package github

import "fmt"

// InvalidRepositoryFullNameError is error raised when repository full name is invalid.
type InvalidRepositoryFullNameError struct {
	fullName string
}

// Error returns an error message.
func (e *InvalidRepositoryFullNameError) Error() string {
	return fmt.Sprintf("acceptable repository full name is 'OWNER/NAME' format, but given name was '%s'", e.fullName)
}

func newInvalidRepositoryFullNameError(fullName string) *InvalidRepositoryFullNameError {
	return &InvalidRepositoryFullNameError{
		fullName: fullName,
	}
}

type AssetNotFoundError struct{}

// Error returns an error message.
func (e *AssetNotFoundError) Error() string {
	return "no asset was found"
}

func newAssetNotFoundError() *AssetNotFoundError {
	return &AssetNotFoundError{}
}

///number of asset patterns and exec binary patterns are not same

type NumberOfAssetPatternsAndExecBinaryPatternsAreNotSameError struct{}
