package github

import "fmt"

// InvalidRepositoryError is error raised when repository full name is invalid.
type InvalidRepositoryError struct {
	fullName string
}

func newInvalidRepositoryError(fullName string) *InvalidRepositoryError {
	return &InvalidRepositoryError{
		fullName: fullName,
	}
}

// Error returns an error message.
func (e *InvalidRepositoryError) Error() string {
	return fmt.Sprintf("acceptable repository full name is 'OWNER/NAME' format, but given name was '%s'", e.fullName)
}

type InvalidPatternError struct{}

func newInvalidPatternError() *InvalidPatternError {
	return &InvalidPatternError{}
}

func (e *InvalidPatternError) Error() string {
	return ""
}

type AssetNotFoundError struct{}

func newAssetNotFoundError() *AssetNotFoundError {
	return &AssetNotFoundError{}
}

func (e *AssetNotFoundError) Error() string {
	return "no asset was found"
}
