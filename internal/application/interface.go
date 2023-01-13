package application

import (
	"context"

	"github.com/shibataka000/go-get-release/internal/domain/model/asset"
	"github.com/shibataka000/go-get-release/internal/domain/model/execbinary"
	"github.com/shibataka000/go-get-release/internal/domain/model/platform"
	"github.com/shibataka000/go-get-release/internal/domain/model/release"
	"github.com/shibataka000/go-get-release/internal/domain/model/repository"
)

// GitHubRepository is repository for GitHub.
type GitHubRepository interface {
	FindRepository(ctx context.Context, owner string, name string) (repository.Repository, error)
	SearchRepository(ctx context.Context, query string) (repository.Repository, error)
	FindReleaseByTag(ctx context.Context, repo repository.Repository, tag string) (release.Release, error)
	LatestRelease(ctx context.Context, repo repository.Repository) (release.Release, error)
	ListAssetMetadata(ctx context.Context, repo repository.Repository, release release.Release) ([]asset.Metadata, error)
}

// HTTPRepository is repository for http.
type HTTPRepository interface {
	DownloadAssetContent(meta asset.Metadata, showProgressBar bool) (asset.Content, error)
}

// FileSystemRepository is repository for filesystem.
type FileSystemRepository interface {
	WriteExecBinaryContent(meta execbinary.Metadata, content execbinary.Content, dir string) error
}

// AssetFactory is factory to make GitHub asset metadata.
type AssetFactory interface {
	FindMetadata(ctx context.Context, repo repository.Repository, release release.Release, platform platform.Platform) (asset.Metadata, error)
}

// ExecBinaryFactory is factory to make executable binary metadata.
type ExecBinaryFactory interface {
	FindMetadata(repo repository.Repository, platform platform.Platform) (execbinary.Metadata, error)
}
