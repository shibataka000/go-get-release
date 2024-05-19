package github

import (
	"context"
	"slices"

	"github.com/google/go-github/v48/github"
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

// newAssetMeat return new AssetMeta object.
func newAssetMeta(downloadURL url.URL, os platform.OS, arch platform.Arch) AssetMeta {
	return AssetMeta{
		DownloadURL: downloadURL,
		OS:          os,
		Arch:        arch,
	}
}

// AssetMetaList is a list of AssetMeta.
type AssetMetaList []AssetMeta

// find AssetMeta by OS/Arch.
func (assets AssetMetaList) find(goos platform.OS, goarch platform.Arch) (AssetMeta, error) {
	i := slices.IndexFunc(assets, func(a AssetMeta) bool {
		return a.OS == goos && a.Arch == goarch
	})
	if i == -1 {
		return AssetMeta{}, &AssetNotFoundError{}
	}
	return assets[i], nil
}

// AssetRepository is repository for Asset.
type AssetRepository struct {
	client *github.Client
}

// NewAssetRepository return new AssetRepository object.
func NewAssetRepository(ctx context.Context, token string) *AssetRepository {
	return &AssetRepository{
		client: newGitHubClient(ctx, token),
	}
}

// listFromAPI return a listFromAPI of AssetMeta in a GitHub release.
func (r *AssetRepository) listFromAPI(ctx context.Context, repo Repository, release Release) (AssetMetaList, error) {
	githubRelease, _, err := r.client.Repositories.GetReleaseByTag(ctx, repo.Owner, repo.Name, release.Tag)
	if err != nil {
		return nil, err
	}

	result := AssetMetaList{}
	for page := 1; page != 0; {
		assets, resp, err := r.client.Repositories.ListReleaseAssets(ctx, repo.Owner, repo.Name, *githubRelease.ID, &github.ListOptions{
			Page: page,
		})
		if err != nil {
			return result, err
		}
		for _, asset := range assets {
			downloadURL := url.URL(asset.GetBrowserDownloadURL())
			goos, goarch := platform.Guess(downloadURL.Base())
			result = append(result, newAssetMeta(downloadURL, goos, goarch))
		}
		page = resp.NextPage
	}
	return result, nil
}

// listFromBuiltIn return a list of AssetMeta in `builtin.yaml`.
func (r *AssetRepository) listFromBuiltIn(repo Repository, release Release) (AssetMetaList, error) {
	// Find AssetMeta from builtin.
	type Record struct {
		Repository Repository
		Assets     AssetMetaList
	}
	records := []Record{}
	err := yaml.Unmarshal(builtin, &records)
	if err != nil {
		return nil, err
	}
	index := slices.IndexFunc(records, func(r Record) bool {
		return r.Repository.Owner == repo.Owner && r.Repository.Name == repo.Name && r.Assets != nil
	})
	if index == -1 {
		return nil, &AssetNotFoundError{}
	}

	// Fill download URL template.
	assets := records[index].Assets
	param := struct {
		Tag    string
		SemVer string
	}{
		Tag:    release.Tag,
		SemVer: release.semver(),
	}
	for i, asset := range assets {
		assets[i].DownloadURL, err = url.Template(asset.DownloadURL).Execute(param)
		if err != nil {
			return nil, err
		}
	}
	return assets, nil
}
