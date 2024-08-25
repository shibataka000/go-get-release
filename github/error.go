package github

import "fmt"

// InvalidRepositoryFullNameFormatError is error raised when repository full name is not 'OWNER/NAME' format.
type InvalidRepositoryFullNameFormatError struct {
	name string
}

// Error returns an error message.
func (e *InvalidRepositoryFullNameFormatError) Error() string {
	return fmt.Sprintf("acceptable repository full name format is 'OWNER/NAME', but given name was '%s'", e.name)
}

// FindingAssetFailureError is error raised when finding asset was failed, especially zero or two or more assets was found.
type FindingAssetFailureError struct {
	assets AssetList
}

// Error returns an error message.
func (e *FindingAssetFailureError) Error() string {
	switch len(e.assets) {
	case 0:
		return "no asset was found"
	default:
		return fmt.Sprintf("too many assets was found: %v", e.assets)
	}
}
