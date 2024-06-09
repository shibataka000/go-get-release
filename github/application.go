package github

import (
	"context"

	"github.com/shibataka000/go-get-release/platform"
)

// ApplicationService.
type ApplicationService struct {
	asset *AssetRepository
	bin   *ExecutableBinaryRepository
}

// NewApplicationService return new ApplicationService object.
func NewApplicationService(asset *AssetRepository, bin *ExecutableBinaryRepository) *ApplicationService {
	return &ApplicationService{
		asset: asset,
		bin:   bin,
	}
}

func (a *ApplicationService) FindAssetMeta(ctx context.Context, repoFullName string, tag string, os platform.OS, arch platform.Arch) (Asset, error) {
	repo, err := newRepositoryFromFullName(repoFullName)
	if err != nil {
		return Asset{}, err
	}
	release := newRelease(tag)
	if asset, err := a.findAssetMetaFromBuiltIn(repo, release, os, arch); err == nil {
		return asset, nil
	}
	return a.findAssetMetaFromAPI(ctx, repo, release, os, arch)
}

func (a *ApplicationService) findAssetMetaFromAPI(ctx context.Context, repo Repository, release Release, os platform.OS, arch platform.Arch) (Asset, error) {
	assets, err := a.asset.listFromAPI(ctx, repo, release)
	if err != nil {
		return Asset{}, err
	}
	return assets.find(os, arch)
}

func (a *ApplicationService) findAssetMetaFromBuiltIn(repo Repository, release Release, os platform.OS, arch platform.Arch) (Asset, error) {
	assets, err := a.asset.listFromBuiltIn(repo, release)
	if err != nil {
		return Asset{}, err
	}
	return assets.find(os, arch)
}

func (a *ApplicationService) FindExecutableBinaryMeta(repo Repository, os platform.OS) ExecutableBinary {
	if bin, err := a.bin.find(repo, os); err == nil {
		return bin
	}
	return newExecutableBinaryMetaFromRepository(repo, os)
}
