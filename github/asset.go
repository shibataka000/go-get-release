package github

import (
	"bytes"
	"context"
	"net/http"
	"net/url"
	"regexp"
	"slices"
	"text/template"

	"github.com/google/go-github/v62/github"
	dist "github.com/shibataka000/go-get-release/distribution"
	"github.com/shibataka000/go-get-release/mime"
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

// newAssetFromString returns a new GitHub release asset object.
// Given download URL must be able to be parsed as URL.
func newAssetFromString(downloadURL string) (Asset, error) {
	url, err := url.Parse(downloadURL)
	if err != nil {
		return Asset{}, err
	}
	return newAsset(url), nil
}

// mustNewAssetFromString returns a new GitHub release asset object.
// Given download URL must be able to be parsed as URL.
// This gets into a panic if the error is non-nil.
func mustNewAssetFromString(downloadURL string) Asset {
	asset, err := newAssetFromString(downloadURL)
	if err != nil {
		panic(err)
	}
	return asset
}

// os returns an os detected by url to download GitHub release asset.
func (a Asset) os() dist.OS {
	os, _ := dist.Detect(a.DownloadURL.String())
	return os
}

// arch returns an arch detected by url to download GitHub release asset.
func (a Asset) arch() dist.Arch {
	_, arch := dist.Detect(a.DownloadURL.String())
	if arch == "" {
		return "amd64"
	}
	return arch
}

// mime returns a mime type.
func (a Asset) mime() mime.Type {
	return mime.Detect(a.DownloadURL.String())
}

// hasExecBinary returns true if GitHub release asset may have executable binary.
func (a Asset) hasExecBinary() bool {
	return a.mime().IsCompressed() || a.mime().IsOctetStream()
}

// isIgnored returns true if this GitHub release asset should be ignored.
func (a Asset) isIgnored() bool {
	return ignoredAssets.matchAny(a)
}

func (a Asset) execBinary() ExecBinary {
	return newExecBinaryWithOS("", a.os())
}

// AssetList is a list of GitHub release asset.
type AssetList []Asset

// find a GitHub release asset which has executable binary for specified os and arch.
func (s AssetList) find(os dist.OS, arch dist.Arch) (Asset, error) {
	index := slices.IndexFunc(s, func(asset Asset) bool {
		return asset.os() == os && asset.arch() == arch && asset.hasExecBinary() && !asset.isIgnored()
	})
	if index == -1 {
		return Asset{}, &AssetNotFoundError{}
	}
	return s[index], nil
}

// AssetTemplate is a template of GitHub release asset.
type AssetTemplate struct {
	downloadURL *template.Template
}

// newAssetTemplate returns a new GitHub release asset template object.
func newAssetTemplate(downloadURL *template.Template) AssetTemplate {
	return AssetTemplate{
		downloadURL: downloadURL,
	}
}

// newAssetTemplateFromString returns a new GitHub release asset template object.
// Given download URL must be able to parsed as template.
func newAssetTemplateFromString(downloadURL string) (AssetTemplate, error) {
	tmpl, err := template.New("").Parse(downloadURL)
	if err != nil {
		return AssetTemplate{}, err
	}
	return newAssetTemplate(tmpl), nil
}

// mustNewAssetTemplateFromString returns a new GitHub release asset template object.
// Given download URL must be able to parsed as template.
// This gets into a panic if the error is non-nil.
func mustNewAssetTemplateFromString(downloadURL string) AssetTemplate {
	asset, err := newAssetTemplateFromString(downloadURL)
	if err != nil {
		panic(err)
	}
	return asset
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
	return newAssetFromString(buf.String())
}

// AssetTemplateList is a list of template of Github release asset.
type AssetTemplateList []AssetTemplate

// execute applies a list of asset template to the GitHub release object, and returns it as a list of GitHub release asset.
func (s AssetTemplateList) execute(release Release) (AssetList, error) {
	assets := AssetList{}
	for _, tmpl := range s {
		asset, err := tmpl.execute(release)
		if err != nil {
			return nil, err
		}
		assets = append(assets, asset)
	}
	return assets, nil
}

// AssetRegexp is a regular expression about GitHub release asset.
type AssetRegexp struct {
	downloadURL *regexp.Regexp
}

// newAssetRegexp returns a new regular expression object about GitHub release asset.
func newAssetRegexp(downloadURL *regexp.Regexp) AssetRegexp {
	return AssetRegexp{
		downloadURL: downloadURL,
	}
}

// newAssetRegexpFromString returns a new regular expression object about GitHub release asset.
// Given download URL must be able to be compiled as regular expression.
func newAssetRegexpFromString(downloadURL string) (AssetRegexp, error) {
	re, err := regexp.Compile(downloadURL)
	if err != nil {
		return AssetRegexp{}, err
	}
	return newAssetRegexp(re), nil
}

// mustNewAssetRegexpFromString returns a new regular expression object about GitHub release asset.
// Given download URL must be able to be compiled as regular expression.
// This gets into a panic if the error is non-nil.
func mustNewAssetRegexpFromString(downloadURL string) AssetRegexp {
	asset, err := newAssetRegexpFromString(downloadURL)
	if err != nil {
		panic(err)
	}
	return asset
}

// match returns true if a GitHub release asset download URL contains any match of the regular expression.
func (a AssetRegexp) match(asset Asset) bool {
	return a.downloadURL.MatchString(asset.DownloadURL.String())
}

// AssetRegexpList is a list of regular expression about GitHub release asset.
type AssetRegexpList []AssetRegexp

// matchAny returns true if any of GitHub release asset contains any match of the regular expression.
func (s AssetRegexpList) matchAny(asset Asset) bool {
	return slices.IndexFunc(s, func(a AssetRegexp) bool {
		return a.match(asset)
	}) != -1
}

// AssetCache is a cache for GitHub release asset.
type AssetCache map[Repository]map[Release]AssetList

// newAssetCache returns a new cache object for GitHub release asset.
func newAssetCache() AssetCache {
	return AssetCache{}
}

// set stores cache with a list of GitHub release asset.
func (c AssetCache) set(repo Repository, release Release, assets AssetList) {
	if _, ok := c[repo]; !ok {
		c[repo] = map[Release]AssetList{}
	}
	c[repo][release] = assets
}

// get returns a list GitHub release asset value in cache.
func (c AssetCache) get(repo Repository, release Release) (AssetList, bool) {
	if _, ok := c[repo]; ok {
		if v, ok := c[repo][release]; ok {
			return v, true
		}
	}
	return nil, false
}

// AssetRepository is a repository for a GitHub release asset.
type AssetRepository struct {
	client *github.Client
	cache  AssetCache
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
		cache:  newAssetCache(),
	}
}

// list returns a list of GitHub release asset.
func (r *AssetRepository) list(ctx context.Context, repo Repository, release Release) (AssetList, error) {
	// Return GitHub release assets in cache if exists.
	if v, ok := r.cache.get(repo, release); ok {
		return v, nil
	}

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
			asset, err := newAssetFromString(downloadURL)
			if err != nil {
				return nil, err
			}
			assets = append(assets, asset)
		}
		page = resp.NextPage
	}

	// List GitHub release assets on external server.
	if tmpls, ok := externalAssets[repo]; ok {
		applied, err := tmpls.execute(release)
		if err != nil {
			return nil, err
		}
		assets = append(assets, applied...)
	}

	// Store cache with results.
	r.cache.set(repo, release, assets)

	return assets, nil
}
