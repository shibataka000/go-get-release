package github

import (
	"archive/tar"
	"archive/zip"
	"io"
)

func newTarReader(r io.Reader) (io.Reader, error) {
	tr := tar.NewReader(r)
	for {
		header, err := tr.Next()
		if err != nil {
			return nil, err
		}
		if header.Typeflag == tar.TypeReg && header.Mode == 0x655 {
			return tr, nil
		}
	}
}

func newZipReader(r io.Reader) (io.Reader, error) {
	zip.NewReader()
	return nil, nil
}
