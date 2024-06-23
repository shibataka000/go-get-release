package mime

import (
	"io"
	"slices"

	"github.com/gabriel-vasile/mimetype"
)

const (
	// Deb
	Deb = "application/vnd.debian.binary-package"
	// Gz
	Gz = "application/gzip"
	// Msi
	Msi = "application/x-ms-installer"
	// Rpm
	Rpm = "application/x-rpm"
	// Tar
	Tar = "application/x-tar"
	// Txt
	Txt = "	text/plain"
	// Xz
	Xz = "application/x-xz"
	// Zip
	Zip = "application/zip"
	// OctedStream
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
