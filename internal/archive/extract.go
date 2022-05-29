package archive

import (
	"archive/tar"
	"archive/zip"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/ulikunitz/xz"
)

// Extract archived or compressed file
func Extract(srcFile, dstDir string) error {
	dstFile := filepath.Join(dstDir, trimExt(filepath.Base(srcFile)))

	switch {
	// archived and compressed
	case strings.HasSuffix(srcFile, ".zip"):
		return extractZip(srcFile, dstDir)
	case strings.HasSuffix(srcFile, ".tar.gz") || strings.HasSuffix(srcFile, ".tgz"):
		return extractTarGz(srcFile, dstDir)
	case strings.HasSuffix(srcFile, ".tar.xz") || strings.HasSuffix(srcFile, ".txz"):
		return extractTarXz(srcFile, dstDir)
	// compressed only
	case strings.HasSuffix(srcFile, ".gz"):
		return extractGz(srcFile, dstFile)
	// unsupported format
	default:
		return fmt.Errorf("unsupported archiving or compression format: %s", srcFile)
	}
}

// Extract zip file
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

			dir := filepath.Dir(path)
			if _, err := os.Stat(dir); os.IsNotExist(err) {
				if err := os.MkdirAll(dir, 0755); err != nil {
					return err
				}
			}

			err := ioutil.WriteFile(path, buf, f.Mode())
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// Extract tar ball
func extractTar(in io.Reader, dstDir string) error {
	tarReader := tar.NewReader(in)

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
			return fmt.Errorf("fail to extract tarball: %s %v", header.Name, header.Typeflag)
		}
	}

	return nil
}

// Extract .tar.gz file
func extractTarGz(srcFile, dstDir string) error {
	in, err := os.Open(srcFile)
	if err != nil {
		return err
	}
	defer in.Close()

	uncompressedStream, err := gzip.NewReader(in)
	if err != nil {
		return err
	}

	return extractTar(uncompressedStream, dstDir)
}

// Extract .tar.xz file
func extractTarXz(srcFile, dstDir string) error {
	in, err := os.Open(srcFile)
	if err != nil {
		return err
	}
	defer in.Close()

	uncompressedStream, err := xz.NewReader(in)
	if err != nil {
		return err
	}

	return extractTar(uncompressedStream, dstDir)
}

// Extract gzip file, which is not archived
func extractGz(srcFile, dstFile string) error {
	in, err := os.Open(srcFile)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(dstFile)
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
