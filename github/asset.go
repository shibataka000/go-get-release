package github

import (
	"context"

	"github.com/google/go-github/v48/github"
	"github.com/shibataka000/go-get-release/platform"
	"github.com/shibataka000/go-get-release/slices"
	"github.com/shibataka000/go-get-release/url"
	"gopkg.in/yaml.v3"
)

// AssetMeta represents a GitHub release asset in a repository.
type AssetMeta struct {
	DownloadURL url.URL       `yaml:"downloadURL"`
	OS          platform.OS   `yaml:"os"`
	Arch        platform.Arch `yaml:"arch"`
}

// hasExecutableBinary return true if AssetMeta has executable binary.
func (a AssetMeta) hasExecutableBinary() bool {
	return true
}

// AssetMetaList is a list of GitHub releaset asset.
type AssetMetaList []AssetMeta

// find AssetMeta which has executable binary and whose OS/Arch are same as passed value.
func (assets AssetMetaList) find(os platform.OS, arch platform.Arch) (AssetMeta, error) {
	asset, err := slices.Find(assets, func(asset AssetMeta) bool {
		return asset.OS == os && asset.Arch == arch && asset.hasExecutableBinary()
	})
	if err != nil {
		return AssetMeta{}, &AssetNotFoundError{}
	}
	return asset, nil
}

// AssetRepository is repository for a GitHub release asset.
type AssetRepository struct {
	client  *github.Client
	factory *AssetFactory
}

// NewAssetRepository return new AssetRepository object.
func NewAssetRepository(ctx context.Context, token string, factory *AssetFactory) *AssetRepository {
	return &AssetRepository{
		client:  newGitHubClient(ctx, token),
		factory: factory,
	}
}

// listFromAPI return a list of AssetMeta in a GitHub release.
func (r *AssetRepository) listFromAPI(ctx context.Context, repo Repository, release Release) (AssetMetaList, error) {
	githubRelease, _, err := r.client.Repositories.GetReleaseByTag(ctx, repo.Owner, repo.Name, release.Tag)
	if err != nil {
		return nil, err
	}
	githubAssets := []*github.ReleaseAsset{}
	for page := 1; page != 0; {
		assets, resp, err := r.client.Repositories.ListReleaseAssets(ctx, repo.Owner, repo.Name, *githubRelease.ID, &github.ListOptions{
			Page: page,
		})
		if err != nil {
			return nil, err
		}
		githubAssets = append(githubAssets, assets...)
		page = resp.NextPage
	}
	return slices.Map[AssetMetaList](githubAssets, func(asset *github.ReleaseAsset) AssetMeta {
		return r.factory.newMetaFromDownloadURL(url.URL(asset.GetBrowserDownloadURL()))
	}), nil
}

// listFromBuiltIn return a list of AssetMeta in `builtin.yaml`.
func (r *AssetRepository) listFromBuiltIn(repo Repository, release Release) (AssetMetaList, error) {
	type Record struct {
		Repository Repository    `yaml:"repository"`
		Assets     AssetMetaList `yaml:"assets"`
	}
	records := []Record{}
	err := yaml.Unmarshal(builtin, &records)
	if err != nil {
		return nil, err
	}
	record, err := slices.Find(records, func(r Record) bool {
		return r.Repository.Owner == repo.Owner && r.Repository.Name == repo.Name && r.Assets != nil
	})
	if err != nil {
		return nil, &AssetNotFoundError{}
	}
	return slices.MapE[AssetMetaList](record.Assets, func(asset AssetMeta) (AssetMeta, error) {
		return r.factory.newMetaFromDownloadURLTemplate(url.Template(asset.DownloadURL), asset.OS, asset.Arch, release)
	})
}

// AssetFactory is factory to create a new GitHub release asset object.
type AssetFactory struct{}

// NewAssetFactory return new AssetFactory object.
func NewAssetFactory() *AssetFactory {
	return &AssetFactory{}
}

// newMeta return new AssetMeta object.
func (f *AssetFactory) newMeta(downloadURL url.URL, os platform.OS, arch platform.Arch) AssetMeta {
	return AssetMeta{
		DownloadURL: downloadURL,
		OS:          os,
		Arch:        arch,
	}
}

func (f *AssetFactory) newMetaFromDownloadURL(downloadURL url.URL) AssetMeta {
	os, arch := platform.Guess(downloadURL.Base())
	return f.newMeta(downloadURL, os, arch)
}

func (f *AssetFactory) newMetaFromDownloadURLTemplate(downloadURL url.Template, os platform.OS, arch platform.Arch, release Release) (AssetMeta, error) {
	param := struct {
		Tag    string
		SemVer string
	}{
		Tag:    release.Tag,
		SemVer: release.semver(),
	}
	url, err := downloadURL.Execute(param)
	if err != nil {
		return AssetMeta{}, err
	}
	return f.newMeta(url, os, arch), nil
}
