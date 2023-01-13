package release

import (
	"fmt"
	"strings"

	"golang.org/x/mod/semver"
)

// Release in GitHub repository
type Release struct {
	id  int64
	tag string
}

// New release instance.
func New(id int64, tag string) Release {
	return Release{
		id:  id,
		tag: tag,
	}
}

// ID return release id.
func (r Release) ID() int64 {
	return r.id
}

// Tag return release tag.
func (r Release) Tag() string {
	return r.tag
}

// SemVer return semver formatted release tag.
// For example, if release tag is "v1.2.3", this return "1.2.3".
func (r Release) SemVer() (string, error) {
	if !semver.IsValid(r.tag) && !semver.IsValid(fmt.Sprintf("v%s", r.tag)) {
		return "", fmt.Errorf("%s is not valid semver", r.tag)
	}
	return strings.TrimLeft(r.tag, "v"), nil
}
