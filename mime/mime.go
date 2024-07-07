package mime

import (
	"mime"
	"path/filepath"
	"slices"
)

const (
	// Gz represents MIME for gzip file.
	Gz = "application/gzip"
	// Tar represents MIME for tarball.
	Tar = "application/x-tar"
	// Xz represents MIME for xz file.
	Xz = "application/x-xz"
	// Zip represents MIME for zip file.
	Zip = "application/zip"
	// OctedStream represents MIME for binary file.
	OctedStream = "application/octet-stream"
)

// MIME.
type MIME string

// Detect mime from name.
func Detect(name string) MIME {
	return MIME(mime.TypeByExtension(filepath.Ext(name)))
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
