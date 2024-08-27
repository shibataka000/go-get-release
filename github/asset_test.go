package github

import (
	"context"
	"os"
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
			asset: mustNewAssetFromString("https://github.com/cli/cli/releases/download/v2.52.0/gh_2.52.0_linux_amd64.tar.gz"),
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
		patterns PatternList
		asset    Asset
	}{
		{
			name: "gh_2.52.0_linux_amd64.tar.gz",
			assets: AssetList{
				mustNewAssetFromString("https://github.com/cli/cli/releases/download/v2.52.0/gh_2.52.0_checksums.txt"),
				mustNewAssetFromString("https://github.com/cli/cli/releases/download/v2.52.0/gh_2.52.0_linux_386.deb"),
				mustNewAssetFromString("https://github.com/cli/cli/releases/download/v2.52.0/gh_2.52.0_linux_386.rpm"),
				mustNewAssetFromString("https://github.com/cli/cli/releases/download/v2.52.0/gh_2.52.0_linux_386.tar.gz"),
				mustNewAssetFromString("https://github.com/cli/cli/releases/download/v2.52.0/gh_2.52.0_linux_amd64.deb"),
				mustNewAssetFromString("https://github.com/cli/cli/releases/download/v2.52.0/gh_2.52.0_linux_amd64.rpm"),
				mustNewAssetFromString("https://github.com/cli/cli/releases/download/v2.52.0/gh_2.52.0_linux_amd64.tar.gz"),
				mustNewAssetFromString("https://github.com/cli/cli/releases/download/v2.52.0/gh_2.52.0_linux_arm64.deb"),
				mustNewAssetFromString("https://github.com/cli/cli/releases/download/v2.52.0/gh_2.52.0_linux_arm64.rpm"),
				mustNewAssetFromString("https://github.com/cli/cli/releases/download/v2.52.0/gh_2.52.0_linux_arm64.tar.gz"),
				mustNewAssetFromString("https://github.com/cli/cli/releases/download/v2.52.0/gh_2.52.0_linux_armv6.deb"),
				mustNewAssetFromString("https://github.com/cli/cli/releases/download/v2.52.0/gh_2.52.0_linux_armv6.rpm"),
				mustNewAssetFromString("https://github.com/cli/cli/releases/download/v2.52.0/gh_2.52.0_linux_armv6.tar.gz"),
				mustNewAssetFromString("https://github.com/cli/cli/releases/download/v2.52.0/gh_2.52.0_macOS_amd64.zip"),
				mustNewAssetFromString("https://github.com/cli/cli/releases/download/v2.52.0/gh_2.52.0_macOS_arm64.zip"),
				mustNewAssetFromString("https://github.com/cli/cli/releases/download/v2.52.0/gh_2.52.0_macOS_universal.pkg"),
				mustNewAssetFromString("https://github.com/cli/cli/releases/download/v2.52.0/gh_2.52.0_windows_386.msi"),
				mustNewAssetFromString("https://github.com/cli/cli/releases/download/v2.52.0/gh_2.52.0_windows_386.zip"),
				mustNewAssetFromString("https://github.com/cli/cli/releases/download/v2.52.0/gh_2.52.0_windows_amd64.msi"),
				mustNewAssetFromString("https://github.com/cli/cli/releases/download/v2.52.0/gh_2.52.0_windows_amd64.zip"),
				mustNewAssetFromString("https://github.com/cli/cli/releases/download/v2.52.0/gh_2.52.0_windows_arm64.zip"),
			},
			patterns: nil,
			asset:    mustNewAssetFromString("https://github.com/cli/cli/releases/download/v2.52.0/gh_2.52.0_linux_amd64.tar.gz"),
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
				mustNewAssetFromString("https://github.com/cli/cli/releases/download/v2.52.0/gh_2.52.0_checksums.txt"),
				mustNewAssetFromString("https://github.com/cli/cli/releases/download/v2.52.0/gh_2.52.0_linux_386.deb"),
				mustNewAssetFromString("https://github.com/cli/cli/releases/download/v2.52.0/gh_2.52.0_linux_386.rpm"),
				mustNewAssetFromString("https://github.com/cli/cli/releases/download/v2.52.0/gh_2.52.0_linux_386.tar.gz"),
				mustNewAssetFromString("https://github.com/cli/cli/releases/download/v2.52.0/gh_2.52.0_linux_amd64.deb"),
				mustNewAssetFromString("https://github.com/cli/cli/releases/download/v2.52.0/gh_2.52.0_linux_amd64.rpm"),
				mustNewAssetFromString("https://github.com/cli/cli/releases/download/v2.52.0/gh_2.52.0_linux_amd64.tar.gz"),
				mustNewAssetFromString("https://github.com/cli/cli/releases/download/v2.52.0/gh_2.52.0_linux_arm64.deb"),
				mustNewAssetFromString("https://github.com/cli/cli/releases/download/v2.52.0/gh_2.52.0_linux_arm64.rpm"),
				mustNewAssetFromString("https://github.com/cli/cli/releases/download/v2.52.0/gh_2.52.0_linux_arm64.tar.gz"),
				mustNewAssetFromString("https://github.com/cli/cli/releases/download/v2.52.0/gh_2.52.0_linux_armv6.deb"),
				mustNewAssetFromString("https://github.com/cli/cli/releases/download/v2.52.0/gh_2.52.0_linux_armv6.rpm"),
				mustNewAssetFromString("https://github.com/cli/cli/releases/download/v2.52.0/gh_2.52.0_linux_armv6.tar.gz"),
				mustNewAssetFromString("https://github.com/cli/cli/releases/download/v2.52.0/gh_2.52.0_macOS_amd64.zip"),
				mustNewAssetFromString("https://github.com/cli/cli/releases/download/v2.52.0/gh_2.52.0_macOS_arm64.zip"),
				mustNewAssetFromString("https://github.com/cli/cli/releases/download/v2.52.0/gh_2.52.0_macOS_universal.pkg"),
				mustNewAssetFromString("https://github.com/cli/cli/releases/download/v2.52.0/gh_2.52.0_windows_386.msi"),
				mustNewAssetFromString("https://github.com/cli/cli/releases/download/v2.52.0/gh_2.52.0_windows_386.zip"),
				mustNewAssetFromString("https://github.com/cli/cli/releases/download/v2.52.0/gh_2.52.0_windows_amd64.msi"),
				mustNewAssetFromString("https://github.com/cli/cli/releases/download/v2.52.0/gh_2.52.0_windows_amd64.zip"),
				mustNewAssetFromString("https://github.com/cli/cli/releases/download/v2.52.0/gh_2.52.0_windows_arm64.zip"),
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

// mustNewAssetFromString returns a new GitHub release asset object.
// Given download URL must be able to be parsed as URL.
// This gets into a panic if the error is non-nil.
func mustNewAssetFromString(downloadURL string) Asset {
	asset, err := newAssetFromString(downloadURL)
	if err != nil {
		panic(err)
	}
	return asset
}
