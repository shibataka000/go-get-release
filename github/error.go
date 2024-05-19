package github

// RepositoryNotFoundError is error raised when no repository was found.
type RepositoryNotFoundError struct{}

func (e *RepositoryNotFoundError) Error() string {
	return "No repository was found."
}

// AssetNotFoundError is error raised when no AssetMeta was found.
type AssetNotFoundError struct{}

func (e *AssetNotFoundError) Error() string {
	return "No asset was found."
}
