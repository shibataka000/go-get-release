package github

import (
	"archive/tar"
	"archive/zip"
	"io"
)

// newExecBinaryReaderInTar returns a [io.Reader] to read exec binary in tarball.
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

// newExecBinaryReaderInZip returns a [io.ReadCloser] to read exec binary in zip file.
// Closing [io.ReadCloser] is caller's responsibility.
func newExecBinaryReaderInZip(r io.ReaderAt, size int64) (io.ReadCloser, error) {
	zr, err := zip.NewReader(r, size)
	if err != nil {
		return nil, err
	}
	for _, f := range zr.File {
		if f.Mode() == 0755 {
			return f.Open()
		}
	}
	return nil, io.EOF
}
