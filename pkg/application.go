package pkg

import (
	"context"
	"io"
)

// ApplicationService for package.
type ApplicationService struct {
	repository *Repository
	factory    *Factory
}

// NewApplicationService return new application service instance.
func NewApplicationService(ctx context.Context, token string) (*ApplicationService, error) {
	repository := NewRepository(ctx, token)
	index, err := repository.LoadBuiltInIndex()
	if err != nil {
		return &ApplicationService{}, err
	}
	factory := NewFactory(index)
	return &ApplicationService{
		repository: repository,
		factory:    factory,
	}, nil
}

// Search package.
func (a *ApplicationService) Search(ctx context.Context, query Query, platform Platform) (Package, error) {
	var err error

	var repo GitHubRepository
	if query.HasFullRepositoryName() {
		repo, err = a.repository.FindGitHubRepository(ctx, query.Repository.Owner, query.Repository.Name)
	} else {
		repo, err = a.repository.SearchGitHubRepository(ctx, query.SearchRepositoryQuery())
	}
	if err != nil {
		return Package{}, err
	}

	var release GitHubRelease
	if query.HasTag() {
		release, err = a.repository.FindGitHubReleaseByTag(ctx, repo, query.Tag)
	} else {
		release, err = a.repository.LatestGitHubRelease(ctx, repo)
	}
	if err != nil {
		return Package{}, err
	}

	assets, err := a.repository.ListGitHubAssets(ctx, repo, release)
	if err != nil {
		return Package{}, err
	}
	asset, err := a.factory.NewAssetMeta(repo, release, assets, platform)
	if err != nil {
		return Package{}, err
	}

	execBinary, err := a.factory.NewExecBinaryMeta(repo, platform)
	if err != nil {
		return Package{}, err
	}

	return New(repo, release, asset, execBinary), nil
}

// Install package.
func (a *ApplicationService) Install(pkg Package, dir string, progressBar io.Writer) error {
	asset, err := a.repository.Download(pkg.Asset.DownloadURL, progressBar)
	if err != nil {
		return err
	}
	execBinary, err := AssetFile(asset).ExecBinary(pkg.ExecBinary.Name)
	if err != nil {
		return err
	}
	return a.repository.WriteFile(File(execBinary), dir, 0755)
}
