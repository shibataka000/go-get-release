package github

import (
	"archive/tar"
	"archive/zip"
	"bytes"
	"io"
)

// newExecBinaryReaderInTar returns a reader to read exec binary in tarball.
func newExecBinaryReaderInTar(r io.Reader) (io.Reader, error) {
	for tr := tar.NewReader(r); ; {
		header, err := tr.Next()
		if err != nil {
			return nil, err
		}
		if header.Mode == 0755 {
			return tr, nil
		}
	}
}

// newExecBinaryReaderInZip returns a reader to read exec binary in zip file.
func newExecBinaryReaderInZip(r io.Reader) (io.Reader, error) {
	br, err := newBytesReader(r)
	if err != nil {
		return nil, err
	}

	zr, err := zip.NewReader(br, br.Size())
	if err != nil {
		return nil, err
	}

	for _, f := range zr.File {
		if f.Mode() == 0755 {
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
