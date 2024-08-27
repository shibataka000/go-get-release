package github

import (
	"context"
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

// FindAsset returns a GitHub release asset which matches given pattern.
func (a *ApplicationService) FindAsset(ctx context.Context, repoFullName string, tag string, assetPatterns []string, execBinaryPatterns []string) (Asset, error) {
	repo, err := newRepositoryFromFullName(repoFullName)
	if err != nil {
		return Asset{}, err
	}

	release := newRelease(tag)

	patterns, err := newPatternListFromStringSlice(assetPatterns, execBinaryPatterns)
	if err != nil {
		return Asset{}, err
	}

	assets, err := a.asset.list(ctx, repo, release)
	if err != nil {
		return Asset{}, err
	}

	return assets.find(patterns)
}
