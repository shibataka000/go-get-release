package github

import (
	"archive/tar"
	"archive/zip"
	"io"
)

func newTarReader(r io.Reader) (io.Reader, error) {
	for tr := tar.NewReader(r); ; {
		header, err := tr.Next()
		if err != nil {
			return nil, err
		}
		if header.Mode == 0655 {
			return tr, nil
		}
	}
}

func newZipReader(r io.ReaderAt, size int64) (io.ReadCloser, error) {
	zr, err := zip.NewReader(r, size)
	if err != nil {
		return nil, err
	}
	for _, f := range zr.File {
		if f.Mode() == 0655 {
			return f.Open()
		}
	}
	return nil, io.EOF
}
