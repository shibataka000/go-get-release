package github

import (
	"context"
)

// ApplicationService.
type ApplicationService struct {
	repository *InfrastructureRepository
}

// SearchResult is result of ApplicationService.Search function.
type SearchResult struct {
	Repository Repository
	Release    Release
	Asset      AssetMeta
	ExecBinary ExecBinaryMeta
}

// NewApplicationService return new application service instance.
func NewApplicationService(repository *InfrastructureRepository) *ApplicationService {
	return &ApplicationService{
		repository: repository,
	}
}

// NewSearchResult return new SearchResult instance.
func NewSearchResult(repo Repository, release Release, asset AssetMeta, execBinary ExecBinaryMeta) SearchResult {
	return SearchResult{
		Repository: repo,
		Release:    release,
		Asset:      asset,
		ExecBinary: execBinary,
	}
}

// Search repository, release, asset, exec binary in GitHub.
func (a *ApplicationService) Search(ctx context.Context, query string, os string, arch string) (SearchResult, error) {
	platform := NewPlatform(OS(os), Arch(arch))

	q, err := ParseSearchQuery(query)
	if err != nil {
		return SearchResult{}, err
	}

	repo, err := a.searchRepository(ctx, q)
	if err != nil {
		return SearchResult{}, err
	}

	release, err := a.findRelease(ctx, q, repo)
	if err != nil {
		return SearchResult{}, err
	}

	assetMeta, err := a.findAssetMeta(ctx, repo, release, platform)
	if err != nil {
		return SearchResult{}, err
	}

	execBinaryMeta, err := a.findExecBinaryMeta(repo, platform)
	if err != nil {
		return SearchResult{}, err
	}

	return NewSearchResult(repo, release, assetMeta, execBinaryMeta), nil
}

// searchRepository search repository in GitHub.
// If query speficy repository's owner and name, this return it.
// Otherwise, this search repository in GitHub and return it.
func (a *ApplicationService) searchRepository(ctx context.Context, query SearchQuery) (Repository, error) {
	if query.HasRepository() {
		return query.Repository, nil
	}
	return a.repository.SearchRepository(ctx, query.SearchRepositoryQuery())
}

// findRelease return release in GitHub.
// If query speficy tag, this return it's release.
// Otherwise, this return latest release.
func (a *ApplicationService) findRelease(ctx context.Context, query SearchQuery, repo Repository) (Release, error) {
	if query.HasRelease() {
		return query.Release, nil
	}
	return a.repository.LatestRelease(ctx, repo)
}

// findAssetMeta return asset meta in GitHub which suits specific platform.
// This try to find asset meta from known.yaml first.
// If not found, this try to find asset meta from GitHub release next.
func (a *ApplicationService) findAssetMeta(ctx context.Context, repo Repository, release Release, platform Platform) (AssetMeta, error) {
	if asset, err := a.findAssetMetaFromKnownEntry(repo, release, platform); err == nil {
		return asset, err
	}
	return a.findAssetMetaFromGitHub(ctx, repo, release, platform)
}

func (a *ApplicationService) findAssetMetaFromGitHub(ctx context.Context, repo Repository, release Release, platform Platform) (AssetMeta, error) {
	assets, err := a.repository.ListAssetMeta(ctx, repo, release)
	if err != nil {
		return AssetMeta{}, err
	}
	return assets.FindByPlatform(platform)
}

func (a *ApplicationService) findAssetMetaFromKnownEntry(repo Repository, release Release, platform Platform) (AssetMeta, error) {
	// knownEntries, err := a.repository.ListKnownEntries()
	// if err != nil {
	// 	return AssetMeta{}, err
	// }
	// knownEntry, err := knownEntries.Find(repo)
	// if err != nil {
	// 	return AssetMeta{}, err
	// }
	// knownEntry.Assets
	return AssetMeta{}, nil
}

// findExecBinaryMeta return meta of exec binary in asset.
// If exec binary meta is defined in known.yaml, this return it.
// Otherwise, this treat repository name as exec binary name and return it.
func (a *ApplicationService) findExecBinaryMeta(repo Repository, platform Platform) (ExecBinaryMeta, error) {
	return ExecBinaryMeta{}, nil
}
