package pkg

import (
	"context"
	"fmt"
	"io"
	"regexp"
)

// ApplicationService.
type ApplicationService struct {
	repository *InfrastructureRepository
	factory    *Factory
}

// Query to search package.
type Query struct {
	Repository Repository
	Tag        string
}

// NewApplicationService return new application service instance.
func NewApplicationService(repository *InfrastructureRepository, factory *Factory) *ApplicationService {
	return &ApplicationService{
		repository: repository,
		factory:    factory,
	}
}

// NewQuery return new query instance.
func NewQuery(repo Repository, tag string) Query {
	return Query{
		Repository: repo,
		Tag:        tag,
	}
}

// Search package.
func (a *ApplicationService) Search(ctx context.Context, query Query, platform Platform) (Package, error) {
	var err error

	var ghRepo GitHubRepository
	if query.HasOwner() {
		ghRepo, err = a.repository.FindGitHubRepository(ctx, query.Repository.Owner, query.Repository.Name)
	} else {
		ghRepo, err = a.repository.SearchGitHubRepository(ctx, query.Repository.Name)
	}
	if err != nil {
		return Package{}, err
	}
	repo := a.factory.NewRepository(ghRepo)

	var ghRelease GitHubRelease
	if query.HasTag() {
		ghRelease, err = a.repository.FindGitHubReleaseByTag(ctx, ghRepo, query.Tag)
	} else {
		ghRelease, err = a.repository.LatestGitHubRelease(ctx, ghRepo)
	}
	if err != nil {
		return Package{}, err
	}
	release := a.factory.NewRelease(ghRelease)

	index, err := a.repository.LoadBuiltInIndex()
	if err != nil {
		return Package{}, err
	}

	var asset Asset
	if index.HasAsset(repo, platform) {
		assetInIndex, err := index.FindAsset(repo, platform)
		if err != nil {
			return Package{}, err
		}
		asset, err = a.factory.NewAssetFromIndex(assetInIndex, release)
		if err != nil {
			return Package{}, err
		}
	} else {
		ghAssets, err := a.repository.ListGitHubAssets(ctx, ghRepo, ghRelease)
		if err != nil {
			return Package{}, err
		}
		asset, err = a.factory.NewAssetFromGitHub(ghAssets, platform)
		if err != nil {
			return Package{}, err
		}
	}

	var execBinary ExecBinary
	if index.HasExecBinary(repo) {
		execBinaryInIndex, err := index.FindExecBinary(repo)
		if err != nil {
			return Package{}, err
		}
		execBinary = a.factory.NewExecBinaryFromIndex(execBinaryInIndex, platform)
	} else {
		execBinary = a.factory.NewExecBinaryFromGitHub(ghRepo, platform)
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

// ParseQuery parse query string and return query instance.
func ParseQuery(query string) (Query, error) {
	re := regexp.MustCompile(`(([^/=]+)/)?([^/=]+)(=([^/=]+))?`)
	submatch := re.FindStringSubmatch(query)
	if submatch == nil || len(submatch) != 6 {
		return Query{}, fmt.Errorf("%s is invalid query", query)
	}
	return NewQuery(NewRepository(submatch[2], submatch[3]), submatch[5]), nil
}

// HasOwner return true if query has repository owner.
func (q Query) HasOwner() bool {
	return q.Repository.Owner != ""
}

// HasTag return true if query has tag.
func (q Query) HasTag() bool {
	return q.Tag != ""
}
