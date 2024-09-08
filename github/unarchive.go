package github

import (
	"archive/tar"
	"archive/zip"
	"bytes"
	"io"
)

// newExecBinaryReaderFromTar returns a reader to read exec binary from tarball.
func newExecBinaryReaderFromTar(r io.Reader) (io.Reader, error) {
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

// newExecBinaryReaderFromZip returns a reader to read exec binary from zip file.
func newExecBinaryReaderFromZip(r io.Reader) (io.Reader, error) {
	br, err := newBytesReader(r)
	if err != nil {
		return nil, err
	}

	zr, err := zip.NewReader(br, br.Size())
	if err != nil {
		return nil, err
	}

	for _, f := range zr.File {
		if f.Mode() == 0655 {
			rc, err := f.Open()
			if err != nil {
				return nil, err
			}
			defer rc.Close()
			return newBytesReader(rc)
		}
	}

	return nil, io.EOF
}

// newBytesReader reads all data from r and return a new [bytes.Reader] object.
func newBytesReader(r io.Reader) (*bytes.Reader, error) {
	b, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(b), nil
}
