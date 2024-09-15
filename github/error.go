package github

import (
	"errors"
)

var (
	ErrInvalidRepositoryFullName       = errors.New("repository full name was not 'OWNER/REPO' format")
	ErrPatternNotMatched               = errors.New("no pattern was matched to any release asset name")
	ErrUnexpectedMIME                  = errors.New("unexpected mime type")
	ErrExtractExecBinaryContentFailure = errors.New("extracting exec binary content from release asset content was failed")
	ErrUnpairablePattern               = errors.New("the number of asset pattern and the number of exec binary pattern are not same")
)
