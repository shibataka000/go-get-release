package github

import (
	"context"
	"os"
	"testing"

	"github.com/shibataka000/go-get-release/platform"
	"github.com/stretchr/testify/require"
)

func TestFindAssetMeta(t *testing.T) {
	tests := []struct {
		name   string
		assets AssetMetaList
		os     platform.OS
		arch   platform.Arch
		asset  AssetMeta
	}{
		{
			name: "linux_amd64",
			assets: AssetMetaList{
				newAssetMeta("", "linux", "amd64"),
				newAssetMeta("", "darwin", "amd64"),
				newAssetMeta("", "windows", "amd64"),
			},
			os:    "linux",
			arch:  "amd64",
			asset: newAssetMeta("", "linux", "amd64"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require := require.New(t)
			asset, err := tt.assets.find(tt.os, tt.arch)
			require.NoError(err)
			require.Equal(tt.asset, asset)
		})
	}
}

func TestListAssetsFromAPI(t *testing.T) {
	tests := []struct {
		name    string
		repo    Repository
		release Release
		assets  AssetMetaList
	}{
		{
			name:    "shibataka000/go-get-release-test",
			repo:    newRepository("shibataka000", "go-get-release-test"),
			release: newRelease("v0.0.2"),
			assets: AssetMetaList{
				newAssetMeta("https://github.com/shibataka000/go-get-release-test/releases/download/v0.0.2/go-get-release_v0.0.2_darwin_amd64", "darwin", "amd64"),
				newAssetMeta("https://github.com/shibataka000/go-get-release-test/releases/download/v0.0.2/go-get-release_v0.0.2_linux_amd64", "linux", "amd64"),
				newAssetMeta("https://github.com/shibataka000/go-get-release-test/releases/download/v0.0.2/go-get-release_v0.0.2_windows_amd64.exe", "windows", "amd64"),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require := require.New(t)
			ctx := context.Background()
			repository := NewAssetRepository(ctx, os.Getenv("GITHUB_TOKEN"))
			assets, err := repository.listFromAPI(ctx, tt.repo, tt.release)
			require.NoError(err)
			require.Equal(tt.assets, assets)
		})
	}
}

func TestListAssetsFromBuiltIn(t *testing.T) {
	tests := []struct {
		name    string
		repo    Repository
		release Release
		assets  AssetMetaList
	}{
		{
			name:    "hashicorp/terraform",
			repo:    newRepository("hashicorp", "terraform"),
			release: newRelease("v1.0.0"),
			assets: AssetMetaList{
				newAssetMeta("https://releases.hashicorp.com/terraform/1.0.0/terraform_1.0.0_linux_amd64.zip", "linux", "amd64"),
				newAssetMeta("https://releases.hashicorp.com/terraform/1.0.0/terraform_1.0.0_darwin_amd64.zip", "darwin", "amd64"),
				newAssetMeta("https://releases.hashicorp.com/terraform/1.0.0/terraform_1.0.0_windows_amd64.zip", "windows", "amd64"),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require := require.New(t)
			ctx := context.Background()
			repository := NewAssetRepository(ctx, os.Getenv("GITHUB_TOKEN"))
			assets, err := repository.listFromBuiltIn(tt.repo, tt.release)
			require.NoError(err)
			require.Equal(tt.assets, assets)
		})
	}
}
