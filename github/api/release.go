package api

import (
	"context"

	"github.com/google/go-github/v48/github"
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
