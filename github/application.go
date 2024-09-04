package github

import (
	"context"
	"os"
)

// ApplicationService.
type ApplicationService struct {
	asset      *AssetRepository
	execBinary *ExecBinaryRepository
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

	patterns, err := newPatternListFromStringArray(assetPatterns, execBinaryPatterns)
	if err != nil {
		return Asset{}, err
	}

	assets, err := a.asset.list(ctx, repo, release)
	if err != nil {
		return Asset{}, err
	}

	return assets.find(patterns)
}

func (a *ApplicationService) Install(asset Asset, execBinary ExecBinary) error {
	assetContent, err := a.asset.download(asset, os.Stdout)
	if err != nil {
		return err
	}
	execBinaryContent, err := assetContent.execBinary()
	if err != nil {
		return err
	}
	return a.execBinary.write(execBinary, execBinaryContent)
}
