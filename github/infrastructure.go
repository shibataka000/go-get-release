package github

import (
	"context"
	"net/http"

	"github.com/google/go-github/v48/github"
	"golang.org/x/oauth2"
)

// InfrastructureRepository is responsible to persist entity.
type InfrastructureRepository struct {
	github *github.Client
}

// NewInfrastructureRepository return new infrastructure repository instance.
func NewInfrastructureRepository(ctx context.Context, token string) *InfrastructureRepository {
	var httpClient *http.Client
	if token != "" {
		tokenSource := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})
		httpClient = oauth2.NewClient(ctx, tokenSource)
	}
	githubClient := github.NewClient(httpClient)
	return &InfrastructureRepository{
		github: githubClient,
	}
}

// SearchRepository search GitHub repository.
func (r *InfrastructureRepository) SearchRepository(ctx context.Context, query string) (Repository, error) {
	result, _, err := r.github.Search.Repositories(ctx, query, &github.SearchOptions{})
	if err != nil {
		return Repository{}, err
	}
	repos := result.Repositories
	if len(repos) == 0 {
		return Repository{}, NewNotFoundError("no repository was found by query '%s'", query)
	}
	repo := repos[0]
	return NewRepository(repo.GetOwner().GetLogin(), repo.GetName()), nil
}

// LatestRelease return latest GitHub release.
func (r *InfrastructureRepository) LatestRelease(ctx context.Context, repo Repository) (Release, error) {
	release, _, err := r.github.Repositories.GetLatestRelease(ctx, repo.Owner, repo.Name)
	if err != nil {
		return Release{}, err
	}
	return NewRelease(release.GetTagName()), nil
}

// ListGitHubAssets list assets in GitHub release.
func (r *InfrastructureRepository) ListAssetMeta(ctx context.Context, repo Repository, release Release) ([]AssetMeta, error) {
	githubRelease, _, err := r.github.Repositories.GetReleaseByTag(ctx, repo.Owner, repo.Name, release.Tag)
	if err != nil {
		return []AssetMeta{}, err
	}
	releaseID := githubRelease.GetID()

	result := []AssetMeta{}
	for page := 1; page != 0; {
		assets, resp, err := r.github.Repositories.ListReleaseAssets(ctx, repo.Owner, repo.Name, releaseID, &github.ListOptions{
			Page:    page,
			PerPage: 100,
		})
		if err != nil {
			return result, err
		}
		for _, asset := range assets {
			downloadURL := URL(asset.GetBrowserDownloadURL())
			result = append(result, NewAssetMeta(downloadURL))
		}
		page = resp.NextPage
	}
	return result, nil
}
