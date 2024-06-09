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

// externals is a map of repository and url template list to download asset on server outside GitHub.
var externals = map[Repository][]AssetTemplate{
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

// AssetTemplate is a template of GitHub release asset in a repository.
type AssetTemplate struct {
	downloadURL *template.Template
}

// AssetList is a list of GitHub releaset asset in a repository.
type AssetList []Asset

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

// os returns an os detected by url to download asset.
func (a Asset) os() platform.OS {
	os, _ := platform.Detect(a.downloadURL.String())
	return os
}

// arch returns an arch detected by url to download asset.
func (a Asset) arch() platform.Arch {
	_, arch := platform.Detect(a.downloadURL.String())
	return arch
}

// mayHaveExecutableBinary returns true if asset may have executable binary.
func (a Asset) mayHaveExecutableBinary() bool {
	return a.mime.IsArchived() || a.mime.IsCompressed() || a.mime.IsOctetStream()
}

// applyDownloadURLTemplateToRelease applies a download url template to the release object, and return it as download url.
func (a AssetTemplate) getDownloadURL(release Release) (*url.URL, error) {
	buf := new(bytes.Buffer)
	data := struct {
		Tag    string
		SemVer string
	}{
		Tag:    release.tag,
		SemVer: release.semver(),
	}
	if err := a.downloadURL.Execute(buf, data); err != nil {
		return nil, err
	}
	return url.Parse(buf.String())
}

// find a GitHub release asset which has executable binary and whose os/arch are same as supplied value.
func (s AssetList) find(os platform.OS, arch platform.Arch) (Asset, error) {
	index := slices.IndexFunc(s, func(asset Asset) bool {
		return asset.os() == os && asset.arch() == arch && asset.mayHaveExecutableBinary()
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

// listExternal returns a list of release assets hosted on server outside of GitHub.
func (r *AssetRepository) listExternal(repo Repository, release Release) (AssetList, error) {
	tmpls, ok := externals[repo]
	if !ok {
		return AssetList{}, nil
	}
	assets := AssetList{}
	for _, tmpl := range tmpls {
		downloadURL, err := tmpl.getDownloadURL(release)
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
