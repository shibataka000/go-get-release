package github

import (
	"bytes"
	"context"
	"net/http"
	"net/url"
	"slices"
	"text/template"

	"github.com/google/go-github/v62/github"
	"github.com/shibataka000/go-get-release/mime"
	"github.com/shibataka000/go-get-release/platform"
	"golang.org/x/oauth2"
)

// Asset represents a GitHub release asset in a repository.
type Asset struct {
	DownloadURL *url.URL
}

// newAsset returns a new GitHub release asset object.
func newAsset(downloadURL *url.URL) Asset {
	return Asset{
		DownloadURL: downloadURL,
	}
}

// os returns an os detected by url to download GitHub release asset.
func (a Asset) os() platform.OS {
	os, _ := platform.Detect(a.DownloadURL.String())
	return os
}

// arch returns an arch detected by url to download GitHub release asset.
func (a Asset) arch() platform.Arch {
	_, arch := platform.Detect(a.DownloadURL.String())
	if arch == "" {
		return "amd64"
	}
	return arch
}

// mime returns a mime type.
func (a Asset) mime() mime.MIME {
	return mime.Detect(a.DownloadURL.String())
}

// hasExecBinary returns true if GitHub release asset may have executable binary.
func (a Asset) hasExecBinary() bool {
	return a.mime().IsCompressed() || a.mime().IsOctetStream()
}

// AssetList is a list of GitHub release asset in a repository.
type AssetList []Asset

// find a GitHub release asset which has executable binary for specified platform.
func (s AssetList) find(os platform.OS, arch platform.Arch) (Asset, error) {
	index := slices.IndexFunc(s, func(asset Asset) bool {
		return asset.os() == os && asset.arch() == arch && asset.hasExecBinary()
	})
	if index == -1 {
		return Asset{}, &AssetNotFoundError{}
	}
	return s[index], nil
}

// AssetTemplate is a template of GitHub release asset in a repository.
type AssetTemplate struct {
	downloadURL *template.Template
}

// newAssetTemplate returns a new GitHub release asset template object.
func newAssetTemplate(downloadURL *template.Template) AssetTemplate {
	return AssetTemplate{
		downloadURL: downloadURL,
	}
}

// execute applies an asset template to the GitHub release object, and returns it as GitHub release asset.
func (a AssetTemplate) execute(release Release) (Asset, error) {
	buf := new(bytes.Buffer)
	data := struct {
		Tag    string
		SemVer string
	}{
		Tag:    release.tag,
		SemVer: release.semver(),
	}
	err := a.downloadURL.Execute(buf, data)
	if err != nil {
		return Asset{}, err
	}
	downloadURL, err := url.Parse(buf.String())
	if err != nil {
		return Asset{}, err
	}
	return newAsset(downloadURL), nil
}

// AssetTemplateList is a list of template of Github release asset in a repository.
type AssetTemplateList []AssetTemplate

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

// list returns a list of GitHub release asset.
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
			downloadURL, err := url.Parse(githubAsset.GetBrowserDownloadURL())
			if err != nil {
				return nil, err
			}
			assets = append(assets, newAsset(downloadURL))
		}
		page = resp.NextPage
	}

	// List GitHub release assets on external server.
	if externalAssets, ok := externalAssets[repo]; ok {
		for _, externalAsset := range externalAssets {
			asset, err := externalAsset.execute(release)
			if err != nil {
				return nil, err
			}
			assets = append(assets, asset)
		}
	}

	return assets, nil
}
