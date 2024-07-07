package mime

import (
	"mime"
	"path/filepath"
	"slices"
)

// MIME.
type MIME string

// Detect mime from name.
func Detect(name string) MIME {
	m := MIME(mime.TypeByExtension(filepath.Ext(name)))
	if m == "" {
		return OctedStream
	}
	return m
}

// IsCompressed returns true if file is compressed.
func (m MIME) IsCompressed() bool {
	return slices.Contains(compressed, m)
}

// IsOctedStream returns true if file is binary.
func (m MIME) IsOctetStream() bool {
	return m == OctedStream
}
