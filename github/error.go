package github

import "fmt"

// InvalidRepositoryFullNameFormatError is error raised when repository full name is not 'OWNER/NAME' format.
type InvalidRepositoryFullNameFormatError struct {
	fullName string
}

func (e *InvalidRepositoryFullNameFormatError) Error() string {
	return fmt.Sprintf("'%s' is not valid repository full name. Its format must be 'OWNER/NAME'.", e.fullName)
}

// AssetNotFoundError is error raised when no AssetMeta was found.
type AssetNotFoundError struct{}

func (e *AssetNotFoundError) Error() string {
	return "No asset was found."
}

// AssetNotFoundError is error raised when no ExecutableBinaryMeta was found.
type ExecutableBinaryNotFoundError struct{}

func (e *ExecutableBinaryNotFoundError) Error() string {
	return "No executable binary was found."
}
