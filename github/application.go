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

// Find returns a GitHub release asset which matches given pattern.
// assetPatterns  must be a regular expression of GitHub release asset name and compilable by [regexp.Compile].
// execBinaryPatterns must be a template of executable binary name and parsable by [text/template.Template.Parse].
func (a *ApplicationService) Find(ctx context.Context, repoFullName string, tag string, assetPatterns []string, execBinaryPatterns []string) (Asset, ExecBinary, error) {
	// New objects.
	repo, err := newRepositoryFromFullName(repoFullName)
	if err != nil {
		return Asset{}, ExecBinary{}, err
	}
	release := newRelease(tag)
	patterns, err := newPatternListFromStringArray(assetPatterns, execBinaryPatterns)
	if err != nil {
		return Asset{}, ExecBinary{}, err
	}

	// Find a GitHub release asset.
	assets, err := a.asset.list(ctx, repo, release)
	if err != nil {
		return Asset{}, ExecBinary{}, err
	}

	asset, pattern, err := find(assets, patterns)
	if err != nil {
		return Asset{}, ExecBinary{}, err
	}

	// Find a executable binary.
	execBinary, err := pattern.apply(asset)
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

	execBinaryContent, err := assetContent.execBinaryContent()
	if err != nil {
		return err
	}

	return a.execBinary.write(execBinary, execBinaryContent, dir)
}

var (
	DefaultAssetPatterns = []string{
		"trivy_0.53.0_Linux-64bit.tar.gz",
	}
	DefaultExecBinaryPatterns = []string{
		"trivy",
	}
)
