package github

import (
	"archive/tar"
	"archive/zip"
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
// It is the caller's responsibility to close the ReadCloser.
func newExecBinaryReaderFromZip(r io.ReaderAt, size int64) (io.ReadCloser, error) {
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
