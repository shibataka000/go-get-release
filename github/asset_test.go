package github

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFindAsset(t *testing.T) {
	tests := []struct {
		name     string
		assets   AssetList
		patterns PatternList
		asset    Asset
	}{
		{
			name: "gh_2.52.0_linux_amd64.tar.gz",
			assets: AssetList{
				newAsset(0, "gh_2.52.0_checksums.txt"),
				newAsset(0, "gh_2.52.0_linux_386.deb"),
				newAsset(0, "gh_2.52.0_linux_386.rpm"),
				newAsset(0, "gh_2.52.0_linux_386.tar.gz"),
				newAsset(0, "gh_2.52.0_linux_amd64.deb"),
				newAsset(0, "gh_2.52.0_linux_amd64.rpm"),
				newAsset(0, "gh_2.52.0_linux_amd64.tar.gz"),
				newAsset(0, "gh_2.52.0_linux_arm64.deb"),
				newAsset(0, "gh_2.52.0_linux_arm64.rpm"),
				newAsset(0, "gh_2.52.0_linux_arm64.tar.gz"),
				newAsset(0, "gh_2.52.0_linux_armv6.deb"),
				newAsset(0, "gh_2.52.0_linux_armv6.rpm"),
				newAsset(0, "gh_2.52.0_linux_armv6.tar.gz"),
				newAsset(0, "gh_2.52.0_macOS_amd64.zip"),
				newAsset(0, "gh_2.52.0_macOS_arm64.zip"),
				newAsset(0, "gh_2.52.0_macOS_universal.pkg"),
				newAsset(0, "gh_2.52.0_windows_386.msi"),
				newAsset(0, "gh_2.52.0_windows_386.zip"),
				newAsset(0, "gh_2.52.0_windows_amd64.msi"),
				newAsset(0, "gh_2.52.0_windows_amd64.zip"),
				newAsset(0, "gh_2.52.0_windows_arm64.zip"),
			},
			patterns: PatternList{},
			asset:    newAsset(0, "gh_2.52.0_linux_amd64.tar.gz"),
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

func TestExtractExecBinaryContentFromAssetContent(t *testing.T) {
	// todo: implement this.
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
				newAsset(0, "gh_2.52.0_checksums.txt"),
				newAsset(0, "gh_2.52.0_linux_386.deb"),
				newAsset(0, "gh_2.52.0_linux_386.rpm"),
				newAsset(0, "gh_2.52.0_linux_386.tar.gz"),
				newAsset(0, "gh_2.52.0_linux_amd64.deb"),
				newAsset(0, "gh_2.52.0_linux_amd64.rpm"),
				newAsset(0, "gh_2.52.0_linux_amd64.tar.gz"),
				newAsset(0, "gh_2.52.0_linux_arm64.deb"),
				newAsset(0, "gh_2.52.0_linux_arm64.rpm"),
				newAsset(0, "gh_2.52.0_linux_arm64.tar.gz"),
				newAsset(0, "gh_2.52.0_linux_armv6.deb"),
				newAsset(0, "gh_2.52.0_linux_armv6.rpm"),
				newAsset(0, "gh_2.52.0_linux_armv6.tar.gz"),
				newAsset(0, "gh_2.52.0_macOS_amd64.zip"),
				newAsset(0, "gh_2.52.0_macOS_arm64.zip"),
				newAsset(0, "gh_2.52.0_macOS_universal.pkg"),
				newAsset(0, "gh_2.52.0_windows_386.msi"),
				newAsset(0, "gh_2.52.0_windows_386.zip"),
				newAsset(0, "gh_2.52.0_windows_amd64.msi"),
				newAsset(0, "gh_2.52.0_windows_amd64.zip"),
				newAsset(0, "gh_2.52.0_windows_arm64.zip"),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require := require.New(t)
			ctx := context.Background()
			repository := NewAssetRepository(ctx, githubToken)
			assets, err := repository.list(ctx, tt.repo, tt.release)
			require.NoError(err)
			require.Equal(tt.assets, assets)
		})
	}
}

func TestDownloadAssetContent(t *testing.T) {
	// todo: implement this.
}

var githubToken = os.Getenv("GITHUB_TOKEN")
