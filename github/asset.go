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
		newAssetTemplate(newTemplate("", "https://releases.hashicorp.com/terraform/{{.SemVer}}/terraform_{{.SemVer}}_linux_amd64.zip")),
		newAssetTemplate(newTemplate("", "https://releases.hashicorp.com/terraform/{{.SemVer}}/terraform_{{.SemVer}}_darwin_amd64.zip")),
		newAssetTemplate(newTemplate("", "https://releases.hashicorp.com/terraform/{{.SemVer}}/terraform_{{.SemVer}}_windows_amd64.zip")),
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
func newAssetTemplate(downloadURL *template.Template) AssetTemplate {
	return AssetTemplate{
		downloadURL: downloadURL,
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

// downloadURLWithRelease applies an asset template to the GitHub release object, and returns it as GitHub release asset.
func (a AssetTemplate) downloadURLWithRelease(release Release) (*url.URL, error) {
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
		return nil, err
	}
	return url.Parse(buf.String())
}

// AssetRepository is a repository for a GitHub release asset.
type AssetRepository struct {
	client         *github.Client
	externalAssets map[Repository]AssetTemplateList
}

// NewAssetRepository returns a new AssetRepository object.
func NewAssetRepository(ctx context.Context, token string) *AssetRepository {
	var httpClient *http.Client
	if token != "" {
		tokenSource := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})
		httpClient = oauth2.NewClient(ctx, tokenSource)
	}
	return &AssetRepository{
		client:         github.NewClient(httpClient),
		externalAssets: externalAssets,
	}
}

// list returns a list of GitHub release asset.
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

	downloadURLs := []*url.URL{}

	// List URL to download GitHub release asset.
	for _, githubAsset := range githubAssets {
		downloadURL, err := url.Parse(githubAsset.GetBrowserDownloadURL())
		if err != nil {
			return nil, err
		}
		downloadURLs = append(downloadURLs, downloadURL)
	}

	// List URL to download GitHub release asset on server outside GitHub.
	if externalAssets, ok := r.externalAssets[repo]; ok {
		for _, externalAsset := range externalAssets {
			downloadURL, err := externalAsset.downloadURLWithRelease(release)
			if err != nil {
				return nil, err
			}
			downloadURLs = append(downloadURLs, downloadURL)
		}
	}

	// Detect MIME.
	assets := AssetList{}
	for _, downloadURL := range downloadURLs {
		resp, err := http.Get(downloadURL.String())
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()
		mime, err := mime.DetectReader(resp.Body)
		if err != nil {
			return nil, err
		}
		assets = append(assets, newAsset(downloadURL, mime))
	}

	return assets, nil
}
