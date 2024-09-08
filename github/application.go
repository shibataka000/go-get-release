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

// NewApplicationService returns a new [ApplicationService] object.
func NewApplicationService(asset *AssetRepository, execBinary *ExecBinaryRepository) *ApplicationService {
	return &ApplicationService{
		asset:      asset,
		execBinary: execBinary,
	}
}

// FindAsset returns a GitHub release asset which matches given pattern.
func (a *ApplicationService) FindAsset(ctx context.Context, repoFullName string, tag string, assetPatterns []string) (Asset, error) {
	repo, err := newRepositoryFromFullName(repoFullName)
	if err != nil {
		return Asset{}, err
	}

	release := newRelease(tag)

	patterns, err := compileAssetPatternList(assetPatterns)
	if err != nil {
		return Asset{}, err
	}

	assets, err := a.asset.list(ctx, repo, release)
	if err != nil {
		return Asset{}, err
	}

	return assets.find(patterns)
}

func (a *ApplicationService) FindExecBinary(asset Asset, assetPatterns []string, execBinaryPatterns []string) (ExecBinary, error) {
	apl, err := compileAssetPatternList(assetPatterns)
	if err != nil {
		return ExecBinary{}, err
	}

	ebpl := newExecBinaryTemplateList(execBinaryPatterns)

	return asset.execBinary(apl, ebpl)
}

func (a *ApplicationService) Install(ctx context.Context, repoFullName string, asset Asset, execBinary ExecBinary, dir string) error {
	repo, err := newRepositoryFromFullName(repoFullName)
	if err != nil {
		return err
	}

	assetContent, err := a.asset.download(ctx, repo, asset, os.Stdout)
	if err != nil {
		return err
	}

	execBinaryContent, err := assetContent.execBinaryContent()
	if err != nil {
		return err
	}

	return a.execBinary.write(execBinary, execBinaryContent, dir)
}
