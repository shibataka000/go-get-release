package github

import (
	"context"

	"github.com/shibataka000/go-get-release/platform"
)

// ApplicationService.
type ApplicationService struct {
	asset *AssetRepository
}

// NewApplicationService returns a new ApplicationService object.
func NewApplicationService(asset *AssetRepository) *ApplicationService {
	return &ApplicationService{
		asset: asset,
	}
}

func (a *ApplicationService) FindAsset(ctx context.Context, repoFullName string, tag string, os platform.OS, arch platform.Arch) (Asset, error) {
	repo, err := newRepositoryFromFullName(repoFullName)
	if err != nil {
		return Asset{}, err
	}

	release := newRelease(tag)

	assets, err := a.asset.list(ctx, repo, release)
	if err != nil {
		return Asset{}, err
	}

	externalAssets, err := a.asset.listExternal(repo, release)
	if err != nil {
		return Asset{}, err
	}

	return append(assets, externalAssets...).find(os, arch)
}
