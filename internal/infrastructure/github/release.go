package github

import (
	"context"

	model "github.com/shibataka000/go-get-release/internal/domain/model/release"
	"github.com/shibataka000/go-get-release/internal/domain/model/repository"
)

// LatestRelease return latest GitHub release.
func (r *Repository) LatestRelease(ctx context.Context, repo repository.Repository) (model.Release, error) {
	release, _, err := r.client.Repositories.GetLatestRelease(ctx, repo.Owner(), repo.Name())
	if err != nil {
		return model.Release{}, err
	}
	return model.New(release.GetID(), release.GetTagName()), nil
}

// FindReleaseByTag return GitHub release by tag.
func (r *Repository) FindReleaseByTag(ctx context.Context, repo repository.Repository, tag string) (model.Release, error) {
	release, _, err := r.client.Repositories.GetReleaseByTag(ctx, repo.Owner(), repo.Name(), tag)
	if err != nil {
		return model.Release{}, err
	}
	return model.New(release.GetID(), release.GetTagName()), nil
}
