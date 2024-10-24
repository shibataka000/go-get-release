package github

import "github.com/gabriel-vasile/mimetype"

func init() {
	mimetype.SetLimit(0)
}
