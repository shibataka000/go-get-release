package github

import (
	"context"
	"io"
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

// Find a GitHub release asset and an executable binary in it and returns them.
// repoFullName should be a repository full name. It should be OWNER/REPO format.
// tag should be a release tag.
func (a *ApplicationService) Find(ctx context.Context, repoFullName string, tag string, patterns map[string]string) (Asset, ExecBinary, error) {
	// New objects.
	repo, err := newRepositoryFromFullName(repoFullName)
	if err != nil {
		return Asset{}, ExecBinary{}, err
	}
	release := newRelease(tag)
	ps, err := newPatternArrayFromStringMap(patterns)
	if err != nil {
		return Asset{}, ExecBinary{}, err
	}

	// Find a GitHub release asset.
	assets, err := a.asset.list(ctx, repo, release)
	if err != nil {
		return Asset{}, ExecBinary{}, err
	}

	asset, pattern, err := find(assets, ps)
	if err != nil {
		return Asset{}, ExecBinary{}, err
	}

	// Find a executable binary.
	execBinary, err := pattern.execute(asset)
	if err != nil {
		return Asset{}, ExecBinary{}, err
	}

	return asset, execBinary, nil
}

func (a *ApplicationService) Install(ctx context.Context, repoFullName string, asset Asset, execBinary ExecBinary, dir string, w io.Writer) error {
	repo, err := newRepositoryFromFullName(repoFullName)
	if err != nil {
		return err
	}

	assetContent, err := a.asset.download(ctx, repo, asset, w)
	if err != nil {
		return err
	}

	execBinaryContent, err := assetContent.extract()
	if err != nil {
		return err
	}

	return a.execBinary.write(execBinary, execBinaryContent, dir)
}
