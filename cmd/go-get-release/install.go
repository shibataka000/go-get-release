package main

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
	"github.com/shibataka000/go-get-release/internal/pkg"
)

func install(name, token, goos, goarch, dir string, showPrompt bool) error {
	p, err := pkg.Find(&pkg.FindInput{
		Name:        name,
		GithubToken: token,
		Goos:        goos,
		Goarch:      goarch,
	})
	if err != nil {
		return err
	}

	if showPrompt {
		fmt.Printf("repo:\t%s/%s\ntag:\t%s\nasset:\t%s\n\n", p.Owner, p.Repo, p.Tag, p.Asset)
		if !prompter.YN("Are you sure to install release binary from above repository?", true) {
			return nil
		}
		fmt.Println()
	}

	tempDir, err := ioutil.TempDir("", "go-get-release-")
	if err != nil {
		return err
	}

	downloadFilePath := filepath.Join(tempDir, p.Asset)
	err = downloadFile(p.DownloadURL, downloadFilePath, showPrompt)
	if err != nil {
		return err
	}

	var downloadBinaryPath string
	if p.IsArchived { // fixme
		err = extract(downloadFilePath, tempDir, p.BinaryName)
		if err != nil {
			return err
		}
		downloadBinaryPath, err = findFile(tempDir, p.BinaryName)
		if err != nil {
			return err
		}
	} else {
		downloadBinaryPath = downloadFilePath
	}

	installBinaryPath := filepath.Join(dir, p.BinaryName)
	err = os.Rename(downloadBinaryPath, installBinaryPath)
	if err != nil {
		return err
	}
	err = os.Chmod(installBinaryPath, 0775)
	if err != nil {
		return err
	}

	err = os.RemoveAll(tempDir)
	if err != nil {
		return err
	}
	return nil
}

func downloadFile(url, filePath string, showProgress bool) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	out, err := os.Create(filePath)
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

func extract(srcFile, dstDir, dstFile string) error {
	if strings.HasSuffix(srcFile, ".zip") {
		return extractZip(srcFile, dstDir)
	} else if strings.HasSuffix(srcFile, ".tar.gz") {
		return extractTarGz(srcFile, dstDir)
	} else if strings.HasSuffix(srcFile, ".tgz") {
		return extractTarGz(srcFile, dstDir)
	} else if strings.HasSuffix(srcFile, ".gz") {
		return extractGz(srcFile, dstDir, dstFile)
	} else {
		return fmt.Errorf("unexpected archive type: %s", srcFile)
	}
}

func extractZip(srcFile, dstDir string) error {
	r, err := zip.OpenReader(srcFile)
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
			path := filepath.Join(dstDir, f.Name)
			os.MkdirAll(path, f.Mode())
		} else {
			buf := make([]byte, f.UncompressedSize)
			_, err = io.ReadFull(rc, buf)
			if err != nil {
				return err
			}

			path := filepath.Join(dstDir, f.Name)
			err := ioutil.WriteFile(path, buf, f.Mode())
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func extractTarGz(srcFile, dstDir string) error {
	inFile, err := os.Open(srcFile)
	if err != nil {
		return err
	}
	defer inFile.Close()

	uncompressedStream, err := gzip.NewReader(inFile)
	if err != nil {
		return err
	}

	tarReader := tar.NewReader(uncompressedStream)

	for {
		header, err := tarReader.Next()

		if err == io.EOF {
			break
		}

		if err != nil {
			return err
		}

		path := filepath.Join(dstDir, header.Name)

		switch header.Typeflag {
		case tar.TypeDir:
			if err := os.MkdirAll(path, 0755); err != nil {
				return err
			}
		case tar.TypeReg:
			dir := filepath.Dir(path)
			if _, err := os.Stat(dir); os.IsNotExist(err) {
				if err := os.MkdirAll(dir, 0755); err != nil {
					return err
				}
			}
			outFile, err := os.Create(path)
			if err != nil {
				return err
			}
			defer outFile.Close()
			if _, err := io.Copy(outFile, tarReader); err != nil {
				return err
			}
		default:
			return fmt.Errorf("fail to extract .tag.gz file: %s %v", header.Name, header.Typeflag)
		}
	}

	return nil
}

func extractGz(srcFile, dstDir, dstFile string) error {
	in, err := os.Open(srcFile)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(filepath.Join(dstDir, dstFile))
	if err != nil {
		return err
	}
	defer out.Close()

	gr, err := gzip.NewReader(in)
	if err != nil {
		return err
	}
	defer gr.Close()

	if _, err := io.Copy(out, gr); err != nil {
		return err
	}
	return nil
}

func findFile(dir, fileName string) (string, error) {
	filePath := ""
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && info.Name() == fileName {
			filePath = path
		}
		return nil
	})
	if err != nil {
		return "", err
	}
	if filePath == "" {
		return "", fmt.Errorf("%s is not found in %s", fileName, dir)
	}
	return filePath, nil
}
