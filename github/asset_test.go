package github

import (
	"context"
	"net/url"
	"os"
	"regexp"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAssetName(t *testing.T) {
	tests := []struct {
		name  string
		asset Asset
	}{
		{
			name:  "gh_2.52.0_linux_amd64.tar.gz",
			asset: mustNewAsset("https://github.com/cli/cli/releases/download/v2.52.0/gh_2.52.0_linux_amd64.tar.gz"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require := require.New(t)
			require.Equal(tt.name, tt.asset.name())
		})
	}
}

func TestAssetListFind(t *testing.T) {
	tests := []struct {
		name     string
		assets   AssetList
		patterns []*regexp.Regexp
		asset    Asset
	}{
		{
			name: "gh_2.52.0_linux_amd64.tar.gz",
			assets: AssetList{
				mustNewAsset("https://github.com/cli/cli/releases/download/v2.52.0/gh_2.52.0_checksums.txt"),
				mustNewAsset("https://github.com/cli/cli/releases/download/v2.52.0/gh_2.52.0_linux_386.deb"),
				mustNewAsset("https://github.com/cli/cli/releases/download/v2.52.0/gh_2.52.0_linux_386.rpm"),
				mustNewAsset("https://github.com/cli/cli/releases/download/v2.52.0/gh_2.52.0_linux_386.tar.gz"),
				mustNewAsset("https://github.com/cli/cli/releases/download/v2.52.0/gh_2.52.0_linux_amd64.deb"),
				mustNewAsset("https://github.com/cli/cli/releases/download/v2.52.0/gh_2.52.0_linux_amd64.rpm"),
				mustNewAsset("https://github.com/cli/cli/releases/download/v2.52.0/gh_2.52.0_linux_amd64.tar.gz"),
				mustNewAsset("https://github.com/cli/cli/releases/download/v2.52.0/gh_2.52.0_linux_arm64.deb"),
				mustNewAsset("https://github.com/cli/cli/releases/download/v2.52.0/gh_2.52.0_linux_arm64.rpm"),
				mustNewAsset("https://github.com/cli/cli/releases/download/v2.52.0/gh_2.52.0_linux_arm64.tar.gz"),
				mustNewAsset("https://github.com/cli/cli/releases/download/v2.52.0/gh_2.52.0_linux_armv6.deb"),
				mustNewAsset("https://github.com/cli/cli/releases/download/v2.52.0/gh_2.52.0_linux_armv6.rpm"),
				mustNewAsset("https://github.com/cli/cli/releases/download/v2.52.0/gh_2.52.0_linux_armv6.tar.gz"),
				mustNewAsset("https://github.com/cli/cli/releases/download/v2.52.0/gh_2.52.0_macOS_amd64.zip"),
				mustNewAsset("https://github.com/cli/cli/releases/download/v2.52.0/gh_2.52.0_macOS_arm64.zip"),
				mustNewAsset("https://github.com/cli/cli/releases/download/v2.52.0/gh_2.52.0_macOS_universal.pkg"),
				mustNewAsset("https://github.com/cli/cli/releases/download/v2.52.0/gh_2.52.0_windows_386.msi"),
				mustNewAsset("https://github.com/cli/cli/releases/download/v2.52.0/gh_2.52.0_windows_386.zip"),
				mustNewAsset("https://github.com/cli/cli/releases/download/v2.52.0/gh_2.52.0_windows_amd64.msi"),
				mustNewAsset("https://github.com/cli/cli/releases/download/v2.52.0/gh_2.52.0_windows_amd64.zip"),
				mustNewAsset("https://github.com/cli/cli/releases/download/v2.52.0/gh_2.52.0_windows_arm64.zip"),
			},
			patterns: []*regexp.Regexp{
				regexp.MustCompile(`gh_.+_linux_amd64.tar.gz`),
			},
			asset: mustNewAsset("https://github.com/cli/cli/releases/download/v2.52.0/gh_2.52.0_linux_amd64.tar.gz"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require := require.New(t)
			asset, err := tt.assets.find(tt.patterns)
			require.NoError(err)
			require.Equal(tt.asset, asset)
		})
	}
}

func TestListAssets(t *testing.T) {
	tests := []struct {
		name    string
		repo    Repository
		release Release
		assets  AssetList
	}{
		{
			name:    "cli/cli",
			repo:    newRepository("cli", "cli"),
			release: newRelease("v2.52.0"),
			assets: AssetList{
				mustNewAsset("https://github.com/cli/cli/releases/download/v2.52.0/gh_2.52.0_checksums.txt"),
				mustNewAsset("https://github.com/cli/cli/releases/download/v2.52.0/gh_2.52.0_linux_386.deb"),
				mustNewAsset("https://github.com/cli/cli/releases/download/v2.52.0/gh_2.52.0_linux_386.rpm"),
				mustNewAsset("https://github.com/cli/cli/releases/download/v2.52.0/gh_2.52.0_linux_386.tar.gz"),
				mustNewAsset("https://github.com/cli/cli/releases/download/v2.52.0/gh_2.52.0_linux_amd64.deb"),
				mustNewAsset("https://github.com/cli/cli/releases/download/v2.52.0/gh_2.52.0_linux_amd64.rpm"),
				mustNewAsset("https://github.com/cli/cli/releases/download/v2.52.0/gh_2.52.0_linux_amd64.tar.gz"),
				mustNewAsset("https://github.com/cli/cli/releases/download/v2.52.0/gh_2.52.0_linux_arm64.deb"),
				mustNewAsset("https://github.com/cli/cli/releases/download/v2.52.0/gh_2.52.0_linux_arm64.rpm"),
				mustNewAsset("https://github.com/cli/cli/releases/download/v2.52.0/gh_2.52.0_linux_arm64.tar.gz"),
				mustNewAsset("https://github.com/cli/cli/releases/download/v2.52.0/gh_2.52.0_linux_armv6.deb"),
				mustNewAsset("https://github.com/cli/cli/releases/download/v2.52.0/gh_2.52.0_linux_armv6.rpm"),
				mustNewAsset("https://github.com/cli/cli/releases/download/v2.52.0/gh_2.52.0_linux_armv6.tar.gz"),
				mustNewAsset("https://github.com/cli/cli/releases/download/v2.52.0/gh_2.52.0_macOS_amd64.zip"),
				mustNewAsset("https://github.com/cli/cli/releases/download/v2.52.0/gh_2.52.0_macOS_arm64.zip"),
				mustNewAsset("https://github.com/cli/cli/releases/download/v2.52.0/gh_2.52.0_macOS_universal.pkg"),
				mustNewAsset("https://github.com/cli/cli/releases/download/v2.52.0/gh_2.52.0_windows_386.msi"),
				mustNewAsset("https://github.com/cli/cli/releases/download/v2.52.0/gh_2.52.0_windows_386.zip"),
				mustNewAsset("https://github.com/cli/cli/releases/download/v2.52.0/gh_2.52.0_windows_amd64.msi"),
				mustNewAsset("https://github.com/cli/cli/releases/download/v2.52.0/gh_2.52.0_windows_amd64.zip"),
				mustNewAsset("https://github.com/cli/cli/releases/download/v2.52.0/gh_2.52.0_windows_arm64.zip"),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require := require.New(t)
			ctx := context.Background()
			repository := NewAssetRepository(ctx, os.Getenv("GH_TOKEN"))
			assets, err := repository.list(ctx, tt.repo, tt.release)
			require.NoError(err)
			require.Equal(tt.assets, assets)
		})
	}
}

// mustNewAsset returns a new GitHub release asset object.
// Given download URL must be able to be parsed as URL.
// This gets into a panic if the error is non-nil.
func mustNewAsset(downloadURL string) Asset {
	url, err := url.Parse(downloadURL)
	if err != nil {
		panic(err)
	}
	return newAsset(url)
}
