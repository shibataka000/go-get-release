package execbinary

import "github.com/shibataka000/go-get-release/internal/domain/model/asset"

// Content of executable binary.
type Content struct {
	body []byte
}

// NewContent return content of executable binary.
func NewContent(body []byte) Content {
	return Content{
		body: body,
	}
}

// NewContentFromAsset return content of executable binary.
func NewContentFromAsset(asset asset.ExecBinary) Content {
	return NewContent(asset.Bytes())
}

// Bytes return bytes of executable binary.
func (c Content) Bytes() []byte {
	return c.body
}
