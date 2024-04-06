package github

import (
	"context"
)

// ApplicationService.
type ApplicationService struct {
	repository *InfrastructureRepository
}

// NewApplicationService return new application service instance.
func NewApplicationService(repository *InfrastructureRepository) *ApplicationService {
	return &ApplicationService{
		repository: repository,
	}
}

// Search package.
func (a *ApplicationService) Search(ctx context.Context, query SearchQuery, platform Platform) (Package, error) {
	var err error

	var ghRepo GitHubRepository
	if query.HasOwner() {
		ghRepo, err = a.repository.FindGitHubRepository(ctx, query.Repository.Owner, query.Repository.Name)
	} else {
		ghRepo, err = a.repository.SearchGitHubRepository(ctx, query.Repository.Name)
	}
	if err != nil {
		return Package{}, err
	}
	repo := a.factory.NewRepository(ghRepo)

	var ghRelease GitHubRelease
	if query.HasTag() {
		ghRelease, err = a.repository.FindGitHubReleaseByTag(ctx, ghRepo, query.Tag)
	} else {
		ghRelease, err = a.repository.LatestGitHubRelease(ctx, ghRepo)
	}
	if err != nil {
		return Package{}, err
	}
	release := a.factory.NewRelease(ghRelease)

	index, err := a.repository.LoadBuiltInIndex()
	if err != nil {
		return Package{}, err
	}

	var asset Asset
	if index.HasAsset(repo, platform) {
		assetInIndex, err := index.FindAsset(repo, platform)
		if err != nil {
			return Package{}, err
		}
		asset, err = a.factory.NewAssetFromIndex(assetInIndex, release)
		if err != nil {
			return Package{}, err
		}
	} else {
		ghAssets, err := a.repository.ListGitHubAssets(ctx, ghRepo, ghRelease)
		if err != nil {
			return Package{}, err
		}
		asset, err = a.factory.NewAssetFromGitHub(ghAssets, platform)
		if err != nil {
			return Package{}, err
		}
	}

	var execBinary ExecBinary
	if index.HasExecBinary(repo) {
		execBinaryInIndex, err := index.FindExecBinary(repo)
		if err != nil {
			return Package{}, err
		}
		execBinary = a.factory.NewExecBinaryFromIndex(execBinaryInIndex, platform)
	} else {
		execBinary = a.factory.NewExecBinaryFromGitHub(ghRepo, platform)
	}

	return New(repo, release, asset, execBinary), nil
}
