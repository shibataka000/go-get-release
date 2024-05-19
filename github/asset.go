package github

import (
	"context"

	"github.com/google/go-github/v48/github"
	"github.com/shibataka000/go-get-release/runtime"
	"github.com/shibataka000/go-get-release/url"
	"gopkg.in/yaml.v3"
)

// AssetMeta represents a GitHub release asset in a repository.
type AssetMeta struct {
	DownloadURL url.URL
	GOOS        runtime.GOOS
	GOARCH      runtime.GOARCH
}

// AssetMetaList is a list of AssetMeta.
type AssetMetaList []AssetMeta

// newAssetMeat return new AssetMeta object.
func newAssetMeta(downloadURL url.URL, goos runtime.GOOS, goarch runtime.GOARCH) AssetMeta {
	return AssetMeta{
		DownloadURL: downloadURL,
		GOOS:        goos,
		GOARCH:      goarch,
	}
}

// AssetRepository is repository for Asset.
type AssetRepository struct {
	client *github.Client
}

// NewAssetRepository return new AssetRepository object.
func NewAssetRepository(ctx context.Context, token string) *AssetRepository {
	return &AssetRepository{
		client: newGitHubClient(ctx, token),
	}
}

// list return lit of AssetMeta in a GitHub release.
func (r *AssetRepository) list(ctx context.Context, repo Repository, release Release) (AssetMetaList, error) {
	githubRelease, _, err := r.client.Repositories.GetReleaseByTag(ctx, repo.Owner, repo.Name, release.Tag)
	if err != nil {
		return nil, err
	}

	result := AssetMetaList{}
	for page := 1; page != 0; {
		assets, resp, err := r.client.Repositories.ListReleaseAssets(ctx, repo.Owner, repo.Name, *githubRelease.ID, &github.ListOptions{
			Page:    page,
			PerPage: 100,
		})
		if err != nil {
			return result, err
		}
		for _, asset := range assets {
			downloadURL := url.URL(asset.GetBrowserDownloadURL())
			goos, goarch := runtime.Guess(downloadURL.Base())
			result = append(result, newAssetMeta(downloadURL, goos, goarch))
		}
		page = resp.NextPage
	}
	return result, nil
}

func (r *AssetRepository) listFromBuiltIn(repo Repository, release Release) (AssetMetaList, error) {
	type Entry struct {
		Repository Repository
		Assets     AssetMetaList
	}
	entries := []Entry{}
	err := yaml.Unmarshal(builtin, &entries)
	if err != nil {
		return nil, err
	}
	entry := Entry{}
	found := false
	for _, e := range entries {
		if e.Repository.Owner == repo.Owner && e.Repository.Name == repo.Name {
			entry = e
			found = true
		}
	}
	if found {
		return nil, &AssetNotFoundError{}
	}
	return entry.Assets, nil
}

// find AssetMeta by GOOS/GOARCH.
func (a AssetMetaList) find(goos runtime.GOOS, goarch runtime.GOARCH) (AssetMeta, error) {
	for _, asset := range a {
		if asset.GOOS == goos && asset.GOARCH == goarch {
			return asset, nil
		}
	}
	return AssetMeta{}, &AssetNotFoundError{}
}

// AssetNotFoundError is error raised when try to find AssetMeta by GOOS/GOARCH but no AssetMeta was found.
type AssetNotFoundError struct{}

// Error return error string.
func (e *AssetNotFoundError) Error() string {
	return "No asset was found."
}
