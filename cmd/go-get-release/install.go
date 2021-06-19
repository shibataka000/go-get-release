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

	downloadPath := filepath.Join(tempDir, p.Asset)
	err = downloadFile(downloadPath, p.DownloadURL, showPrompt)
	if err != nil {
		return err
	}

	var oldBinaryPath string
	if p.IsArchived {
		err = extract(downloadPath, tempDir, p.BinaryName)
		if err != nil {
			return err
		}
		oldBinaryPath, err = searchBinaryFilePath(tempDir, p.BinaryName)
		if err != nil {
			return err
		}
	} else {
		oldBinaryPath = downloadPath
	}

	newBinaryPath := filepath.Join(dir, p.BinaryName)
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

func extract(src, dst, binaryName string) error {
	if strings.HasSuffix(src, ".zip") {
		return extractZip(src, dst)
	} else if strings.HasSuffix(src, ".tar.gz") {
		return extractTarGz(src, dst)
	} else if strings.HasSuffix(src, ".gz") {
		return extractGz(src, dst, binaryName)
	} else {
		return fmt.Errorf("unexpected archive type: %s", src)
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

	for {
		header, err := tarReader.Next()

		if err == io.EOF {
			break
		}

		if err != nil {
			return err
		}

		path := filepath.Join(dst, header.Name)

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

func extractGz(src, dst, binaryName string) error {
	in, err := os.Open(src)
	if err != nil {
		panic(err)
	}
	defer in.Close()

	out, err := os.Create(filepath.Join(dst, binaryName))
	if err != nil {
		panic(err)
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
		return "", fmt.Errorf("no binary file found in archived file: %s", path)
	}
	return binaryPath, nil
}
