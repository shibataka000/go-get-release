package mime

import (
	"io"

	"github.com/gabriel-vasile/mimetype"
)

type MIME string

func DetectReader(r io.Reader) (MIME, error) {
	mime, err := mimetype.DetectReader(r)
	if err != nil {
		return "", err
	}
	return MIME(mime.String()), nil
}

func (m MIME) IsArchived() bool {
	return true
}

func (m MIME) IsCompressed() bool {
	return true
}

func (m MIME) IsOctetStream() bool {
	return true
}
