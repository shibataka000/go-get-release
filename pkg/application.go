package pkg

import (
	"context"
	"fmt"
	"io"
	"regexp"
)

// ApplicationService for package.
type ApplicationService struct {
	repository *InfrastructureRepository
	factory    *Factory
}

// Query to search package.
type Query struct {
	Repository GitHubRepository
	Tag        string
}

// NewApplicationService return new application service instance.
func NewApplicationService(ctx context.Context, token string) (*ApplicationService, error) {
	repository := NewInfrastructureRepository(ctx, token)
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

// NewQuery return new query instance to search package.
func NewQuery(repo GitHubRepository, tag string) Query {
	return Query{
		Repository: repo,
		Tag:        tag,
	}
}

// ParseQuery parse query string and return query instance.
func ParseQuery(query string) (Query, error) {
	re := regexp.MustCompile(`(([^/=]+)/)?([^/=]+)(=([^/=]+))?`)
	submatch := re.FindStringSubmatch(query)
	if submatch == nil || len(submatch) != 6 {
		return Query{}, fmt.Errorf("%s is invalid query", query)
	}
	return NewQuery(NewGitHubRepository(submatch[2], submatch[3]), submatch[5]), nil
}

// SearchRepositoryQuery return query string to search GitHub repository by SearchRepository function.
func (q Query) SearchRepositoryQuery() string {
	if q.Repository.Owner == "" {
		return q.Repository.Name
	}
	return q.Repository.FullName()
}

// HasFullRepositoryName return true if query string has both of owner and name of GitHub repository.
func (q Query) HasFullRepositoryName() bool {
	return q.Repository.Owner != "" && q.Repository.Name != ""
}

// HasTag return true if query string has GitHub release tag.
func (q Query) HasTag() bool {
	return q.Tag != ""
}
