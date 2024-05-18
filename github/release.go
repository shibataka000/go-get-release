package github

import (
	"context"
	"fmt"
	"strings"

	"github.com/google/go-github/v48/github"
	"golang.org/x/mod/semver"
)

// Release represents a GitHub release in a repository.
type Release struct {
	Tag string
}

// newRelease return new Release object.
func newRelease(tag string) Release {
	return Release{
		Tag: tag,
	}
}

// semVer return semver formatted release tag.
// For example, if release tag is "v1.2.3", this return "1.2.3".
func (r Release) semVer() (string, error) {
	if !semver.IsValid(r.Tag) && !semver.IsValid(fmt.Sprintf("v%s", r.Tag)) {
		return "", fmt.Errorf("%s is not valid semver", r.Tag)
	}
	return strings.TrimLeft(r.Tag, "v"), nil
}

// ReleaseRepository is repository for Release.
type ReleaseRepository struct {
	client *github.Client
}

// NewReleaseRepository return new ReleaseRepository object.
func NewReleaseRepository(ctx context.Context, token string) *ReleaseRepository {
	return &ReleaseRepository{
		client: newGitHubClient(ctx, token),
	}
}

// Lastet return latest release in a repository.
func (r *ReleaseRepository) latest(ctx context.Context, repo Repository) (Release, error) {
	release, _, err := r.client.Repositories.GetLatestRelease(ctx, repo.Owner, repo.Name)
	if err != nil {
		return Release{}, err
	}
	return newRelease(release.GetTagName()), nil
}
