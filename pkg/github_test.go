package pkg

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGitHubAssetHasExecBinary(t *testing.T) {
	tests := []struct {
		name          string
		asset         GitHubAsset
		hasExecBinary bool
	}{
		{
			name:          "https://github.com/cli/cli/releases/download/v2.21.0/gh_2.21.0_linux_amd64.tar.gz",
			asset:         NewGitHubAsset("https://github.com/cli/cli/releases/download/v2.21.0/gh_2.21.0_linux_amd64.tar.gz"),
			hasExecBinary: true,
		},
		{
			name:          "https://github.com/cli/cli/releases/download/v2.21.0/gh_2.21.0_checksums.txt",
			asset:         NewGitHubAsset("https://github.com/cli/cli/releases/download/v2.21.0/gh_2.21.0_checksums.txt"),
			hasExecBinary: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := require.New(t)
			assert.Equal(tt.hasExecBinary, tt.asset.HasExecBinary())
		})
	}
}

func TestGitHubAssetPlatform(t *testing.T) {
	tests := []struct {
		name     string
		asset    GitHubAsset
		platform Platform
	}{
		{
			name:     "https://github.com/cli/cli/releases/download/v2.21.0/gh_2.21.0_linux_amd64.tar.gz",
			asset:    NewGitHubAsset("https://github.com/cli/cli/releases/download/v2.21.0/gh_2.21.0_linux_amd64.tar.gz"),
			platform: NewPlatform("linux", "amd64"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := require.New(t)
			platform, err := tt.asset.Platform()
			assert.NoError(err)
			assert.Equal(tt.platform, platform)
		})
	}
}

func TestFilterGitHubAssetByPlatform(t *testing.T) {
	tests := []struct {
		name     string
		assets   []GitHubAsset
		platform Platform
		filtered []GitHubAsset
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
			filtered: []GitHubAsset{
				NewGitHubAsset("https://github.com/cli/cli/releases/download/v2.21.1/gh_2.21.1_linux_amd64.tar.gz"),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := require.New(t)
			filtered := FilterGitHubAssetByPlatform(tt.assets, tt.platform)
			assert.Equal(tt.filtered, filtered)
		})
	}
}
