package github

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func NewInfrastructureRepositoryForTest(ctx context.Context, t *testing.T) *InfrastructureRepository {
	t.Helper()
	return NewInfrastructureRepository(ctx, os.Getenv("GITHUB_TOKEN"))
}

func TestInfrastructureRepositorySearchRepository(t *testing.T) {
	tests := []struct {
		name       string
		query      string
		repository Repository
	}{
		{
			name:       "terraform",
			query:      "terraform",
			repository: NewRepository("hashicorp", "terraform"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require := require.New(t)
			ctx := context.Background()
			repository := NewInfrastructureRepositoryForTest(ctx, t)
			repo, err := repository.SearchRepository(ctx, tt.query)
			require.NoError(err)
			require.Equal(tt.repository, repo)
		})
	}
}

func TestInfrastructureRepositoryLatestRelease(t *testing.T) {
	tests := []struct {
		name       string
		repository Repository
		release    Release
	}{
		{
			name:       "shibataka000/go-get-release-test",
			repository: NewRepository("shibataka000", "go-get-release-test"),
			release:    NewRelease("v0.0.2"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require := require.New(t)
			ctx := context.Background()
			repository := NewInfrastructureRepositoryForTest(ctx, t)
			release, err := repository.LatestRelease(ctx, tt.repository)
			require.NoError(err)
			require.Equal(tt.release, release)
		})
	}
}

func TestInfrastructureRepositoryListGitHubAssets(t *testing.T) {
	tests := []struct {
		name       string
		repository Repository
		release    Release
		assets     []AssetMeta
	}{
		{
			name:       "cli/cli",
			repository: NewRepository("cli", "cli"),
			release:    NewRelease("v2.21.1"),
			assets: []AssetMeta{
				NewAssetMeta("https://github.com/cli/cli/releases/download/v2.21.1/gh_2.21.1_checksums.txt"),
				NewAssetMeta("https://github.com/cli/cli/releases/download/v2.21.1/gh_2.21.1_linux_386.deb"),
				NewAssetMeta("https://github.com/cli/cli/releases/download/v2.21.1/gh_2.21.1_linux_386.rpm"),
				NewAssetMeta("https://github.com/cli/cli/releases/download/v2.21.1/gh_2.21.1_linux_386.tar.gz"),
				NewAssetMeta("https://github.com/cli/cli/releases/download/v2.21.1/gh_2.21.1_linux_amd64.deb"),
				NewAssetMeta("https://github.com/cli/cli/releases/download/v2.21.1/gh_2.21.1_linux_amd64.rpm"),
				NewAssetMeta("https://github.com/cli/cli/releases/download/v2.21.1/gh_2.21.1_linux_amd64.tar.gz"),
				NewAssetMeta("https://github.com/cli/cli/releases/download/v2.21.1/gh_2.21.1_linux_arm64.deb"),
				NewAssetMeta("https://github.com/cli/cli/releases/download/v2.21.1/gh_2.21.1_linux_arm64.rpm"),
				NewAssetMeta("https://github.com/cli/cli/releases/download/v2.21.1/gh_2.21.1_linux_arm64.tar.gz"),
				NewAssetMeta("https://github.com/cli/cli/releases/download/v2.21.1/gh_2.21.1_linux_armv6.deb"),
				NewAssetMeta("https://github.com/cli/cli/releases/download/v2.21.1/gh_2.21.1_linux_armv6.rpm"),
				NewAssetMeta("https://github.com/cli/cli/releases/download/v2.21.1/gh_2.21.1_linux_armv6.tar.gz"),
				NewAssetMeta("https://github.com/cli/cli/releases/download/v2.21.1/gh_2.21.1_macOS_amd64.tar.gz"),
				NewAssetMeta("https://github.com/cli/cli/releases/download/v2.21.1/gh_2.21.1_windows_386.zip"),
				NewAssetMeta("https://github.com/cli/cli/releases/download/v2.21.1/gh_2.21.1_windows_amd64.msi"),
				NewAssetMeta("https://github.com/cli/cli/releases/download/v2.21.1/gh_2.21.1_windows_amd64.zip"),
				NewAssetMeta("https://github.com/cli/cli/releases/download/v2.21.1/gh_2.21.1_windows_arm64.zip"),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require := require.New(t)
			ctx := context.Background()
			repository := NewInfrastructureRepositoryForTest(ctx, t)
			assets, err := repository.ListAssetMeta(ctx, tt.repository, tt.release)
			require.NoError(err)
			require.Equal(tt.assets, assets)
		})
	}
}
