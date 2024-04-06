package github

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestURLFileName(t *testing.T) {
	tests := []struct {
		name     string
		url      URL
		filename FileName
	}{
		{
			name:     "gh_2.21.0_linux_amd64.tar.gz",
			url:      "https://github.com/cli/cli/releases/download/v2.21.0/gh_2.21.0_linux_amd64.tar.gz",
			filename: "gh_2.21.0_linux_amd64.tar.gz",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require := require.New(t)
			require.Equal(tt.filename, tt.url.FileName())
		})
	}
}

func TestURLTemplateRenderWithRelease(t *testing.T) {
	tests := []struct {
		name    string
		tmpl    URLTemplate
		release Release
		url     URL
	}{
		{
			name:    "https://github.com/viaduct-ai/kustomize-sops/releases/download/{{.Tag}}/ksops_{{.SemVer}}_Linux_x86_64.tar.gz",
			tmpl:    "https://github.com/viaduct-ai/kustomize-sops/releases/download/{{.Tag}}/ksops_{{.SemVer}}_Linux_x86_64.tar.gz",
			release: NewRelease("v4.1.0"),
			url:     "https://github.com/viaduct-ai/kustomize-sops/releases/download/v4.1.0/ksops_4.1.0_Linux_x86_64.tar.gz",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require := require.New(t)
			url, err := tt.tmpl.RenderWithRelease(tt.release)
			require.NoError(err)
			require.Equal(tt.url, url)
		})
	}
}
