package github

import (
	"context"
	"net/url"
	"os"
	"testing"

	"github.com/shibataka000/go-get-release/mime"
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
			name:  "https://github.com/cli/cli/releases/download/v2.21.1/gh_2.21.1_linux_amd64.tar.gz",
			asset: newAsset(newURL("https://github.com/cli/cli/releases/download/v2.21.1/gh_2.21.1_linux_amd64.tar.gz"), mime.Gz),
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

func TestAssetHasExecutableBinary(t *testing.T) {
	tests := []struct {
		name                string
		asset               Asset
		hasExecutableBinary bool
	}{
		{
			name:                "https://github.com/cli/cli/releases/download/v2.21.1/gh_2.21.1_linux_amd64.tar.gz",
			asset:               newAsset(newURL("https://github.com/cli/cli/releases/download/v2.21.1/gh_2.21.1_linux_amd64.tar.gz"), mime.Gz),
			hasExecutableBinary: true,
		},
		{
			name:                "https://github.com/cli/cli/releases/download/v2.21.1/gh_2.21.1_checksums.txt",
			asset:               newAsset(newURL("https://github.com/cli/cli/releases/download/v2.21.1/gh_2.21.1_checksums.txt"), mime.Txt),
			hasExecutableBinary: false,
		},
		{
			name:                "https://github.com/cli/cli/releases/download/v2.21.1/gh_2.21.1_linux_386.deb",
			asset:               newAsset(newURL("https://github.com/cli/cli/releases/download/v2.21.1/gh_2.21.1_linux_386.deb"), mime.Deb),
			hasExecutableBinary: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require := require.New(t)
			require.Equal(tt.hasExecutableBinary, tt.asset.hasExecutableBinary())
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
			name: "FindGitHubCLIReleaseAssetForLinuxAndAmd64",
			assets: AssetList{
				newAsset(newURL("https://github.com/cli/cli/releases/download/v2.21.1/gh_2.21.1_checksums.txt"), mime.Txt),
				newAsset(newURL("https://github.com/cli/cli/releases/download/v2.21.1/gh_2.21.1_linux_386.deb"), mime.Deb),
				newAsset(newURL("https://github.com/cli/cli/releases/download/v2.21.1/gh_2.21.1_linux_386.rpm"), mime.Rpm),
				newAsset(newURL("https://github.com/cli/cli/releases/download/v2.21.1/gh_2.21.1_linux_386.tar.gz"), mime.Gz),
				newAsset(newURL("https://github.com/cli/cli/releases/download/v2.21.1/gh_2.21.1_linux_amd64.deb"), mime.Deb),
				newAsset(newURL("https://github.com/cli/cli/releases/download/v2.21.1/gh_2.21.1_linux_amd64.rpm"), mime.Rpm),
				newAsset(newURL("https://github.com/cli/cli/releases/download/v2.21.1/gh_2.21.1_linux_amd64.tar.gz"), mime.Gz),
				newAsset(newURL("https://github.com/cli/cli/releases/download/v2.21.1/gh_2.21.1_linux_arm64.deb"), mime.Deb),
				newAsset(newURL("https://github.com/cli/cli/releases/download/v2.21.1/gh_2.21.1_linux_arm64.rpm"), mime.Rpm),
				newAsset(newURL("https://github.com/cli/cli/releases/download/v2.21.1/gh_2.21.1_linux_arm64.tar.gz"), mime.Gz),
				newAsset(newURL("https://github.com/cli/cli/releases/download/v2.21.1/gh_2.21.1_linux_armv6.deb"), mime.Deb),
				newAsset(newURL("https://github.com/cli/cli/releases/download/v2.21.1/gh_2.21.1_linux_armv6.rpm"), mime.Rpm),
				newAsset(newURL("https://github.com/cli/cli/releases/download/v2.21.1/gh_2.21.1_linux_armv6.tar.gz"), mime.Gz),
				newAsset(newURL("https://github.com/cli/cli/releases/download/v2.21.1/gh_2.21.1_macOS_amd64.tar.gz"), mime.Gz),
				newAsset(newURL("https://github.com/cli/cli/releases/download/v2.21.1/gh_2.21.1_windows_386.zip"), mime.Zip),
				newAsset(newURL("https://github.com/cli/cli/releases/download/v2.21.1/gh_2.21.1_windows_amd64.msi"), mime.Msi),
				newAsset(newURL("https://github.com/cli/cli/releases/download/v2.21.1/gh_2.21.1_windows_amd64.zip"), mime.Zip),
				newAsset(newURL("https://github.com/cli/cli/releases/download/v2.21.1/gh_2.21.1_windows_arm64.zip"), mime.Zip),
			},
			os:    "linux",
			arch:  "amd64",
			asset: newAsset(newURL("https://github.com/cli/cli/releases/download/v2.21.1/gh_2.21.1_linux_amd64.tar.gz"), mime.Gz),
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
			template: newAssetTemplate(newTemplate("https://releases.hashicorp.com/terraform/{{.SemVer}}/terraform_{{.SemVer}}_linux_amd64.zip"), mime.Zip),
			release:  newRelease("v1.8.5"),
			asset:    newAsset(newURL("https://releases.hashicorp.com/terraform/1.8.5/terraform_1.8.5_linux_amd64.zip"), mime.Zip),
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
			release: newRelease("v2.21.1"),
			assets: AssetList{
				newAsset(newURL("https://github.com/cli/cli/releases/download/v2.21.1/gh_2.21.1_checksums.txt"), mime.Txt),
				newAsset(newURL("https://github.com/cli/cli/releases/download/v2.21.1/gh_2.21.1_linux_386.deb"), mime.Deb),
				newAsset(newURL("https://github.com/cli/cli/releases/download/v2.21.1/gh_2.21.1_linux_386.rpm"), mime.Rpm),
				newAsset(newURL("https://github.com/cli/cli/releases/download/v2.21.1/gh_2.21.1_linux_386.tar.gz"), mime.Gz),
				newAsset(newURL("https://github.com/cli/cli/releases/download/v2.21.1/gh_2.21.1_linux_amd64.deb"), mime.Deb),
				newAsset(newURL("https://github.com/cli/cli/releases/download/v2.21.1/gh_2.21.1_linux_amd64.rpm"), mime.Rpm),
				newAsset(newURL("https://github.com/cli/cli/releases/download/v2.21.1/gh_2.21.1_linux_amd64.tar.gz"), mime.Gz),
				newAsset(newURL("https://github.com/cli/cli/releases/download/v2.21.1/gh_2.21.1_linux_arm64.deb"), mime.Deb),
				newAsset(newURL("https://github.com/cli/cli/releases/download/v2.21.1/gh_2.21.1_linux_arm64.rpm"), mime.Rpm),
				newAsset(newURL("https://github.com/cli/cli/releases/download/v2.21.1/gh_2.21.1_linux_arm64.tar.gz"), mime.Gz),
				newAsset(newURL("https://github.com/cli/cli/releases/download/v2.21.1/gh_2.21.1_linux_armv6.deb"), mime.Deb),
				newAsset(newURL("https://github.com/cli/cli/releases/download/v2.21.1/gh_2.21.1_linux_armv6.rpm"), mime.Rpm),
				newAsset(newURL("https://github.com/cli/cli/releases/download/v2.21.1/gh_2.21.1_linux_armv6.tar.gz"), mime.Gz),
				newAsset(newURL("https://github.com/cli/cli/releases/download/v2.21.1/gh_2.21.1_macOS_amd64.tar.gz"), mime.Gz),
				newAsset(newURL("https://github.com/cli/cli/releases/download/v2.21.1/gh_2.21.1_windows_386.zip"), mime.Zip),
				newAsset(newURL("https://github.com/cli/cli/releases/download/v2.21.1/gh_2.21.1_windows_amd64.msi"), mime.Msi),
				newAsset(newURL("https://github.com/cli/cli/releases/download/v2.21.1/gh_2.21.1_windows_amd64.zip"), mime.Zip),
				newAsset(newURL("https://github.com/cli/cli/releases/download/v2.21.1/gh_2.21.1_windows_arm64.zip"), mime.Zip),
			},
		},
		{
			name:    "hashicorp/terraform",
			repo:    newRepository("hashicorp", "terraform"),
			release: newRelease("v1.8.0"),
			assets: AssetList{
				newAsset(newURL("https://releases.hashicorp.com/terraform/1.8.0/terraform_1.8.0_linux_amd64.zip"), mime.Zip),
				newAsset(newURL("https://releases.hashicorp.com/terraform/1.8.0/terraform_1.8.0_darwin_amd64.zip"), mime.Zip),
				newAsset(newURL("https://releases.hashicorp.com/terraform/1.8.0/terraform_1.8.0_windows_amd64.zip"), mime.Zip),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require := require.New(t)
			ctx := context.Background()
			r := NewAssetRepository(ctx, os.Getenv("GITHUB_TOKEN"), 0)
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
