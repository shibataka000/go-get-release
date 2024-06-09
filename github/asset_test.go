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
		assets AssetList
		os     platform.OS
		arch   platform.Arch
		asset  Asset
	}{
		{
			name: "linux_amd64",
			assets: AssetList{
				newAsset("linux-amd64", ""),
				newAsset("darwin-amd64", ""),
				newAsset("windows-amd64", ""),
			},
			os:    "linux",
			arch:  "amd64",
			asset: newAsset("linux-amd64", ""),
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

func TestListAssets(t *testing.T) {

	tests := []struct {
		name    string
		repo    Repository
		release Release
		assets  AssetList
	}{
		{
			name:    "shibataka000/go-get-release-test",
			repo:    newRepository("shibataka000", "go-get-release-test"),
			release: newRelease("v0.0.2"),
			assets: AssetList{
				newAsset("https://github.com/shibataka000/go-get-release-test/releases/download/v0.0.2/go-get-release_v0.0.2_darwin_amd64", ""),
				newAsset("https://github.com/shibataka000/go-get-release-test/releases/download/v0.0.2/go-get-release_v0.0.2_linux_amd64", ""),
				newAsset("https://github.com/shibataka000/go-get-release-test/releases/download/v0.0.2/go-get-release_v0.0.2_windows_amd64.exe", ""),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require := require.New(t)
			ctx := context.Background()
			repository := NewAssetRepository(ctx, os.Getenv("GITHUB_TOKEN"))
			assets, err := repository.list(ctx, tt.repo, tt.release)
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
		assets  AssetList
	}{
		{
			name:    "hashicorp/terraform",
			repo:    newRepository("hashicorp", "terraform"),
			release: newRelease("v1.0.0"),
			assets: AssetList{
				newAsset("https://releases.hashicorp.com/terraform/1.0.0/terraform_1.0.0_linux_amd64.zip", ""),
				newAsset("https://releases.hashicorp.com/terraform/1.0.0/terraform_1.0.0_darwin_amd64.zip", ""),
				newAsset("https://releases.hashicorp.com/terraform/1.0.0/terraform_1.0.0_windows_amd64.zip", ""),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require := require.New(t)
			ctx := context.Background()
			repository := NewAssetRepository(ctx, os.Getenv("GITHUB_TOKEN"))
			assets, err := repository.listExternal(tt.repo, tt.release)
			require.NoError(err)
			require.Equal(tt.assets, assets)
		})
	}
}
