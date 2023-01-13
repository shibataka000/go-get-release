package asset

import (
	"archive/tar"
	"archive/zip"
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/ulikunitz/xz"
)

// Content of GitHub asset.
type Content struct {
	filename string
	body     []byte
}

// NewContent return content of GitHub asset.
func NewContent(filename string, body []byte) Content {
	return Content{
		filename: filename,
		body:     body,
	}
}

// FindExecBinary return executable binary in GitHub asset.
func (c Content) FindExecBinary(filename string) (ExecBinary, error) {
	var err error
	in := bytes.NewReader(c.body)
	buf := new(bytes.Buffer)

	switch c.ext() {
	case "":
		_, err = io.Copy(buf, in)
	case ".exe":
		_, err = io.Copy(buf, in)
	case ".gz":
		err = extractGzip(buf, in)
	case ".tgz":
		err = copyFileInTarGzip(buf, in, filename)
	case ".tar.gz":
		err = copyFileInTarGzip(buf, in, filename)
	case ".txz":
		err = copyFileInTarXz(buf, in, filename)
	case ".tar.xz":
		err = copyFileInTarXz(buf, in, filename)
	case ".zip":
		err = copyFileInZip(buf, in, filename)
	default:
		err = fmt.Errorf("unsupported file format: %s", c.ext())
	}

	if err != nil {
		return ExecBinary{}, err
	}

	return NewExecBinary(buf.Bytes()), nil
}

// ext return asset name extension.
func (c Content) ext() string {
	ext := filepath.Ext(c.filename)
	name := strings.TrimSuffix(c.filename, ext)
	if filepath.Ext(name) == ".tar" {
		ext = ".tar" + ext
	}
	return ext
}

// extractGzip extract gzip compressed file.
func extractGzip(out io.Writer, in io.Reader) error {
	gzipIn, err := gzip.NewReader(in)
	if err != nil {
		return err
	}
	defer gzipIn.Close()

	_, err = io.Copy(out, gzipIn)
	return err
}

// copyFileInTarGzip extract and copy content from specific file in .tar.gz file.
func copyFileInTarGzip(out io.Writer, in io.Reader, filename string) error {
	gzipIn, err := gzip.NewReader(in)
	if err != nil {
		return err
	}
	defer gzipIn.Close()

	return copyFileInTar(out, gzipIn, filename)
}

// copyFileInTarGzip extract and copy content from specific file in .tar.xz file.
func copyFileInTarXz(out io.Writer, in io.Reader, filename string) error {
	xzIn, err := xz.NewReader(in)
	if err != nil {
		return err
	}

	return copyFileInTar(out, xzIn, filename)
}

// copyFileInTar extract and copy content from specific file in tarball.
func copyFileInTar(out io.Writer, in io.Reader, filename string) error {
	tarIn := tar.NewReader(in)

	for {
		header, err := tarIn.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		switch header.Typeflag {
		case tar.TypeDir:
			// do nothing
		case tar.TypeReg:
			if filepath.Base(header.Name) == filename {
				_, err := io.Copy(out, tarIn)
				return err
			}
		default:
			return fmt.Errorf("unexpected typeflag: %v", header.Typeflag)
		}
	}

	return fmt.Errorf("file '%s' was not found in tarball", filename)
}

// copyFileInTarGzip extract and copy content from specific file in zip file.
func copyFileInZip(out io.Writer, in io.Reader, filename string) error {
	tempFile, err := os.CreateTemp("", "*.zip")
	if err != nil {
		return err
	}
	_, err = io.Copy(tempFile, in)
	if err != nil {
		return err
	}
	defer os.Remove(tempFile.Name())
	defer tempFile.Close()

	rc, err := zip.OpenReader(tempFile.Name())
	if err != nil {
		return err
	}
	defer rc.Close()

	for _, f := range rc.File {
		if !f.FileInfo().IsDir() && filepath.Base(f.Name) == filename {
			fileIn, err := f.Open()
			if err != nil {
				return err
			}
			defer fileIn.Close()

			_, err = io.Copy(out, fileIn)
			return err
		}
	}

	return fmt.Errorf("file '%s' was not found in zip file", filename)
}
