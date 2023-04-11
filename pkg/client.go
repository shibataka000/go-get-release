package pkg

import (
	"context"
	"io"
)

// Client for package.
type Client struct {
	repository *Repository
	factory    *Factory
}

// NewClient return new client instance.
func NewClient(ctx context.Context, token string) (*Client, error) {
	repository := NewRepository(ctx, token)
	index, err := repository.LoadBuiltInIndex()
	if err != nil {
		return &Client{}, err
	}
	factory := NewFactory(index)
	return &Client{
		repository: repository,
		factory:    factory,
	}, nil
}

// Search package.
func (c *Client) Search(ctx context.Context, query Query, platform Platform) (Package, error) {
	var err error

	var repo GitHubRepository
	if query.HasFullRepositoryName() {
		repo, err = c.repository.FindGitHubRepository(ctx, query.Repository.Owner, query.Repository.Name)
	} else {
		repo, err = c.repository.SearchGitHubRepository(ctx, query.SearchRepositoryQuery())
	}
	if err != nil {
		return Package{}, err
	}

	var release GitHubRelease
	if query.HasTag() {
		release, err = c.repository.FindGitHubReleaseByTag(ctx, repo, query.Tag)
	} else {
		release, err = c.repository.LatestGitHubRelease(ctx, repo)
	}
	if err != nil {
		return Package{}, err
	}

	assets, err := c.repository.ListGitHubAssets(ctx, repo, release)
	if err != nil {
		return Package{}, err
	}
	asset, err := c.factory.NewAssetMeta(repo, release, assets, platform)
	if err != nil {
		return Package{}, err
	}

	execBinary, err := c.factory.NewExecBinaryMeta(repo, platform)
	if err != nil {
		return Package{}, err
	}

	return New(repo, release, asset, execBinary), nil
}

// Install package.
func (c *Client) Install(pkg Package, dir string, progressBar io.Writer) error {
	asset, err := c.repository.Download(pkg.Asset.DownloadURL, progressBar)
	if err != nil {
		return err
	}
	execBinary, err := AssetFile(asset).ExecBinary(pkg.ExecBinary.Name)
	if err != nil {
		return err
	}
	return c.repository.WriteFile(File(execBinary), dir, 0755)
}
