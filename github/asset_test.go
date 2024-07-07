package github

import (
	"context"
	"encoding/csv"
	"fmt"
	"io"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"testing"

	"github.com/shibataka000/go-get-release/mime"
	"github.com/shibataka000/go-get-release/platform"
	"github.com/stretchr/testify/require"
)

func TestAssetOS(t *testing.T) {
	tests, err := readAssetTestCase(t)
	require.NoError(t, err)

	for _, tt := range tests {
		name := tt.asset.DownloadURL.String()
		t.Run(name, func(t *testing.T) {
			require := require.New(t)
			if tt.os != "" {
				require.Equal(tt.os, tt.asset.os())
			}
		})
	}
}

func TestAssetArch(t *testing.T) {
	tests, err := readAssetTestCase(t)
	require.NoError(t, err)

	for _, tt := range tests {
		name := tt.asset.DownloadURL.String()
		t.Run(name, func(t *testing.T) {
			require := require.New(t)
			if tt.arch != "" {
				require.Equal(tt.arch, tt.asset.arch())
			}
		})
	}
}

func TestAssetMIME(t *testing.T) {
	tests, err := readAssetTestCase(t)
	require.NoError(t, err)

	for _, tt := range tests {
		name := tt.asset.DownloadURL.String()
		t.Run(name, func(t *testing.T) {
			require := require.New(t)
			if tt.mime != "" {
				require.Equal(tt.mime, tt.asset.mime())
			}
		})
	}
}

func TestAssetHasExecBinary(t *testing.T) {
	tests, err := readAssetTestCase(t)
	require.NoError(t, err)

	for _, tt := range tests {
		name := tt.asset.DownloadURL.String()
		t.Run(name, func(t *testing.T) {
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

// AssetTestCase is a test case about a GitHub release asset.
type AssetTestCase struct {
	repo          Repository
	release       Release
	asset         Asset
	os            platform.OS
	arch          platform.Arch
	mime          mime.MIME
	hasExecBinary bool
}

// readAssetTestCase return a test case about a GitHub release asset.
func readAssetTestCase(t *testing.T) ([]AssetTestCase, error) {
	t.Helper()

	path := filepath.Join(".", "testdata", "assets.csv")
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	r := csv.NewReader(file)

	tests := []AssetTestCase{}
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if strings.HasPrefix(record[0], "#") {
			continue
		}

		if len(record) != 8 {
			return nil, &InvalidAssetTestCaseError{record}
		}

		repo := newRepository(record[0], record[1])
		release := newRelease(record[2])
		downloadURL, err := url.Parse(record[3])
		if err != nil {
			return nil, err
		}
		asset := newAsset(downloadURL)
		os := platform.OS(record[4])
		arch := platform.Arch(record[5])
		mime := mime.MIME(record[6])
		hasExecBinary, err := strconv.ParseBool(record[7])
		if err != nil {
			return nil, err
		}

		tests = append(tests, AssetTestCase{
			repo:          repo,
			release:       release,
			asset:         asset,
			os:            os,
			arch:          arch,
			mime:          mime,
			hasExecBinary: hasExecBinary,
		})
	}
	return tests, nil
}

// InvalidAssetTestCaseError is error raised when asset test case is invalid.
type InvalidAssetTestCaseError struct {
	record []string
}

func (e *InvalidAssetTestCaseError) Error() string {
	return fmt.Sprintf("Asset test case is invalid: %v", e.record)
}
