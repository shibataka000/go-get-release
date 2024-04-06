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
func NewSearchResult(repository Repository, release Release, assetMeta AssetMeta, execBinaryMeta ExecBinaryMeta) SearchResult {
	return SearchResult{
		Repository: repository,
		Release:    release,
		Asset:      assetMeta,
		ExecBinary: execBinaryMeta,
	}
}

// Search repository, release, asset, exec binary in GitHub.
func (a *ApplicationService) Search(ctx context.Context, queryStr string, os string, arch string) (SearchResult, error) {
	platform := NewPlatform(OS(os), Arch(arch))

	query, err := ParseQuery(queryStr)
	if err != nil {
		return SearchResult{}, err
	}

	repository, err := a.searchRepository(ctx, query)
	if err != nil {
		return SearchResult{}, err
	}

	release, err := a.findRelease(ctx, query, repository)
	if err != nil {
		return SearchResult{}, err
	}

	assetMeta, err := a.findAssetMeta(ctx, repository, release, platform)
	if err != nil {
		return SearchResult{}, err
	}

	execBinaryMeta, err := a.findExecBinaryMeta(repository, platform)
	if err != nil {
		return SearchResult{}, err
	}

	return NewSearchResult(repository, release, assetMeta, execBinaryMeta), nil
}

// searchRepository search repository in GitHub.
// If query speficy repository's owner and name, this return it.
// Otherwise, this search repository in GitHub and return it.
func (a *ApplicationService) searchRepository(ctx context.Context, query SearchQuery) (Repository, error) {
	if query.HasOwner() {
		return query.Repository, nil
	}
	return a.repository.SearchRepository(ctx, query.Repository.Name)
}

// findRelease return release in GitHub.
// If query speficy tag, this return it's release.
// Otherwise, this return latest release.
func (a *ApplicationService) findRelease(ctx context.Context, query SearchQuery, repository Repository) (Release, error) {
	if query.HasTag() {
		return query.Release, nil
	}
	return a.repository.LatestRelease(ctx, repository)
}

// findAssetMeta return asset meta in GitHub which suits specific platform.
// This try to find asset meta from known.yaml first.
// If not found, this try to find asset meta from GitHub release next.
func (a *ApplicationService) findAssetMeta(ctx context.Context, repository Repository, release Release, platform Platform) (AssetMeta, error) {
	return AssetMeta{}, nil
}

// findExecBinaryMeta return meta of exec binary in asset.
// If exec binary meta is defined in known.yaml, this return it.
// Otherwise, this treat repository name as exec binary name and return it.
func (a *ApplicationService) findExecBinaryMeta(repository Repository, platform Platform) (ExecBinaryMeta, error) {
	return ExecBinaryMeta{}, nil
}
