package github

import (
	"fmt"

	"github.com/gabriel-vasile/mimetype"
)

// InvalidRepositoryError is error raised when repository is invalid.
type InvalidRepositoryError struct {
	format string
	a      []any
}

// newInvalidRepositoryError returns a new InvalidRepositoryError object.
func newInvalidRepositoryError(format string, a ...any) *InvalidRepositoryError {
	return &InvalidRepositoryError{
		format: format,
		a:      a,
	}
}

// Error returns an error message.
func (e *InvalidRepositoryError) Error() string {
	return fmt.Sprintf("repository was invalid: %s", fmt.Sprintf(e.format, e.a...))
}

// InvalidPatternError is error raised when pattern is invalid.
type InvalidPatternError struct {
	format string
	a      []any
}

// newInvalidPatternError returns a new InvalidPatternError object.
func newInvalidPatternError(format string, a ...any) *InvalidPatternError {
	return &InvalidPatternError{
		format: format,
		a:      a,
	}
}

// Error returns an error message.
func (e *InvalidPatternError) Error() string {
	return fmt.Sprintf("pattern was invalid: %s", fmt.Sprintf(e.format, e.a...))
}

type AssetNotFoundError struct{}

func newAssetNotFoundError() *AssetNotFoundError {
	return &AssetNotFoundError{}
}

func (e *AssetNotFoundError) Error() string {
	return "no asset was found"
}

type UnsupportedMIMEError struct {
	mime *mimetype.MIME
}

func newUnsupportedMIMEError(mime *mimetype.MIME) *UnsupportedMIMEError {
	return &UnsupportedMIMEError{mime}
}

func (e *UnsupportedMIMEError) Error() string {
	return fmt.Sprintf("mime type was unsupported: %s", e.mime.String())
}
