package github

import (
	"context"

	"github.com/shibataka000/go-get-release/platform"
)

// ApplicationService.
type ApplicationService struct {
	asset *AssetRepository
}

// NewApplicationService return new ApplicationService object.
func NewApplicationService(asset *AssetRepository) *ApplicationService {
	return &ApplicationService{
		asset: asset,
	}
}

// FindAsset return a GitHub release asset in a repository whose OS/Arch are same to passed values.
func (a *ApplicationService) FindAssetMeta(ctx context.Context, repoNameWithOwner string, tag string, os platform.OS, arch platform.Arch) (AssetMeta, error) {
	repo, err := newRepositoryFromFullName(repoNameWithOwner)
	if err != nil {
		return AssetMeta{}, err
	}
	release := newRelease(tag)
	if asset, err := a.findAssetMetaFromBuiltIn(repo, release, os, arch); err == nil {
		return asset, nil
	}
	return a.findAssetMetaFromAPI(ctx, repo, release, os, arch)
}

func (a *ApplicationService) findAssetMetaFromAPI(ctx context.Context, repo Repository, release Release, os platform.OS, arch platform.Arch) (AssetMeta, error) {
	assets, err := a.asset.listFromAPI(ctx, repo, release)
	if err != nil {
		return AssetMeta{}, err
	}
	return assets.find(os, arch)
}

func (a *ApplicationService) findAssetMetaFromBuiltIn(repo Repository, release Release, os platform.OS, arch platform.Arch) (AssetMeta, error) {
	assets, err := a.asset.listFromBuiltIn(repo, release)
	if err != nil {
		return AssetMeta{}, err
	}
	return assets.find(os, arch)
}

// GetExecutableBinaryMeta return executable binary metadata from GitHub repository.
func (a *ApplicationService) GetExecutableBinaryMeta(repo Repository) ExecutableBinaryMeta {
	return newExecutableBinaryMetaFromRepository(repo, "")
}
