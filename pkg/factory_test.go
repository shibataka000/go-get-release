package pkg

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFactoryNewAssetMeta(t *testing.T) {
	tests := []struct {
		name       string
		repository GitHubRepository
		release    GitHubRelease
		assets     []GitHubAsset
		platform   Platform
		asset      AssetMeta
	}{
		{
			name:       "cli/cli",
			repository: NewGitHubRepository("cli", "cli"),
			release:    NewGitHubRelease(0, "v2.21.1"),
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
			platform: NewPlatform("linux", "amd64"),
			asset:    NewAssetMeta("https://github.com/cli/cli/releases/download/v2.21.1/gh_2.21.1_linux_amd64.tar.gz"),
		},
		{
			name:       "hashicorp/terraform",
			repository: NewGitHubRepository("hashicorp", "terraform"),
			release:    NewGitHubRelease(0, "v0.12.20"),
			assets:     []GitHubAsset{},
			platform:   NewPlatform("linux", "amd64"),
			asset:      NewAssetMeta("https://releases.hashicorp.com/terraform/0.12.20/terraform_0.12.20_linux_amd64.zip"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := require.New(t)
			index, err := LoadIndexForTest(t)
			assert.NoError(err)
			factory := NewFactory(index)
			asset, err := factory.NewAssetMeta(tt.repository, tt.release, tt.assets, tt.platform)
			assert.NoError(err)
			assert.Equal(tt.asset, asset)
		})
	}
}

func TestFactoryNewExecBinaryMeta(t *testing.T) {
	tests := []struct {
		name       string
		index      Index
		repository GitHubRepository
		platform   Platform
		execBinary ExecBinaryMeta
	}{
		{
			name: "cli/cli",
			index: NewIndex([]RepositoryInIndex{
				NewRepositoryInIndex("cli", "cli", []AssetInIndex{}, NewExecBinaryInIndex("gh")),
			}),
			repository: NewGitHubRepository("cli", "cli"),
			platform:   NewPlatform("linux", "amd64"),
			execBinary: NewExecBinaryMeta("gh"),
		},
		{
			name:       "cli/cli",
			index:      NewIndex([]RepositoryInIndex{}),
			repository: NewGitHubRepository("cli", "cli"),
			platform:   NewPlatform("linux", "amd64"),
			execBinary: NewExecBinaryMeta("cli"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := require.New(t)
			factory := NewFactory(tt.index)
			execBinary, err := factory.NewExecBinaryMeta(tt.repository, tt.platform)
			assert.NoError(err)
			assert.Equal(tt.execBinary, execBinary)
		})
	}
}

func TestFindAssetByPlatform(t *testing.T) {
	tests := []struct {
		name     string
		assets   []GitHubAsset
		platform Platform
		asset    GitHubAsset
	}{
		{
			name: "cli/cli",
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
			platform: NewPlatform("linux", "amd64"),
			asset:    NewGitHubAsset("https://github.com/cli/cli/releases/download/v2.21.1/gh_2.21.1_linux_amd64.tar.gz"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := require.New(t)
			asset, err := findAssetByPlatform(tt.assets, tt.platform)
			assert.NoError(err)
			assert.Equal(tt.asset, asset)
		})
	}
}
