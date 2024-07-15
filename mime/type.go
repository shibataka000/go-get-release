package mime

import (
	"mime"
	"path/filepath"
	"slices"
)

// Type represents a MIME type.
type Type string

// Detect mime type from name.
func Detect(name string) Type {
	t := Type(mime.TypeByExtension(filepath.Ext(name)))
	if t == "" {
		return OctedStream
	}
	return t
}

// IsCompressed returns true if mime type is compressed one.
func (t Type) IsCompressed() bool {
	return slices.Contains(compressed, t)
}

// IsOctedStream returns true if mime type is octed stream.
func (t Type) IsOctetStream() bool {
	return t == OctedStream
}
