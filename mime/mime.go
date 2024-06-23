package mime

import (
	"io"
	"slices"

	"github.com/gabriel-vasile/mimetype"
)

const (
	// Deb represents MIME for debian binary package file.
	Deb = "application/vnd.debian.binary-package"
	// Gz represents MIME for gzip file.
	Gz = "application/gzip"
	// Msi represents MIME for microsoft installer file.
	Msi = "application/x-ms-installer"
	// Rpm represents MIME for rpm package file.
	Rpm = "application/x-rpm"
	// Tar represents MIME for tarball.
	Tar = "application/x-tar"
	// Txt represents MIME for text file.
	Txt = "	text/plain"
	// Xz represents MIME for xz file.
	Xz = "application/x-xz"
	// Zip represents MIME for zip file.
	Zip = "application/zip"
	// OctedStream represents MIME for binary file.
	OctedStream = "application/octet-stream"
)

// MIME.
type MIME string

// DetectReader returns the MIME type of the provided reader.
func DetectReader(r io.Reader, limit uint32) (MIME, error) {
	mimetype.SetLimit(limit)
	mime, err := mimetype.DetectReader(r)
	if err != nil {
		return "", err
	}
	return MIME(mime.String()), nil
}

// IsArchived returns true if file is archived.
func (m MIME) IsArchived() bool {
	return slices.Contains([]MIME{Tar, Zip}, m)
}

// IsCompressed returns true if file is compressed.
func (m MIME) IsCompressed() bool {
	return slices.Contains([]MIME{Gz, Zip, Xz}, m)
}

// IsOctedStream returns true if file is binary.
func (m MIME) IsOctetStream() bool {
	return m == OctedStream
}
