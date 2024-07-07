package github

import "fmt"

// InvalidRepositoryFullNameFormatError is error raised when repository full name is not 'OWNER/NAME' format.
type InvalidRepositoryFullNameFormatError struct {
	name string
}

// Error returns an error message.
func (e *InvalidRepositoryFullNameFormatError) Error() string {
	return fmt.Sprintf("Acceptable repository full name format is 'OWNER/NAME', but given name was '%s'.", e.name)
}

// AssetNotFoundError is error raised when GitHub release asset was not found.
type AssetNotFoundError struct{}

// Error returns an error message.
func (e *AssetNotFoundError) Error() string {
	return "No asset was found."
}
