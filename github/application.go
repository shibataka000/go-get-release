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
func (a *ApplicationService) Find(ctx context.Context, repoFullName string, tag string, assetPatterns []string, execBinaryTemplates []string) (Asset, ExecBinary, error) {
	repo, err := newRepositoryFromFullName(repoFullName)
	if err != nil {
		return Asset{}, ExecBinary{}, err
	}

	release := newRelease(tag)

	patterns, err := newPatternList(assetPatterns, execBinaryTemplates)
	if err != nil {
		return Asset{}, ExecBinary{}, err
	}

	assets, err := a.asset.list(ctx, repo, release)
	if err != nil {
		return Asset{}, ExecBinary{}, err
	}

	asset, err := assets.find(patterns)
	if err != nil {
		return Asset{}, ExecBinary{}, err
	}

	pattern, err := patterns.find(asset)
	if err != nil {
		return Asset{}, ExecBinary{}, err
	}

	execBinary, err := pattern.renderExecBinary(asset)

	return asset, execBinary, nil
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
