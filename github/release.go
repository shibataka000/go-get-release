package github

import (
	"fmt"
	"strings"

	"golang.org/x/mod/semver"
)

// Release represents a GitHub release in a repository.
type Release struct {
	tag string
}

// newRelease returns new GitHub release object.
func newRelease(tag string) Release {
	return Release{
		tag: tag,
	}
}

// semver returns semver formatted release tag.
// For example, if release tag is "v1.2.3", this returns "1.2.3".
// If release tag is not valid format, this returns empty string.
func (r Release) semver() string {
	switch {
	case semver.IsValid(r.tag):
		return strings.TrimLeft(r.tag, "v")
	case semver.IsValid(fmt.Sprintf("v%s", r.tag)):
		return r.tag
	default:
		return ""
	}
}
