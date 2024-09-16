package github

import (
	"errors"
)

var (
	ErrInvalidRepositoryFullName          = errors.New("repository full name was not 'OWNER/REPO' format")
	ErrPatternNotMatched                  = errors.New("no pattern matched with any release asset name")
	ErrUnexpectedMIME                     = errors.New("unexpected mime type")
	ErrExtractingExecBinaryContentFailure = errors.New("extracting exec binary content from release asset content was failed")
)
