package github

import (
	"errors"
)

var (
	ErrInvalidRepositoryFullName       = errors.New("repository full name was not 'OWNER/REPO' format")
	ErrAssetNotFound                   = errors.New("no asset was found")
	ErrUnsupportedMIME                 = errors.New("mime type was not supported")
	ErrGettingExecBinaryContentFailure = errors.New("getting exec binary content from release asset was failed")
)
