package asset

import (
	"context"

	"github.com/shibataka000/go-get-release/internal/domain/model/asset"
	"github.com/shibataka000/go-get-release/internal/domain/model/platform"
	"github.com/shibataka000/go-get-release/internal/domain/model/release"
	"github.com/shibataka000/go-get-release/internal/domain/model/repository"
	specification "github.com/shibataka000/go-get-release/internal/domain/specification/asset"
)

// GitHubRepository is interface for GitHub.
type GitHubRepository interface {
	ListAssetMetadata(ctx context.Context, repo repository.Repository, release release.Release) ([]asset.Metadata, error)
}

// EmbedRepository is interface for embed data.
type EmbedRepository interface {
	FindAssetMetadata(repo repository.Repository, release release.Release, platform platform.Platform) (asset.Metadata, error)
	HasAssetMetadata(repo repository.Repository, release release.Release, platform platform.Platform) bool
}

// Service is domain service for asset.
type Service interface {
	FindMetadata(meta []asset.Metadata, spec specification.FindMetadata) (asset.Metadata, error)
}
