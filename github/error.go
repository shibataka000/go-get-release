package github

import "fmt"

// InvalidRepositoryFullNameFormatError is error raised when repository full name is not 'OWNER/NAME' format.
type InvalidRepositoryFullNameFormatError struct {
	name string
}

func (e *InvalidRepositoryFullNameFormatError) Error() string {
	return fmt.Sprintf("Acceptable repository full name format is 'OWNER/NAME', but given name was '%s'.", e.name)
}

// AssetNotFoundError is error raised when GitHub release asset was not found.
type AssetNotFoundError struct{}

func (e *AssetNotFoundError) Error() string {
	return "No asset was found."
}
