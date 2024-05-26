package github

import (
	"context"
	"slices"

	"github.com/google/go-github/v48/github"
	"github.com/shibataka000/go-get-release/file"
	"github.com/shibataka000/go-get-release/platform"
	"github.com/shibataka000/go-get-release/url"
	"gopkg.in/yaml.v3"
)

// AssetMeta represents a GitHub release asset in a repository.
type AssetMeta struct {
	DownloadURL url.URL       `yaml:"downloadURL"`
	OS          platform.OS   `yaml:"os"`
	Arch        platform.Arch `yaml:"arch"`
}

// newAssetMeta return new AssetMeta object.
func newAssetMeta(downloadURL url.URL, os platform.OS, arch platform.Arch) AssetMeta {
	return AssetMeta{
		DownloadURL: downloadURL,
		OS:          os,
		Arch:        arch,
	}
}

// fileName return file name of GitHub release asset.
func (a AssetMeta) fileName() file.Name {
	return file.Name(a.DownloadURL.Base())
}

// hasExecutableBinary return true if AssetMeta may have executable binary.
func (a AssetMeta) hasExecutableBinary() bool {
	exts := []string{"", ".exe", ".linux", ".darwin", ".linux-amd64", ".darwin-amd64", ".amd64", ".gz", ".xz", ".zip", ".tar"}
	return slices.Contains(exts, a.fileName().Ext())
}

// AssetMetaList is a list of GitHub releaset asset.
type AssetMetaList []AssetMeta

// find AssetMeta which has executable binary and whose OS/Arch are same as passed value.
func (assets AssetMetaList) find(os platform.OS, arch platform.Arch) (AssetMeta, error) {
	index := slices.IndexFunc(assets, func(asset AssetMeta) bool {
		return asset.hasExecutableBinary() && asset.OS == os && asset.Arch == arch
	})
	if index == -1 {
		return AssetMeta{}, &AssetNotFoundError{}
	}
	return assets[index], nil
}

// AssetRepository is repository for a GitHub release asset.
type AssetRepository struct {
	client *github.Client
}

// NewAssetRepository return new AssetRepository object.
func NewAssetRepository(ctx context.Context, token string) *AssetRepository {
	return &AssetRepository{
		client: newGitHubClient(ctx, token),
	}
}

// listFromAPI return a list of AssetMeta in a GitHub release using GitHub API.
func (r *AssetRepository) listFromAPI(ctx context.Context, repo Repository, release Release) (AssetMetaList, error) {
	// Get GitHub release ID.
	githubRelease, _, err := r.client.Repositories.GetReleaseByTag(ctx, repo.Owner, repo.Name, release.Tag)
	if err != nil {
		return nil, err
	}
	releaseID := *githubRelease.ID

	// List GitHub release assets.
	githubAssets := []*github.ReleaseAsset{}
	for page := 1; page != 0; {
		assets, resp, err := r.client.Repositories.ListReleaseAssets(ctx, repo.Owner, repo.Name, releaseID, &github.ListOptions{
			Page: page,
		})
		if err != nil {
			return nil, err
		}
		githubAssets = append(githubAssets, assets...)
		page = resp.NextPage
	}

	// Create AssetMetaList from list of GitHub release asset.
	result := AssetMetaList{}
	for _, githubAsset := range githubAssets {
		downloadURL := url.URL(githubAsset.GetBrowserDownloadURL())
		os, arch := platform.Guess(downloadURL.Base())
		result = append(result, newAssetMeta(downloadURL, os, arch))
	}

	return result, nil
}

// listFromBuiltIn return a list of AssetMeta from built-in data.
func (r *AssetRepository) listFromBuiltIn(repo Repository, release Release) (AssetMetaList, error) {
	// Define structure about record in built-in data to unmarshal them.
	type Record struct {
		Repository Repository    `yaml:"repository"`
		Assets     AssetMetaList `yaml:"assets"`
	}

	// Unmarshal built-in data.
	records := []Record{}
	err := yaml.Unmarshal(builtin, &records)
	if err != nil {
		return nil, err
	}

	// Find record in built-in data by repository.
	index := slices.IndexFunc(records, func(r Record) bool {
		return r.Repository.Owner == repo.Owner && r.Repository.Name == repo.Name && r.Assets != nil
	})
	if index == -1 {
		return nil, &AssetNotFoundError{}
	}
	record := records[index]

	// Apply a download URL template to the GitHub release object.
	result := AssetMetaList{}
	param := struct {
		Tag    string
		SemVer string
	}{
		Tag:    release.Tag,
		SemVer: release.semver(),
	}
	for _, asset := range record.Assets {
		downloadURLTemplate := url.Template(asset.DownloadURL)
		downloadURL, err := downloadURLTemplate.Execute(param)
		if err != nil {
			return nil, err
		}
		result = append(result, newAssetMeta(downloadURL, asset.OS, asset.Arch))
	}

	return result, nil
}
