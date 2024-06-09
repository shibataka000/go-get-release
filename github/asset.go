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

// assetDownloadURLTemplatesHostedOnOutside is a list of url template to download asset hosted on server outside of GitHub.
var assetDownloadURLTemplatesHostedOnOutside = map[string][]url.Template{
	"hashicorp/terraform": {
		"https://releases.hashicorp.com/terraform/{{.SemVer}}/terraform_{{.SemVer}}_linux_amd64.zip",
		"https://releases.hashicorp.com/terraform/{{.SemVer}}/terraform_{{.SemVer}}_darwin_amd64.zip",
		"https://releases.hashicorp.com/terraform/{{.SemVer}}/terraform_{{.SemVer}}_windows_amd64.zip",
	},
}

// Asset represents a GitHub release asset in a repository.
type Asset struct {
	downloadURL url.URL
	mime        mime.MIME
}

// AssetList is a list of GitHub releaset asset.
type AssetList []Asset

// newAsset returns a new GitHub release asset object.
func newAsset(downloadURL url.URL, mime mime.MIME) Asset {
	return Asset{
		downloadURL: downloadURL,
		mime:        mime,
	}
}

// os returns an os detected by url to download asset.
func (a Asset) os() platform.OS {
	os, _ := platform.Detect(string(a.downloadURL))
	return os
}

// arch returns an arch detected by url to download assetl.
func (a Asset) arch() platform.Arch {
	_, arch := platform.Detect(string(a.downloadURL))
	return arch
}

// hasExecutableBinary returns true if asset may have executable binary.
func (a Asset) hasExecutableBinary() bool {
	return a.mime.IsArchived() || a.mime.IsCompressed() || a.mime.IsOctetStream()
}

// find a GitHub release asset which has executable binary and whose os/arch are same as supplied value.
func (s AssetList) find(os platform.OS, arch platform.Arch) (Asset, error) {
	index := slices.IndexFunc(s, func(asset Asset) bool {
		return asset.os() == os && asset.arch() == arch && asset.hasExecutableBinary()
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

// get returns a new GitHub release asset object.
func (r *AssetRepository) get(downloadURL url.URL) (Asset, error) {
	resp, err := http.Get(string(downloadURL))
	if err != nil {
		return Asset{}, err
	}
	defer resp.Body.Close()

	mime, err := mime.DetectReader(resp.Body)
	if err != nil {
		return Asset{}, err
	}

	return newAsset(downloadURL, mime), nil
}

// list returns a list of GitHub release asset in a GitHub release.
func (r *AssetRepository) list(ctx context.Context, repo Repository, release Release) (AssetList, error) {
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
	assets := AssetList{}
	for _, githubAsset := range githubAssets {
		downloadURL := url.URL(githubAsset.GetBrowserDownloadURL())
		asset, err := r.get(downloadURL)
		if err != nil {
			return nil, err
		}
		assets = append(assets, asset)
	}

	return assets, nil
}

// listExternal returns a list of release assets hosted on server outside of GitHub.
func (r *AssetRepository) listExternal(repo Repository, release Release) (AssetList, error) {
	tmpls, ok := assetDownloadURLTemplatesHostedOnOutside[repo.fullName()]
	if !ok {
		return AssetList{}, nil
	}
	assets := AssetList{}
	for _, tmpl := range tmpls {
		downloadURL, err := applyDownloadURLTemplateToRelease(tmpl, release)
		if err != nil {
			return nil, err
		}
		asset, err := r.get(downloadURL)
		if err != nil {
			return nil, err
		}
		assets = append(assets, asset)
	}
	return assets, nil
}

// applyDownloadURLTemplateToRelease applies a download url template to the release object, and return it as download url.
func applyDownloadURLTemplateToRelease(downloadURL url.Template, release Release) (url.URL, error) {
	data := struct {
		Tag    string
		SemVer string
	}{
		Tag:    release.tag,
		SemVer: release.semver(),
	}
	return downloadURL.Execute(data)
}
