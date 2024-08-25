package github

// Release represents a GitHub release.
type Release struct {
	tag string
}

// newRelease returns a new GitHub release object.
func newRelease(tag string) Release {
	return Release{
		tag: tag,
	}
}
