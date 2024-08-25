package github

import (
	"context"
	"net/http"
	"net/url"
	"path"
	"regexp"
	"slices"

	"github.com/google/go-github/v62/github"
	"golang.org/x/oauth2"
)

// Asset represents a GitHub release asset.
type Asset struct {
	DownloadURL *url.URL
}

// newAsset returns a new GitHub release asset object.
func newAsset(downloadURL *url.URL) Asset {
	return Asset{
		DownloadURL: downloadURL,
	}
}

// name returns a name of GitHub release asset.
func (a Asset) name() string {
	return path.Base(a.DownloadURL.String())
}

// AssetList represents a list of GitHub release assets.
type AssetList []Asset

// find a GitHub release asset which matches any of given patterns.
// If one or more assets match patterns, this returns first one.
func (al AssetList) find(patterns []*regexp.Regexp) (Asset, error) {
	index := slices.IndexFunc(al, func(a Asset) bool {
		return slices.ContainsFunc(patterns, func(p *regexp.Regexp) bool {
			return p.Match([]byte(a.name()))
		})
	})
	if index == -1 {
		return Asset{}, &AssetNotFoundError{}
	}
	return al[index], nil
}

// AssetRepository is a repository for a GitHub release asset.
type AssetRepository struct {
	client *github.Client
}

// NewAssetRepository returns a new AssetRepository object.
func NewAssetRepository(ctx context.Context, token string) *AssetRepository {
	var httpClient *http.Client
	if token != "" {
		tokenSource := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})
		httpClient = oauth2.NewClient(ctx, tokenSource)
	}
	return &AssetRepository{
		client: github.NewClient(httpClient),
	}
}

// list returns a list of GitHub release assets.
func (r *AssetRepository) list(ctx context.Context, repo Repository, release Release) (AssetList, error) {
	assets := AssetList{}

	// Get GitHub release ID.
	githubRelease, _, err := r.client.Repositories.GetReleaseByTag(ctx, repo.owner, repo.name, release.tag)
	if err != nil {
		return nil, err
	}
	releaseID := *githubRelease.ID

	// List GitHub release assets.
	for page := 1; page != 0; {
		githubAssets, resp, err := r.client.Repositories.ListReleaseAssets(ctx, repo.owner, repo.name, releaseID, &github.ListOptions{
			Page: page,
		})
		if err != nil {
			return nil, err
		}
		for _, githubAsset := range githubAssets {
			downloadURL := githubAsset.GetBrowserDownloadURL()
			url, err := url.Parse(downloadURL)
			if err != nil {
				return nil, err
			}
			assets = append(assets, newAsset(url))
		}
		page = resp.NextPage
	}

	return assets, nil
}
