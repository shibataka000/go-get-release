package github

import (
	"context"
	"regexp"
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
func (a *ApplicationService) FindAsset(ctx context.Context, repoFullName string, tag string, patterns []*regexp.Regexp) (Asset, error) {
	repo, err := newRepositoryFromFullName(repoFullName)
	if err != nil {
		return Asset{}, err
	}

	release := newRelease(tag)

	assets, err := a.asset.list(ctx, repo, release)
	if err != nil {
		return Asset{}, err
	}

	return assets.find(patterns)
}
