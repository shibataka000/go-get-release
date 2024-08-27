package github

import "fmt"

// InvalidRepositoryFullNameError is error raised when repository full name is invalid.
type InvalidRepositoryFullNameError struct {
	name string
}

// Error returns an error message.
func (e *InvalidRepositoryFullNameError) Error() string {
	return fmt.Sprintf("acceptable repository full name is 'OWNER/NAME' format, but given was '%s'", e.name)
}

type InvalidPatternError struct {
	pattern string
}

// Error returns an error message.
func (e *InvalidPatternError) Error() string {
	return fmt.Sprintf("acceptable pattern is 'ASSET_NAME_PATTERN=EXEC_BINARY_NAME_PATTERN' format, but given was '%s'", e.pattern)
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

// gh release install --asset (p<name>.*)linux --exec-binary
