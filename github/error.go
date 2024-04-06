package github

import (
	"fmt"
)

// NotFoundError raised if something is not found.
type NotFoundError struct {
	Message string
}

// UnsupportedFileFormatError raised if file format is not supported.
type UnsupportedFileFormatError struct {
	Format string
}

// InvalidSemVerError raised if version is not semver.
type InvalidSemVerError struct {
	Version string
}

// NewNotFoundError return new NotFoundError instance.
func NewNotFoundError(format string, a ...any) NotFoundError {
	return NotFoundError{
		Message: fmt.Sprintf(format, a...),
	}
}

// NewUnsupportedFileFormatError return new UnsupportedFileFormatError instance.
func NewUnsupportedFileFormatError(format string) UnsupportedFileFormatError {
	return UnsupportedFileFormatError{
		Format: format,
	}
}

// NewInvalidSemVerError return new InvalidSemVerError instance.
func NewInvalidSemVerError(version string) InvalidSemVerError {
	return InvalidSemVerError{
		Version: version,
	}
}

// Error return error message.
func (e NotFoundError) Error() string {
	return e.Message
}

// Error return error message.
func (e UnsupportedFileFormatError) Error() string {
	return fmt.Sprintf("unsupported file format: %s", e.Format)
}

func (e InvalidSemVerError) Error() string {
	return fmt.Sprintf("invalid semver: %s", e.Version)
}
