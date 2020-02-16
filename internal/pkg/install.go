package pkg

import (
	"archive/tar"
	"archive/zip"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/Songmu/prompter"
	"github.com/cheggaaa/pb"
)

// Option to execute commands.
type Option struct {
	GithubToken string
	OS          string
	Arch        string
	InstallDir  string
	ShowPrompt  bool
}

// Install golan release binary.
func Install(pkgName string, option *Option) error {
	pkgInfo, err := getPkgInfo(pkgName, option)
	if err != nil {
		return err
	}

	if option.ShowPrompt {
		fmt.Printf("repo:\t%s/%s\ntag:\t%s\nasset:\t%s\n\n", pkgInfo.owner, pkgInfo.repo, pkgInfo.tag, pkgInfo.asset)
		if !prompter.YN("Are you sure to install release binary from above repository?", true) {
			return nil
		}
		fmt.Println()
	}

	tempDir, err := ioutil.TempDir("", "go-get-release-")
	if err != nil {
		return err
	}

	downloadPath := filepath.Join(tempDir, pkgInfo.asset)
	err = downloadFile(downloadPath, pkgInfo.downloadURL, option.ShowPrompt)
	if err != nil {
		return err
	}

	var oldBinaryPath string
	if pkgInfo.isArchived {
		err = extract(downloadPath, tempDir)
		if err != nil {
			return err
		}
		oldBinaryPath, err = searchBinaryFilePath(tempDir, pkgInfo.binary)
		if err != nil {
			return nil
		}
	} else {
		oldBinaryPath = downloadPath
	}

	newBinaryPath := filepath.Join(option.InstallDir, pkgInfo.binary)
	err = os.Rename(oldBinaryPath, newBinaryPath)
	if err != nil {
		return err
	}
	err = os.Chmod(newBinaryPath, 0775)
	if err != nil {
		return err
	}

	err = os.RemoveAll(tempDir)
	if err != nil {
		return err
	}
	return nil
}

func downloadFile(filepath, url string, showProgress bool) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	r := resp.Body
	if showProgress {
		bar := pb.Full.Start64(resp.ContentLength)
		r = bar.NewProxyReader(resp.Body)
	}
	_, err = io.Copy(out, r)
	return err
}

func extract(src, dst string) error {
	if strings.HasSuffix(src, ".zip") {
		return extractZip(src, dst)
	} else if strings.HasSuffix(src, ".tar.gz") {
		return extractTarGz(src, dst)
	} else {
		return fmt.Errorf("Unexpected archive type: %s", src)
	}
}

func extractZip(src, dst string) error {
	r, err := zip.OpenReader(src)
	if err != nil {
		return err
	}
	defer r.Close()

	for _, f := range r.File {
		rc, err := f.Open()
		if err != nil {
			return err
		}
		defer rc.Close()

		if f.FileInfo().IsDir() {
			path := filepath.Join(dst, f.Name)
			os.MkdirAll(path, f.Mode())
		} else {
			buf := make([]byte, f.UncompressedSize)
			_, err = io.ReadFull(rc, buf)
			if err != nil {
				return err
			}

			path := filepath.Join(dst, f.Name)
			err := ioutil.WriteFile(path, buf, f.Mode())
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func extractTarGz(src, dst string) error {
	inFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer inFile.Close()

	uncompressedStream, err := gzip.NewReader(inFile)
	if err != nil {
		return err
	}

	tarReader := tar.NewReader(uncompressedStream)

	for true {
		header, err := tarReader.Next()

		if err == io.EOF {
			break
		}

		if err != nil {
			return err
		}

		switch header.Typeflag {
		case tar.TypeDir:
			if err := os.Mkdir(filepath.Join(dst, header.Name), 0755); err != nil {
				return err
			}
		case tar.TypeReg:
			outFile, err := os.Create(filepath.Join(dst, header.Name))
			if err != nil {
				return err
			}
			defer outFile.Close()
			if _, err := io.Copy(outFile, tarReader); err != nil {
				return err
			}
		default:
			return fmt.Errorf("Fail to extract .tag.gz file: %s %v", header.Name, header.Typeflag)
		}
	}

	return nil
}

func searchBinaryFilePath(path, binaryName string) (string, error) {
	binaryPath := ""
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && info.Name() == binaryName {
			binaryPath = path
		}
		return nil
	})
	if err != nil {
		return "", err
	}
	if binaryPath == "" {
		return "", fmt.Errorf("No binary file found in archived file: %s", path)
	}
	return binaryPath, nil
}
