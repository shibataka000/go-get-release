package github

import (
	"errors"
)

var (
	ErrInvalidRepositoryFullName          = errors.New("repository full name was not 'OWNER/REPO' format")
	ErrUnpairablePattern                  = errors.New("the number of asset patterns and the number of exec binary patterns were not same")
	ErrPatternNotMatched                  = errors.New("no pattern matched with any release asset name")
	ErrUnexpectedMIME                     = errors.New("unexpected mime type")
	ErrExtractingExecBinaryContentFailure = errors.New("extracting exec binary content from release asset content was failed")
)
