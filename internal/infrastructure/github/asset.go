package github

import (
	"context"

	"github.com/google/go-github/v48/github"
	model "github.com/shibataka000/go-get-release/internal/domain/model/asset"
	"github.com/shibataka000/go-get-release/internal/domain/model/release"
	"github.com/shibataka000/go-get-release/internal/domain/model/repository"
)

// List assets in GitHub release.
func (r *Repository) ListAssetMetadata(ctx context.Context, repo repository.Repository, release release.Release) ([]model.Metadata, error) {
	meta := []model.Metadata{}
	for page := 1; page != 0; {
		assets, resp, err := r.client.Repositories.ListReleaseAssets(ctx, repo.Owner(), repo.Name(), release.ID(), &github.ListOptions{
			Page:    page,
			PerPage: 100,
		})
		if err != nil {
			return meta, err
		}
		for _, asset := range assets {
			meta = append(meta, model.NewMetadata(asset.GetBrowserDownloadURL()))
		}
		page = resp.NextPage
	}
	return meta, nil
}
