package github

import (
	"context"
	"net/url"
	"os"
	"testing"

	"github.com/shibataka000/go-get-release/platform"
	"github.com/stretchr/testify/require"
)

func TestAssetPlatform(t *testing.T) {
	tests := []struct {
		name  string
		asset Asset
		os    platform.OS
		arch  platform.Arch
	}{
		{
			name:  "https://github.com/cli/cli/releases/download/v2.51.1/gh_2.51.1_linux_amd64.tar.gz",
			asset: newAsset(newURL("https://github.com/cli/cli/releases/download/v2.51.1/gh_2.51.1_linux_amd64.tar.gz")),
			os:    "linux",
			arch:  "amd64",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require := require.New(t)
			require.Equal(tt.os, tt.asset.os())
			require.Equal(tt.arch, tt.asset.arch())
		})
	}
}

func TestAssetHasExecBinary(t *testing.T) {
	tests := []struct {
		name          string
		asset         Asset
		hasExecBinary bool
	}{
		{
			name:          "https://github.com/cli/cli/releases/download/v2.51.1/gh_2.51.1_linux_amd64.tar.gz",
			asset:         newAsset(newURL("https://github.com/cli/cli/releases/download/v2.51.1/gh_2.51.1_linux_amd64.tar.gz")),
			hasExecBinary: true,
		},
		{
			name:          "https://github.com/cli/cli/releases/download/v2.51.1/gh_2.51.1_linux_386.deb",
			asset:         newAsset(newURL("https://github.com/cli/cli/releases/download/v2.51.1/gh_2.51.1_linux_386.deb")),
			hasExecBinary: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require := require.New(t)
			require.Equal(tt.hasExecBinary, tt.asset.hasExecBinary())
		})
	}
}

func TestAssetListFind(t *testing.T) {
	tests := []struct {
		name   string
		assets AssetList
		os     platform.OS
		arch   platform.Arch
		asset  Asset
	}{
		{
			name: "https://github.com/cli/cli/releases/download/v2.51.1/gh_2.51.1_linux_amd64.tar.gz",
			assets: AssetList{
				newAsset(newURL("https://github.com/cli/cli/releases/download/v2.51.1/gh_2.51.1_checksums.txt")),
				newAsset(newURL("https://github.com/cli/cli/releases/download/v2.51.1/gh_2.51.1_linux_386.deb")),
				newAsset(newURL("https://github.com/cli/cli/releases/download/v2.51.1/gh_2.51.1_linux_386.rpm")),
				newAsset(newURL("https://github.com/cli/cli/releases/download/v2.51.1/gh_2.51.1_linux_386.tar.gz")),
				newAsset(newURL("https://github.com/cli/cli/releases/download/v2.51.1/gh_2.51.1_linux_amd64.deb")),
				newAsset(newURL("https://github.com/cli/cli/releases/download/v2.51.1/gh_2.51.1_linux_amd64.rpm")),
				newAsset(newURL("https://github.com/cli/cli/releases/download/v2.51.1/gh_2.51.1_linux_amd64.tar.gz")),
				newAsset(newURL("https://github.com/cli/cli/releases/download/v2.51.1/gh_2.51.1_linux_arm64.deb")),
				newAsset(newURL("https://github.com/cli/cli/releases/download/v2.51.1/gh_2.51.1_linux_arm64.rpm")),
				newAsset(newURL("https://github.com/cli/cli/releases/download/v2.51.1/gh_2.51.1_linux_arm64.tar.gz")),
				newAsset(newURL("https://github.com/cli/cli/releases/download/v2.51.1/gh_2.51.1_linux_armv6.deb")),
				newAsset(newURL("https://github.com/cli/cli/releases/download/v2.51.1/gh_2.51.1_linux_armv6.rpm")),
				newAsset(newURL("https://github.com/cli/cli/releases/download/v2.51.1/gh_2.51.1_linux_armv6.tar.gz")),
				newAsset(newURL("https://github.com/cli/cli/releases/download/v2.51.1/gh_2.51.1_macOS_amd64.tar.gz")),
				newAsset(newURL("https://github.com/cli/cli/releases/download/v2.51.1/gh_2.51.1_windows_386.zip")),
				newAsset(newURL("https://github.com/cli/cli/releases/download/v2.51.1/gh_2.51.1_windows_amd64.msi")),
				newAsset(newURL("https://github.com/cli/cli/releases/download/v2.51.1/gh_2.51.1_windows_amd64.zip")),
				newAsset(newURL("https://github.com/cli/cli/releases/download/v2.51.1/gh_2.51.1_windows_arm64.zip")),
			},
			os:    "linux",
			arch:  "amd64",
			asset: newAsset(newURL("https://github.com/cli/cli/releases/download/v2.51.1/gh_2.51.1_linux_amd64.tar.gz")),
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

func TestAssetTemplateExecute(t *testing.T) {
	tests := []struct {
		name     string
		template AssetTemplate
		release  Release
		asset    Asset
	}{
		{
			name:     "https://releases.hashicorp.com/terraform/1.8.5/terraform_1.8.5_linux_amd64.zip",
			template: newAssetTemplate(newTemplate("https://releases.hashicorp.com/terraform/{{.SemVer}}/terraform_{{.SemVer}}_linux_amd64.zip")),
			release:  newRelease("v1.8.5"),
			asset:    newAsset(newURL("https://releases.hashicorp.com/terraform/1.8.5/terraform_1.8.5_linux_amd64.zip")),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require := require.New(t)
			asset, err := tt.template.execute(tt.release)
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
			release: newRelease("v2.51.1"),
			assets: AssetList{
				newAsset(newURL("https://github.com/cli/cli/releases/download/v2.51.1/gh_2.51.1_checksums.txt")),
				newAsset(newURL("https://github.com/cli/cli/releases/download/v2.51.1/gh_2.51.1_linux_386.deb")),
				newAsset(newURL("https://github.com/cli/cli/releases/download/v2.51.1/gh_2.51.1_linux_386.rpm")),
				newAsset(newURL("https://github.com/cli/cli/releases/download/v2.51.1/gh_2.51.1_linux_386.tar.gz")),
				newAsset(newURL("https://github.com/cli/cli/releases/download/v2.51.1/gh_2.51.1_linux_amd64.deb")),
				newAsset(newURL("https://github.com/cli/cli/releases/download/v2.51.1/gh_2.51.1_linux_amd64.rpm")),
				newAsset(newURL("https://github.com/cli/cli/releases/download/v2.51.1/gh_2.51.1_linux_amd64.tar.gz")),
				newAsset(newURL("https://github.com/cli/cli/releases/download/v2.51.1/gh_2.51.1_linux_arm64.deb")),
				newAsset(newURL("https://github.com/cli/cli/releases/download/v2.51.1/gh_2.51.1_linux_arm64.rpm")),
				newAsset(newURL("https://github.com/cli/cli/releases/download/v2.51.1/gh_2.51.1_linux_arm64.tar.gz")),
				newAsset(newURL("https://github.com/cli/cli/releases/download/v2.51.1/gh_2.51.1_linux_armv6.deb")),
				newAsset(newURL("https://github.com/cli/cli/releases/download/v2.51.1/gh_2.51.1_linux_armv6.rpm")),
				newAsset(newURL("https://github.com/cli/cli/releases/download/v2.51.1/gh_2.51.1_linux_armv6.tar.gz")),
				newAsset(newURL("https://github.com/cli/cli/releases/download/v2.51.1/gh_2.51.1_macOS_amd64.tar.gz")),
				newAsset(newURL("https://github.com/cli/cli/releases/download/v2.51.1/gh_2.51.1_windows_386.zip")),
				newAsset(newURL("https://github.com/cli/cli/releases/download/v2.51.1/gh_2.51.1_windows_amd64.msi")),
				newAsset(newURL("https://github.com/cli/cli/releases/download/v2.51.1/gh_2.51.1_windows_amd64.zip")),
				newAsset(newURL("https://github.com/cli/cli/releases/download/v2.51.1/gh_2.51.1_windows_arm64.zip")),
			},
		},
		{
			name:    "hashicorp/terraform",
			repo:    newRepository("hashicorp", "terraform"),
			release: newRelease("v1.8.5"),
			assets: AssetList{
				newAsset(newURL("https://releases.hashicorp.com/terraform/1.8.5/terraform_1.8.5_linux_amd64.zip")),
				newAsset(newURL("https://releases.hashicorp.com/terraform/1.8.5/terraform_1.8.5_darwin_amd64.zip")),
				newAsset(newURL("https://releases.hashicorp.com/terraform/1.8.5/terraform_1.8.5_windows_amd64.zip")),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require := require.New(t)
			ctx := context.Background()
			r := NewAssetRepository(ctx, os.Getenv("GITHUB_TOKEN"))
			assets, err := r.list(ctx, tt.repo, tt.release)
			require.NoError(err)
			require.Equal(tt.assets, assets)
		})
	}
}

// newURL parses a raw url into a URL structure.
// This gets into a panic if the error is non-nil.
func newURL(rawURL string) *url.URL {
	parsed, err := url.Parse(rawURL)
	if err != nil {
		panic(err)
	}
	return parsed
}
