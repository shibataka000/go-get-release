package github

import (
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
		if tt.OS != "" {
			name := tt.AssetDownloadURL
			t.Run(name, func(t *testing.T) {
				require := require.New(t)
				downloadURL, err := url.Parse(tt.AssetDownloadURL)
				require.NoError(err)
				asset := newAsset(downloadURL)
				require.Equal(tt.OS, asset.os())
			})
		}
	}
}

func TestAssetArch(t *testing.T) {
	tests, err := readAssetTestCase(t)
	require.NoError(t, err)

	for _, tt := range tests {
		if tt.Arch != "" {
			name := tt.AssetDownloadURL
			t.Run(name, func(t *testing.T) {
				require := require.New(t)
				downloadURL, err := url.Parse(tt.AssetDownloadURL)
				require.NoError(err)
				asset := newAsset(downloadURL)
				require.Equal(tt.Arch, asset.arch())
			})
		}
	}
}

func TestAssetMIME(t *testing.T) {
	tests, err := readAssetTestCase(t)
	require.NoError(t, err)

	for _, tt := range tests {
		if tt.MIME != "" {
			name := tt.AssetDownloadURL
			t.Run(name, func(t *testing.T) {
				require := require.New(t)
				downloadURL, err := url.Parse(tt.AssetDownloadURL)
				require.NoError(err)
				asset := newAsset(downloadURL)
				require.Equal(tt.MIME, asset.mime())
			})
		}
	}
}

func TestAssetHasExecBinary(t *testing.T) {
	tests, err := readAssetTestCase(t)
	require.NoError(t, err)

	for _, tt := range tests {
		name := tt.AssetDownloadURL
		t.Run(name, func(t *testing.T) {
			require := require.New(t)
			downloadURL, err := url.Parse(tt.AssetDownloadURL)
			require.NoError(err)
			asset := newAsset(downloadURL)
			require.Equal(tt.HasExecBinary, asset.hasExecBinary())
		})
	}
}

func TestAssetListFind(t *testing.T) {
	// tests, err := readAssetTestData(t)
	// require.NoError(t, err)

	// for _, tt := range tests {
	// 	if !tt.hasExecBinary {
	// 		continue
	// 	}

	// 	name := tt.asset.DownloadURL.String()
	// 	t.Run(name, func(t *testing.T) {
	// 		require := require.New(t)
	// 		assets := tests.assets(tt.repo, tt.release)
	// 		asset, err := assets.find(tt.os, tt.arch)
	// 		require.NoError(err)
	// 		require.Equal(tt.asset, asset)
	// 	})
	// }
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
	// tests, err := readAssetTestData(t)
	// require.NoError(t, err)

	// for _, repo := range tests.repositories() {
	// 	for _, release := range tests.releases(repo) {
	// 		name := ""
	// 		t.Run(name, func(t *testing.T) {
	// 			require := require.New(t)
	// 			ctx := context.Background()
	// 			r := NewAssetRepository(ctx, os.Getenv("GITHUB_TOKEN"))
	// 			assets, err := r.list(ctx, repo, release)
	// 			require.NoError(err)
	// 			require.Equal(tests.assets(repo, release), assets)

	// 		})
	// 	}
	// }
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
func readAssetTestCase(t *testing.T) ([]*AssetTestCase, error) {
	t.Helper()

	path := filepath.Join(".", "testdata", "assets.csv")
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	tests := []*AssetTestCase{}

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
