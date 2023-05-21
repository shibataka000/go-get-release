package pkg

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFactoryNewRepository(t *testing.T) {
	tests := []struct {
		name   string
		ghRepo GitHubRepository
		repo   Repository
	}{
		{
			name:   "hashicorp/terraform",
			ghRepo: NewGitHubRepository("hashicorp", "terraform"),
			repo:   NewRepository("hashicorp", "terraform"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := require.New(t)
			factory := NewFactory()
			assert.Equal(tt.repo, factory.NewRepository(tt.ghRepo))
		})
	}
}

func TestFactoryNewRelease(t *testing.T) {
	tests := []struct {
		name      string
		ghRelease GitHubRelease
		release   Release
	}{
		{
			name:      "v0.0.1",
			ghRelease: NewGitHubRelease(0, "v0.0.1"),
			release:   NewRelease("v0.0.1"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := require.New(t)
			factory := NewFactory()
			assert.Equal(tt.release, factory.NewRelease(tt.ghRelease))
		})
	}
}

func TestFactoryNewAssetFromIndex(t *testing.T) {
	tests := []struct {
		name         string
		assetInIndex AssetInIndex
		release      Release
		asset        Asset
	}{
		{
			name:         "hashicorp/terraform",
			assetInIndex: NewAssetInIndex("https://releases.hashicorp.com/terraform/{{.SemVer}}/terraform_{{.SemVer}}_linux_amd64.zip", "linux", "amd64"),
			release:      NewRelease("v1.0.0"),
			asset:        NewAsset("https://releases.hashicorp.com/terraform/1.0.0/terraform_1.0.0_linux_amd64.zip"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := require.New(t)
			factory := NewFactory()
			asset, err := factory.NewAssetFromIndex(tt.assetInIndex, tt.release)
			assert.NoError(err)
			assert.Equal(tt.asset, asset)
		})
	}
}

func TestFactoryNewAssetFromGitHub(t *testing.T) {
	tests := []struct {
		name     string
		ghAssets []GitHubAsset
		platform Platform
		asset    Asset
	}{
		{
			name: "cli/cli",
			ghAssets: []GitHubAsset{
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
			asset:    NewAsset("https://github.com/cli/cli/releases/download/v2.21.1/gh_2.21.1_linux_amd64.tar.gz"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := require.New(t)
			factory := NewFactory()
			asset, err := factory.NewAssetFromGitHub(tt.ghAssets, tt.platform)
			assert.NoError(err)
			assert.Equal(tt.asset, asset)
		})
	}
}

func TestFactoryNewExecBinaryFromIndex(t *testing.T) {
	tests := []struct {
		name              string
		execBinaryInIndex ExecBinaryInIndex
		platform          Platform
		execBinary        ExecBinary
	}{
		{
			name:              "terraform",
			execBinaryInIndex: NewExecBinaryInIndex("terraform"),
			platform:          NewPlatform("linux", "amd64"),
			execBinary:        NewExecBinary("terraform"),
		},
		{
			name:              "terraform.exe",
			execBinaryInIndex: NewExecBinaryInIndex("terraform"),
			platform:          NewPlatform("windows", "amd64"),
			execBinary:        NewExecBinary("terraform.exe"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := require.New(t)
			factory := NewFactory()
			execBinary := factory.NewExecBinaryFromIndex(tt.execBinaryInIndex, tt.platform)
			assert.Equal(tt.execBinary, execBinary)
		})
	}
}

func TestFactoryNewExecBinaryFromGitHub(t *testing.T) {
	tests := []struct {
		name       string
		ghRepo     GitHubRepository
		platform   Platform
		execBinary ExecBinary
	}{
		{
			name:       "terraform",
			ghRepo:     NewGitHubRepository("hashicorp", "terraform"),
			platform:   NewPlatform("linux", "amd64"),
			execBinary: NewExecBinary("terraform"),
		},
		{
			name:       "terraform.exe",
			ghRepo:     NewGitHubRepository("hashicorp", "terraform"),
			platform:   NewPlatform("windows", "amd64"),
			execBinary: NewExecBinary("terraform.exe"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := require.New(t)
			factory := NewFactory()
			execBinary := factory.NewExecBinaryFromGitHub(tt.ghRepo, tt.platform)
			assert.Equal(tt.execBinary, execBinary)
		})
	}
}

func TestFactoryNewExecBinaryWithPlatform(t *testing.T) {
	tests := []struct {
		name       string
		baseName   FileName
		platform   Platform
		execBinary ExecBinary
	}{
		{
			name:       "terraform",
			baseName:   NewFileName("terraform"),
			platform:   NewPlatform("linux", "amd64"),
			execBinary: NewExecBinary("terraform"),
		},
		{
			name:       "terraform.exe",
			baseName:   NewFileName("terraform"),
			platform:   NewPlatform("windows", "amd64"),
			execBinary: NewExecBinary("terraform.exe"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := require.New(t)
			factory := NewFactory()
			execBinary := factory.NewExecBinaryWithPlatform(tt.baseName, tt.platform)
			assert.Equal(tt.execBinary, execBinary)
		})
	}
}
