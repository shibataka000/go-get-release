package github

import (
	"fmt"
	"strings"

	"golang.org/x/mod/semver"
)

// Release is release in GitHub.
type Release struct {
	ID  int64
	Tag string
}

// NewRelease return new release instance.
func NewRelease(id int64, tag string) Release {
	return Release{
		ID:  id,
		Tag: tag,
	}
}

// SemVer return semver formatted release tag.
// For example, if release tag is "v1.2.3", this return "1.2.3".
func (r Release) SemVer() (string, error) {
	if !semver.IsValid(r.Tag) && !semver.IsValid(fmt.Sprintf("v%s", r.Tag)) {
		return "", NewInvalidSemVerError(r.Tag)
	}
	return strings.TrimLeft(r.Tag, "v"), nil
}
