package github

import (
	"bytes"
	"context"
	"net/http"
	"net/url"
	"slices"
	"text/template"

	"github.com/google/go-github/v48/github"
	"github.com/shibataka000/go-get-release/mime"
	"github.com/shibataka000/go-get-release/platform"
	"golang.org/x/oauth2"
)

// externalAssets is a map of repository and a list of GitHub release asset template on server outside GitHub.
var externalAssets = map[Repository]AssetTemplateList{
	newRepository("hashicorp", "terraform"): {
		newAssetTemplate(newTemplate("", "https://releases.hashicorp.com/terraform/{{.SemVer}}/terraform_{{.SemVer}}_linux_amd64.zip"), ""),
		newAssetTemplate(newTemplate("", "https://releases.hashicorp.com/terraform/{{.SemVer}}/terraform_{{.SemVer}}_darwin_amd64.zip"), ""),
		newAssetTemplate(newTemplate("", "https://releases.hashicorp.com/terraform/{{.SemVer}}/terraform_{{.SemVer}}_windows_amd64.zip"), ""),
	},
}

// Asset represents a GitHub release asset in a repository.
type Asset struct {
	downloadURL *url.URL
	mime        mime.MIME
}

// AssetList is a list of GitHub release asset in a repository.
type AssetList []Asset

// AssetTemplate is a template of GitHub release asset in a repository.
type AssetTemplate struct {
	downloadURL *template.Template
	mime        mime.MIME
}

// AssetTemplateList is a list of template of Github release asset in a repository.
type AssetTemplateList []AssetTemplate

// newAsset returns a new GitHub release asset object.
func newAsset(downloadURL *url.URL, mime mime.MIME) Asset {
	return Asset{
		downloadURL: downloadURL,
		mime:        mime,
	}
}

// newAssetTemplate returns a new GitHub release asset template object.
func newAssetTemplate(downloadURL *template.Template, mime mime.MIME) AssetTemplate {
	return AssetTemplate{
		downloadURL: downloadURL,
		mime:        mime,
	}
}

// os returns an os detected by url to download GitHub release asset.
func (a Asset) os() platform.OS {
	os, _ := platform.Detect(a.downloadURL.String())
	return os
}

// arch returns an arch detected by url to download GitHub release asset.
func (a Asset) arch() platform.Arch {
	_, arch := platform.Detect(a.downloadURL.String())
	return arch
}

// mayHaveExecutableBinary returns true if GitHub release asset may have executable binary.
func (a Asset) mayHaveExecutableBinary() bool {
	return a.mime.IsArchived() || a.mime.IsCompressed() || a.mime.IsOctetStream()
}

// find a GitHub release asset which has executable binary for specified platform.
func (s AssetList) find(os platform.OS, arch platform.Arch) (Asset, error) {
	index := slices.IndexFunc(s, func(asset Asset) bool {
		return asset.os() == os && asset.arch() == arch && asset.mayHaveExecutableBinary()
	})
	if index == -1 {
		return Asset{}, &AssetNotFoundError{}
	}
	return s[index], nil
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
	return newAsset(downloadURL, a.mime), nil
}

// execute applies a list of asset template to the GitHub release object, and returns them as list of GitHub release asset.
func (s AssetTemplateList) execute(release Release) (AssetList, error) {
	assets := AssetList{}
	for _, a := range s {
		asset, err := a.execute(release)
		if err != nil {
			return nil, err
		}
		assets = append(assets, asset)
	}
	return assets, nil
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

// get returns a new GitHub release asset.
func (r *AssetRepository) get(downloadURL *url.URL) (Asset, error) {
	resp, err := http.Get(downloadURL.String())
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

// list returns a list of GitHub release asset in GitHub.
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
		downloadURL, err := url.Parse(githubAsset.GetBrowserDownloadURL())
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

// ExternalAssetRepository is a repository for a GitHub release asset on server outside GitHub.
type ExternalAssetRepository struct {
	assets map[Repository]AssetTemplateList
}

// NewExternalAssetRepository returns a new ExternalAssetRepository object.
func NewExternalAssetRepository() *ExternalAssetRepository {
	return &ExternalAssetRepository{
		assets: externalAssets,
	}
}

// list returns a list of GitHub release asset on server outside GitHub.
func (r *ExternalAssetRepository) list(repo Repository, release Release) (AssetList, error) {
	tmpls, ok := r.assets[repo]
	if !ok {
		return AssetList{}, nil
	}
	return tmpls.execute(release)
}
