package github

import (
	"context"
	"net/url"
	"os"
	"path/filepath"
	"testing"

	"github.com/gocarina/gocsv"
	dist "github.com/shibataka000/go-get-release/distribution"
	"github.com/shibataka000/go-get-release/mime"
	"github.com/stretchr/testify/require"
)

func TestAssetOS(t *testing.T) {
	tests, err := readAssetTestCase(t)
	require.NoError(t, err)

	for _, tt := range tests {
		if tt.OS == "" {
			continue
		}
		name := tt.AssetDownloadURL
		t.Run(name, func(t *testing.T) {
			require := require.New(t)
			asset, err := tt.asset()
			require.NoError(err)
			require.Equal(tt.OS, asset.os())
		})
	}
}

func TestAssetArch(t *testing.T) {
	tests, err := readAssetTestCase(t)
	require.NoError(t, err)

	for _, tt := range tests {
		if tt.Arch == "" {
			continue
		}
		name := tt.AssetDownloadURL
		t.Run(name, func(t *testing.T) {
			require := require.New(t)
			asset, err := tt.asset()
			require.NoError(err)
			require.Equal(tt.Arch, asset.arch())
		})
	}
}

func TestAssetMIME(t *testing.T) {
	tests, err := readAssetTestCase(t)
	require.NoError(t, err)

	for _, tt := range tests {
		if tt.MIME == "" {
			continue
		}
		name := tt.AssetDownloadURL
		t.Run(name, func(t *testing.T) {
			require := require.New(t)
			asset, err := tt.asset()
			require.NoError(err)
			require.Equal(tt.MIME, asset.mime())
		})
	}
}

func TestAssetHasExecBinary(t *testing.T) {
	tests, err := readAssetTestCase(t)
	require.NoError(t, err)

	for _, tt := range tests {
		name := tt.AssetDownloadURL
		t.Run(name, func(t *testing.T) {
			require := require.New(t)
			asset, err := tt.asset()
			require.NoError(err)
			require.Equal(tt.HasExecBinary, asset.hasExecBinary())
		})
	}
}

func TestAssetListFind(t *testing.T) {
	tests, err := readAssetTestCase(t)
	require.NoError(t, err)

	for _, tt := range tests {
		if !tt.HasExecBinary {
			continue
		}
		name := tt.AssetDownloadURL
		t.Run(name, func(t *testing.T) {
			require := require.New(t)

			downloadURL, err := url.Parse(tt.AssetDownloadURL)
			require.NoError(err)
			except := newAsset(downloadURL)

			assets, err := tests.assetsByRelease(tt.repository(), tt.release())
			require.NoError(err)
			actual, err := assets.find(tt.OS, tt.Arch)
			require.NoError(err)

			require.Equal(except, actual)
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

func TestAssetTemplateListExecute(t *testing.T) {
	tests := []struct {
		name      string
		templates AssetTemplateList
		release   Release
		assets    AssetList
	}{
		{
			name:      "https://releases.hashicorp.com/terraform/1.8.5/terraform_1.8.5_linux_amd64.zip",
			templates: externalAssets[newRepository("hashicorp", "terraform")],
			release:   newRelease("v1.8.5"),
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
			assets, err := tt.templates.execute(tt.release)
			require.NoError(err)
			require.Equal(tt.assets, assets)
		})
	}
}

func TestListAssets(t *testing.T) {
	tests, err := readAssetTestCase(t)
	require.NoError(t, err)

	ctx := context.Background()
	r := NewAssetRepository(ctx, os.Getenv("GITHUB_TOKEN"))

	for _, tt := range tests {
		name := tt.AssetDownloadURL
		t.Run(name, func(t *testing.T) {
			require := require.New(t)

			except, err := tests.assetsByRelease(tt.repository(), tt.release())
			require.NoError(err)

			actual, err := r.list(ctx, tt.repository(), tt.release())
			require.NoError(err)

			require.Equal(except, actual)
		})
	}
}

// AssetTestCase is a test case about a GitHub release asset.
type AssetTestCase struct {
	Owner            string    `csv:"owner"`
	Repo             string    `csv:"repo"`
	Tag              string    `csv:"tag"`
	AssetDownloadURL string    `csv:"asset_download_url"`
	OS               dist.OS   `csv:"os"`
	Arch             dist.Arch `csv:"arch"`
	MIME             mime.Type `csv:"mime"`
	HasExecBinary    bool      `csv:"has_exec_binary"`
}

// readAssetTestCase return a list of test case about a GitHub release asset.
func readAssetTestCase(t *testing.T) (AssetTestCaseList, error) {
	t.Helper()

	path := filepath.Join(".", "testdata", "assets.csv")
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	tests := AssetTestCaseList{}
	if err := gocsv.UnmarshalFile(file, &tests); err != nil {
		return nil, err
	}
	return tests, nil
}

// repository returns a GitHub repository.
func (c AssetTestCase) repository() Repository {
	return newRepository(c.Owner, c.Repo)
}

// repository returns a GitHub release.
func (c AssetTestCase) release() Release {
	return newRelease(c.Tag)
}

// asset returns a GitHub release asset.
func (c AssetTestCase) asset() (Asset, error) {
	url, err := url.Parse(c.AssetDownloadURL)
	if err != nil {
		return Asset{}, err
	}
	return newAsset(url), nil
}

// AssetTestCaseList is a list of test case about a GitHub release asset.
type AssetTestCaseList []AssetTestCase

// assetsByRelease returns a list of GitHub release asset which is contained by given release.
func (s AssetTestCaseList) assetsByRelease(repo Repository, release Release) (AssetList, error) {
	assets := AssetList{}
	for _, t := range s {
		if t.repository() == repo && t.release() == release {
			asset, err := t.asset()
			if err != nil {
				return nil, err
			}
			assets = append(assets, asset)
		}
	}
	return assets, nil
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
