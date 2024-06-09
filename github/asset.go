package github

import (
	"context"
	"net/http"
	"slices"

	"github.com/google/go-github/v48/github"
	"github.com/shibataka000/go-get-release/mime"
	"github.com/shibataka000/go-get-release/platform"
	"github.com/shibataka000/go-get-release/url"
	"golang.org/x/oauth2"
)

// Asset represents a GitHub release asset in a repository.
type Asset struct {
	DownloadURL url.URL
	MIME        mime.MIME
}

// newAsset returns a new GitHub release asset object.
func newAsset(downloadURL url.URL, mime mime.MIME) Asset {
	return Asset{
		DownloadURL: downloadURL,
		MIME:        mime,
	}
}

// os returns an os detected by asset file name.
func (a Asset) os() platform.OS {
	os, _ := platform.Detect(a.DownloadURL.Base())
	return os
}

// arch returns an arch detected by asset file name.
func (a Asset) arch() platform.Arch {
	_, arch := platform.Detect(a.DownloadURL.Base())
	return arch
}

// hasExecutableBinary returns true if asset may have executable binary.
func (a Asset) hasExecutableBinary() bool {
	return a.MIME.IsArchived() || a.MIME.IsCompressed() || a.MIME.IsOctetStream()
}

// AssetList is a list of GitHub releaset asset.
type AssetList []Asset

// find a GitHub release asset metadata which has executable binary and whose os/arch are same as supplied value.
func (s AssetList) find(os platform.OS, arch platform.Arch) (Asset, error) {
	index := slices.IndexFunc(s, func(asset Asset) bool {
		return asset.hasExecutableBinary() && asset.os() == os && asset.arch() == arch
	})
	if index == -1 {
		return Asset{}, &AssetNotFoundError{}
	}
	return s[index], nil
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

// listFromAPI returns a list of GitHub release asset metadata in a GitHub release using GitHub API.
func (r *AssetRepository) listFromAPI(ctx context.Context, repo Repository, release Release) (AssetList, error) {
	// Get GitHub release ID.
	githubRelease, _, err := r.client.Repositories.GetReleaseByTag(ctx, repo.owner, repo.name, release.tag)
	if err != nil {
		return nil, err
	}
	releaseID := *githubRelease.ID

	// List GitHub release assets.
	githubAssets := []*github.ReleaseAsset{}
	for page := 1; page != 0; {
		assets, resp, err := r.client.Repositories.ListReleaseAssets(ctx, repo.owner, repo.name, releaseID, &github.ListOptions{
			Page: page,
		})
		if err != nil {
			return nil, err
		}
		githubAssets = append(githubAssets, assets...)
		page = resp.NextPage
	}

	// Create AssetList from list of GitHub release asset.
	result := AssetList{}
	for _, githubAsset := range githubAssets {
		downloadURL := url.URL(githubAsset.GetBrowserDownloadURL())
		result = append(result, newAsset(downloadURL, ""))
	}

	return result, nil
}

// listFromBuiltIn returns a list of GitHub release asset metadata from built-in data.
func (r *AssetRepository) listFromBuiltIn(repo Repository, release Release) (AssetList, error) {
	list, err := loadBuiltInData()
	if err != nil {
		return nil, err
	}

	data, err := list.find(repo)
	if err != nil {
		return nil, err
	}

	assets := AssetList{}

	for _, asset := range data.Assets {
		downloadURL, err := asset.render(release)
		if err != nil {
			return nil, err
		}
		assets = append(assets, newAsset(downloadURL, ""))
	}

	return assets, nil
}
