package github

import (
	"errors"
)

var (
	ErrInvalidRepositoryFullName = errors.New("repository full name was not 'OWNER/REPO' format")
	ErrUnpairablePattern         = errors.New("the number of asset patterns and the number of exec binary patterns were not same")
	ErrAssetNotFound             = errors.New("no asset was found")
	ErrUnsupportedMIME           = errors.New("mime type was not supported")
)
