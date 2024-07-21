package github

import (
	"cmp"
	"context"
	"fmt"
	"net/url"
	"os"
	"path/filepath"
	"slices"
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
		name := tt.Asset.DownloadURL.String()
		t.Run(name, func(t *testing.T) {
			require := require.New(t)
			require.Equal(tt.OS, tt.Asset.os())
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
		name := tt.Asset.DownloadURL.String()
		t.Run(name, func(t *testing.T) {
			require := require.New(t)
			require.Equal(tt.Arch, tt.Asset.arch())
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
		name := tt.Asset.DownloadURL.String()
		t.Run(name, func(t *testing.T) {
			require := require.New(t)
			require.Equal(tt.MIME, tt.Asset.mime())
		})
	}
}

func TestAssetHasExecBinary(t *testing.T) {
	tests, err := readAssetTestCase(t)
	require.NoError(t, err)

	for _, tt := range tests {
		name := tt.Asset.DownloadURL.String()
		t.Run(name, func(t *testing.T) {
			require := require.New(t)
			require.Equal(tt.HasExecBinary, tt.Asset.hasExecBinary())
		})
	}
}

func TestAssetIgnored(t *testing.T) {
	tests, err := readAssetTestCase(t)
	require.NoError(t, err)

	for _, tt := range tests {
		name := tt.Asset.DownloadURL.String()
		t.Run(name, func(t *testing.T) {
			require := require.New(t)
			require.Equal(tt.Ignored, tt.Asset.ignored())
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
		name := tt.Asset.DownloadURL.String()
		t.Run(name, func(t *testing.T) {
			require := require.New(t)
			asset, err := tests.listAssetsByRelease(tt.Repository, tt.Release).find(tt.OS, tt.Arch)
			require.NoError(err)
			require.Equal(tt.Asset, asset)
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

func TestAssetCacheGet(t *testing.T) {
	tests := []struct {
		name    string
		cache   AssetCache
		repo    Repository
		release Release
		v       AssetList
		ok      bool
	}{
		{
			name:    "not hit",
			cache:   AssetCache{},
			repo:    newRepository("hashicorp", "terraform"),
			release: newRelease("v1.8.5"),
			v:       AssetList{},
			ok:      false,
		},
		{
			name: "hit",
			cache: AssetCache{
				newRepository("hashicorp", "terraform"): {
					newRelease("v1.8.5"): AssetList{
						newAsset(newURL("https://releases.hashicorp.com/terraform/1.8.5/terraform_1.8.5_linux_amd64.zip")),
					},
				},
			},
			repo:    newRepository("hashicorp", "terraform"),
			release: newRelease("v1.8.5"),
			v: AssetList{
				newAsset(newURL("https://releases.hashicorp.com/terraform/1.8.5/terraform_1.8.5_linux_amd64.zip")),
			},
			ok: true,
		},
		{
			name: "not hit",
			cache: AssetCache{
				newRepository("hashicorp", "terraform"): {
					newRelease("v1.8.5"): AssetList{
						newAsset(newURL("https://releases.hashicorp.com/terraform/1.8.5/terraform_1.8.5_linux_amd64.zip")),
					},
				},
			},
			repo:    newRepository("hashicorp", "terraform"),
			release: newRelease("v1.8.6"),
			v:       AssetList{},
			ok:      false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require := require.New(t)
			v, ok := tt.cache.get(tt.repo, tt.release)
			require.Equal(tt.ok, ok)
			if ok {
				require.Equal(tt.v, v)
			}
		})
	}
}

func TestAssetCacheSet(t *testing.T) {
	tests := []struct {
		name    string
		before  AssetCache
		after   AssetCache
		repo    Repository
		release Release
		assets  AssetList
	}{
		{
			name:   "Set assets whose repository and release are new",
			before: AssetCache{},
			after: AssetCache{
				newRepository("hashicorp", "terraform"): {
					newRelease("v1.8.5"): AssetList{
						newAsset(newURL("https://releases.hashicorp.com/terraform/1.8.5/terraform_1.8.5_linux_amd64.zip")),
					},
				},
			},
			repo:    newRepository("hashicorp", "terraform"),
			release: newRelease("v1.8.5"),
			assets: AssetList{
				newAsset(newURL("https://releases.hashicorp.com/terraform/1.8.5/terraform_1.8.5_linux_amd64.zip")),
			},
		},
		{
			name: "Set assets whose release is new",
			before: AssetCache{
				newRepository("hashicorp", "terraform"): {
					newRelease("v1.8.4"): AssetList{
						newAsset(newURL("https://releases.hashicorp.com/terraform/1.8.5/terraform_1.8.4_linux_amd64.zip")),
					},
				},
			},
			after: AssetCache{
				newRepository("hashicorp", "terraform"): {
					newRelease("v1.8.4"): AssetList{
						newAsset(newURL("https://releases.hashicorp.com/terraform/1.8.5/terraform_1.8.4_linux_amd64.zip")),
					},
					newRelease("v1.8.5"): AssetList{
						newAsset(newURL("https://releases.hashicorp.com/terraform/1.8.5/terraform_1.8.5_linux_amd64.zip")),
					},
				},
			},
			repo:    newRepository("hashicorp", "terraform"),
			release: newRelease("v1.8.5"),
			assets: AssetList{
				newAsset(newURL("https://releases.hashicorp.com/terraform/1.8.5/terraform_1.8.5_linux_amd64.zip")),
			},
		},
		{
			name: "Set assets whose repository is new",
			before: AssetCache{
				newRepository("hashicorp", "terraform"): {
					newRelease("v1.8.5"): AssetList{
						newAsset(newURL("https://releases.hashicorp.com/terraform/1.8.5/terraform_1.8.5_linux_amd64.zip")),
					},
				},
			},
			after: AssetCache{
				newRepository("hashicorp", "terraform"): {
					newRelease("v1.8.5"): AssetList{
						newAsset(newURL("https://releases.hashicorp.com/terraform/1.8.5/terraform_1.8.5_linux_amd64.zip")),
					},
				},
				newRepository("cli", "cli"): {
					newRelease("v2.52.0"): AssetList{
						newAsset(newURL("https://github.com/cli/cli/releases/download/v2.52.0/gh_2.52.0_linux_amd64.tar.gz")),
					},
				},
			},
			repo:    newRepository("cli", "cli"),
			release: newRelease("v2.52.0"),
			assets: AssetList{
				newAsset(newURL("https://github.com/cli/cli/releases/download/v2.52.0/gh_2.52.0_linux_amd64.tar.gz")),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require := require.New(t)
			cache := tt.before
			cache.set(tt.repo, tt.release, tt.assets)
			require.Equal(tt.after, cache)
		})
	}
}

func TestListAssets(t *testing.T) {
	tests, err := readAssetTestCase(t)
	require.NoError(t, err)

	ctx := context.Background()
	repository := NewAssetRepository(ctx, os.Getenv("GITHUB_TOKEN"))

	done := []string{}

	for _, tt := range tests {
		name := fmt.Sprintf("%s/%s/%s", tt.Repository.owner, tt.Repository.name, tt.Release.tag)
		if slices.Contains(done, name) {
			continue
		}
		t.Run(name, func(t *testing.T) {
			require := require.New(t)

			except := tests.listAssetsByRelease(tt.Repository, tt.Release)
			slices.SortFunc(except, func(a, b Asset) int {
				return cmp.Compare(a.DownloadURL.String(), b.DownloadURL.String())
			})

			actual, err := repository.list(ctx, tt.Repository, tt.Release)
			require.NoError(err)
			slices.SortFunc(actual, func(a, b Asset) int {
				return cmp.Compare(a.DownloadURL.String(), b.DownloadURL.String())
			})

			require.Equal(except, actual)
		})
		done = append(done, name)
	}
}

func TestAssetUnmarshalCSV(t *testing.T) {
	tests := []struct {
		name  string
		value string
		asset Asset
	}{
		{
			name:  "https://github.com/cli/cli/releases/download/v2.52.0/gh_2.52.0_linux_amd64.tar.gz",
			value: "https://github.com/cli/cli/releases/download/v2.52.0/gh_2.52.0_linux_amd64.tar.gz",
			asset: newAsset(newURL("https://github.com/cli/cli/releases/download/v2.52.0/gh_2.52.0_linux_amd64.tar.gz")),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require := require.New(t)
			asset := Asset{}
			asset.UnmarshalCSV(tt.value)
			require.Equal(tt.asset, asset)
		})
	}
}

// UnmarshalCSV converts the CSV string as GitHub release asset.
func (a *Asset) UnmarshalCSV(value string) error {
	url, err := url.Parse(value)
	if err != nil {
		return err
	}
	a.DownloadURL = url
	return nil
}

// AssetTestCase is a test case about a GitHub release asset.
type AssetTestCase struct {
	Repository    Repository `csv:"repository"`
	Release       Release    `csv:"release"`
	Asset         Asset      `csv:"asset"`
	OS            dist.OS    `csv:"os"`
	Arch          dist.Arch  `csv:"arch"`
	MIME          mime.Type  `csv:"mime"`
	HasExecBinary bool       `csv:"has_exec_binary"`
	Ignored       bool       `csv:"ignored"`
	ExecBinary    ExecBinary `csv:"exec_binary"`
}

// AssetTestCaseList is a list of test case about a GitHub release asset.
type AssetTestCaseList []AssetTestCase

// listAssetsByRelease returns a list of GitHub release asset which is contained by given release in GitHub repository.
func (s AssetTestCaseList) listAssetsByRelease(repo Repository, release Release) AssetList {
	assets := AssetList{}
	for _, t := range s {
		if t.Repository == repo && t.Release == release && !slices.Contains(assets, t.Asset) {
			assets = append(assets, t.Asset)
		}
	}
	return assets
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

// newURL parses a raw url into a URL structure.
// This gets into a panic if the error is non-nil.
func newURL(rawURL string) *url.URL {
	parsed, err := url.Parse(rawURL)
	if err != nil {
		panic(err)
	}
	return parsed
}
