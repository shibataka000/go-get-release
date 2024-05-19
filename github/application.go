package github

import (
	"context"

	"github.com/shibataka000/go-get-release/platform"
)

// ApplicationService.
type ApplicationService struct {
	repository *RepositoryRepository
	release    *ReleaseRepository
	asset      *AssetRepository
}

// NewApplicationService return new ApplicationService object.
func NewApplicationService(repository *RepositoryRepository, release *ReleaseRepository, asset *AssetRepository) *ApplicationService {
	return &ApplicationService{
		repository: repository,
		release:    release,
		asset:      asset,
	}
}

// SearchRepository return a GitHub repository searched by query.
func (a *ApplicationService) SearchRepository(ctx context.Context, query string) (Repository, error) {
	return a.repository.search(ctx, query)
}

// LatestRelease return a latest GitHub release in a repository.
func (a *ApplicationService) LatestRelease(ctx context.Context, repo Repository) (Release, error) {
	return a.release.latest(ctx, repo)
}

// FindAsset return a GitHub release asset in a repository whose OS/Arch are same to passed values.
func (a *ApplicationService) FindAsset(ctx context.Context, repo Repository, release Release, os platform.OS, arch platform.Arch) (AssetMeta, error) {
	assets1, err := a.asset.listFromAPI(ctx, repo, release)
	if err != nil {
		return AssetMeta{}, err
	}
	assets2, err := a.asset.listFromBuiltIn(repo, release)
	if err != nil {
		return AssetMeta{}, err
	}
	assets := append(assets1, assets2...)
	return assets.find(os, arch)
}

// GetExecutableBinaryMeta return executable binary metadata from GitHub repository.
func (a *ApplicationService) GetExecutableBinaryMeta(repo Repository) ExecutableBinaryMeta {
	return newExecutableBinaryMetaFromRepository(repo)
}