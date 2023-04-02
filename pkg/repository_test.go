package pkg

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func NewRepositoryForTest(ctx context.Context, t *testing.T) *Repository {
	t.Helper()
	return NewRepository(ctx, os.Getenv("GITHUB_TOKEN"))
}

func TestRepositorySearchGitHubRepository(t *testing.T) {
	tests := []struct {
		name       string
		query      string
		repository GitHubRepository
	}{
		{
			name:       "terraform",
			query:      "terraform",
			repository: NewGitHubRepository("hashicorp", "terraform"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := require.New(t)
			ctx := context.Background()
			repository := NewRepositoryForTest(ctx, t)
			repo, err := repository.SearchGitHubRepository(ctx, tt.query)
			assert.NoError(err)
			assert.Equal(tt.repository, repo)
		})
	}
}

func TestRepositoryFindGitHubRepository(t *testing.T) {
	tests := []struct {
		name       string
		owner      string
		repoName   string
		repository GitHubRepository
	}{
		{
			name:       "hashicorp/terraform",
			owner:      "hashicorp",
			repoName:   "terraform",
			repository: NewGitHubRepository("hashicorp", "terraform"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := require.New(t)
			ctx := context.Background()
			repository := NewRepositoryForTest(ctx, t)
			repo, err := repository.FindGitHubRepository(ctx, tt.owner, tt.repoName)
			assert.NoError(err)
			assert.Equal(tt.repository, repo)
		})
	}
}

func TestRepositoryLatestGitHubRelease(t *testing.T) {
	tests := []struct {
		name       string
		repository GitHubRepository
		release    GitHubRelease
	}{
		{
			name:       "shibataka000/go-get-release-test",
			repository: NewGitHubRepository("shibataka000", "go-get-release-test"),
			release:    NewGitHubRelease(23712476, "v0.0.2"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := require.New(t)
			ctx := context.Background()
			repository := NewRepositoryForTest(ctx, t)
			release, err := repository.LatestGitHubRelease(ctx, tt.repository)
			assert.NoError(err)
			assert.Equal(tt.release, release)
		})
	}
}

func TestRepositoryFindGitHubReleaseByTag(t *testing.T) {
	tests := []struct {
		name       string
		repository GitHubRepository
		tag        string
		release    GitHubRelease
	}{
		{
			name:       "shibataka000/go-get-release-test",
			repository: NewGitHubRepository("shibataka000", "go-get-release-test"),
			tag:        "v0.0.1",
			release:    NewGitHubRelease(23712441, "v0.0.1"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := require.New(t)
			ctx := context.Background()
			repository := NewRepositoryForTest(ctx, t)
			release, err := repository.FindGitHubReleaseByTag(ctx, tt.repository, tt.tag)
			assert.NoError(err)
			assert.Equal(tt.release, release)
		})
	}
}

func TestRepositoryListGitHubAssets(t *testing.T) {
	tests := []struct {
		name       string
		repository GitHubRepository
		release    GitHubRelease
		assets     []GitHubAsset
	}{
		{
			name:       "cli/cli",
			repository: NewGitHubRepository("cli", "cli"),
			release:    NewGitHubRelease(87098162, "v2.21.1"),
			assets: []GitHubAsset{
				NewGitHubAsset("https://github.com/cli/cli/releases/download/v2.21.1/gh_2.21.1_checksums.txt"),
				NewGitHubAsset("https://github.com/cli/cli/releases/download/v2.21.1/gh_2.21.1_linux_386.deb"),
				NewGitHubAsset("https://github.com/cli/cli/releases/download/v2.21.1/gh_2.21.1_linux_386.rpm"),
				NewGitHubAsset("https://github.com/cli/cli/releases/download/v2.21.1/gh_2.21.1_linux_386.tar.gz"),
				NewGitHubAsset("https://github.com/cli/cli/releases/download/v2.21.1/gh_2.21.1_linux_amd64.deb"),
				NewGitHubAsset("https://github.com/cli/cli/releases/download/v2.21.1/gh_2.21.1_linux_amd64.rpm"),
				NewGitHubAsset("https://github.com/cli/cli/releases/download/v2.21.1/gh_2.21.1_linux_amd64.tar.gz"),
				NewGitHubAsset("https://github.com/cli/cli/releases/download/v2.21.1/gh_2.21.1_linux_arm64.deb"),
				NewGitHubAsset("https://github.com/cli/cli/releases/download/v2.21.1/gh_2.21.1_linux_arm64.rpm"),
				NewGitHubAsset("https://github.com/cli/cli/releases/download/v2.21.1/gh_2.21.1_linux_arm64.tar.gz"),
				NewGitHubAsset("https://github.com/cli/cli/releases/download/v2.21.1/gh_2.21.1_linux_armv6.deb"),
				NewGitHubAsset("https://github.com/cli/cli/releases/download/v2.21.1/gh_2.21.1_linux_armv6.rpm"),
				NewGitHubAsset("https://github.com/cli/cli/releases/download/v2.21.1/gh_2.21.1_linux_armv6.tar.gz"),
				NewGitHubAsset("https://github.com/cli/cli/releases/download/v2.21.1/gh_2.21.1_macOS_amd64.tar.gz"),
				NewGitHubAsset("https://github.com/cli/cli/releases/download/v2.21.1/gh_2.21.1_windows_386.zip"),
				NewGitHubAsset("https://github.com/cli/cli/releases/download/v2.21.1/gh_2.21.1_windows_amd64.msi"),
				NewGitHubAsset("https://github.com/cli/cli/releases/download/v2.21.1/gh_2.21.1_windows_amd64.zip"),
				NewGitHubAsset("https://github.com/cli/cli/releases/download/v2.21.1/gh_2.21.1_windows_arm64.zip"),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := require.New(t)
			ctx := context.Background()
			repository := NewRepositoryForTest(ctx, t)
			assets, err := repository.ListGitHubAssets(ctx, tt.repository, tt.release)
			assert.NoError(err)
			assert.Equal(tt.assets, assets)
		})
	}
}

func TestDownload(t *testing.T) {
	tests := []struct {
		name string
		url  URL
		file File
	}{
		{
			name: "https://raw.githubusercontent.com/shibataka000/go-get-release-test/master/README.md",
			url:  "https://raw.githubusercontent.com/shibataka000/go-get-release-test/master/README.md",
			file: NewFile("README.md", []byte("# go-get-release-test")),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := require.New(t)
			ctx := context.Background()
			repository := NewRepositoryForTest(ctx, t)
			file, err := repository.Download(tt.url)
			assert.NoError(err)
			assert.Equal(tt.file, file)
		})
	}
}

func TestDownloadAsset(t *testing.T) {
	tests := []struct {
		name string
		url  URL
		file AssetFile
	}{
		{
			name: "https://github.com/shibataka000/go-get-release-test/releases/download/v0.0.2/testdata.txt",
			url:  "https://github.com/shibataka000/go-get-release-test/releases/download/v0.0.2/testdata.txt",
			file: NewAssetFile("testdata.txt", []byte("helloworld\n")),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := require.New(t)
			ctx := context.Background()
			repository := NewRepositoryForTest(ctx, t)
			file, err := repository.DownloadAsset(tt.url)
			assert.NoError(err)
			assert.Equal(tt.file, file)
		})
	}
}
