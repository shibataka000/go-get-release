package asset

// Release is interface for GitHub release entity.
type Release interface {
	Tag() string
	SemVer() (string, error)
}
