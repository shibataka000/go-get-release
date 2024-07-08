package github

import (
	"context"
	"net/url"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAssetOS(t *testing.T) {
	tests, err := readAssetTestData(t)
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
	tests, err := readAssetTestData(t)
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
	tests, err := readAssetTestData(t)
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
	tests, err := readAssetTestData(t)
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
	tests, err := readAssetTestData(t)
	require.NoError(t, err)

	for _, tt := range tests {
		if !tt.hasExecBinary {
			continue
		}

		name := tt.asset.DownloadURL.String()
		t.Run(name, func(t *testing.T) {
			require := require.New(t)
			assets := tests.assets(tt.repo, tt.release)
			asset, err := assets.find(tt.os, tt.arch)
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
	tests, err := readAssetTestData(t)
	require.NoError(t, err)

	for _, repo := range tests.repositories() {
		for _, release := range tests.releases(repo) {
			name := ""
			t.Run(name, func(t *testing.T) {
				require := require.New(t)
				ctx := context.Background()
				r := NewAssetRepository(ctx, os.Getenv("GITHUB_TOKEN"))
				assets, err := r.list(ctx, repo, release)
				require.NoError(err)
				require.Equal(tests.assets(repo, release), assets)

			})
		}
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
