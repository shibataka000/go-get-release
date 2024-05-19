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

// semver return semver formatted release tag.
// For example, if release tag is "v1.2.3", this return "1.2.3".
// If release tag is invalid format, this returns empty string.
func (r Release) semver() string {
	switch {
	case semver.IsValid(r.Tag):
		return strings.TrimLeft(r.Tag, "v")
	case semver.IsValid(fmt.Sprintf("v%s", r.Tag)):
		return r.Tag
	default:
		return ""
	}
}

// ReleaseRepository is repository for a GitHub release.
type ReleaseRepository struct {
	client  *github.Client
	factory *ReleaseFactory
}

// NewReleaseRepository return new ReleaseRepository object.
func NewReleaseRepository(ctx context.Context, token string, factory *ReleaseFactory) *ReleaseRepository {
	return &ReleaseRepository{
		client:  newGitHubClient(ctx, token),
		factory: factory,
	}
}

// lastet return latest release in a repository.
func (r *ReleaseRepository) latest(ctx context.Context, repo Repository) (Release, error) {
	release, _, err := r.client.Repositories.GetLatestRelease(ctx, repo.Owner, repo.Name)
	if err != nil {
		return Release{}, err
	}
	return r.factory.new(release.GetTagName()), nil
}

// ReleaseFactory is factory to create a new GitHub release object.
type ReleaseFactory struct{}

// NewReleaseFactory return new ReleaseFactory object.
func NewReleaseFactory() *ReleaseFactory {
	return &ReleaseFactory{}
}

// new GitHub release object.
func (f *ReleaseFactory) new(tag string) Release {
	return Release{
		Tag: tag,
	}
}
