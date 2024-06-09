package github

import (
	"context"
	"net/http"
	"slices"

	"github.com/google/go-github/v48/github"
	"github.com/shibataka000/go-get-release/file"
	"github.com/shibataka000/go-get-release/mime"
	"github.com/shibataka000/go-get-release/platform"
	"github.com/shibataka000/go-get-release/url"
	"golang.org/x/oauth2"
)

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

// name returns an asset file name.
func (a Asset) name() file.Name {
	return file.Name(a.downloadURL.Base())
}

// os returns an os detected by asset file name.
func (a Asset) os() platform.OS {
	os, _ := platform.Detect(string(a.name()))
	return os
}

// arch returns an arch detected by asset file name.
func (a Asset) arch() platform.Arch {
	_, arch := platform.Detect(string(a.name()))
	return arch
}

// hasExecutableBinary returns true if asset may have executable binary.
func (a Asset) hasExecutableBinary() bool {
	return a.mime.IsArchived() || a.mime.IsCompressed() || a.mime.IsOctetStream()
}

// find a GitHub release asset which has executable binary and whose os/arch are same as supplied value.
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

func (r *AssetRepository) get(downloadURL url.URL) (Asset, error) {

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
	result := AssetList{}
	for _, githubAsset := range githubAssets {
		asset, err := r.get(url.URL(githubAsset.GetBrowserDownloadURL()))
		if err != nil {
			return nil, err
		}
		result = append(result, asset)
	}

	return result, nil
}

var externalAssets = map[string][]url.Template{
	"hashicorp/terraform": {
		"https://releases.hashicorp.com/terraform/{{.SemVer}}/terraform_{{.SemVer}}_linux_amd64.zip",
		"https://releases.hashicorp.com/terraform/{{.SemVer}}/terraform_{{.SemVer}}_darwin_amd64.zip",
		"https://releases.hashicorp.com/terraform/{{.SemVer}}/terraform_{{.SemVer}}_windows_amd64.zip",
	},
}

// listExternalAssets returns a list of release assets hosted on server out of GitHub.
func (r *AssetRepository) listExternalAssets(repo Repository, release Release) (AssetList, error) {
	templates, ok := externalAssets[repo.fullName()]
	if !ok {
		return AssetList{}, nil
	}
	assets := AssetList{}
	for _, tmpl := range templates {
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
