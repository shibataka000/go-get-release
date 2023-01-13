package asset

import (
	"context"

	"github.com/shibataka000/go-get-release/internal/domain/model/asset"
	"github.com/shibataka000/go-get-release/internal/domain/model/platform"
	"github.com/shibataka000/go-get-release/internal/domain/model/release"
	"github.com/shibataka000/go-get-release/internal/domain/model/repository"
	specification "github.com/shibataka000/go-get-release/internal/domain/specification/asset"
)

// Factory to make GitHub asset metadata instance.
type Factory struct {
	github  GitHubRepository
	embed   EmbedRepository
	service Service
}

// New factory instance.
func New(github GitHubRepository, embed EmbedRepository, service Service) Factory {
	return Factory{
		github:  github,
		embed:   embed,
		service: service,
	}
}

// FindMetadata return GitHub asset metadata.
// If metadata is defined in embed data, return it.
// Otherwise, find metadata from GitHub release and return it.
func (f Factory) FindMetadata(ctx context.Context, repo repository.Repository, release release.Release, platform platform.Platform) (asset.Metadata, error) {
	if f.embed.HasAssetMetadata(repo, release, platform) {
		return f.embed.FindAssetMetadata(repo, release, platform)
	}
	meta, err := f.github.ListAssetMetadata(ctx, repo, release)
	if err != nil {
		return asset.Metadata{}, err
	}
	spec := specification.NewFindMetadataByPlatform(platform)
	return f.service.FindMetadata(meta, spec)
}
