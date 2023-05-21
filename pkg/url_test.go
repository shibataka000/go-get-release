package pkg

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
			url:      NewURL("https://github.com/cli/cli/releases/download/v2.21.0/gh_2.21.0_linux_amd64.tar.gz"),
			filename: NewFileName("gh_2.21.0_linux_amd64.tar.gz"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := require.New(t)
			assert.Equal(tt.filename, tt.url.FileName())
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
			tmpl:    NewURLTemplate("https://github.com/viaduct-ai/kustomize-sops/releases/download/{{.Tag}}/ksops_{{.SemVer}}_Linux_x86_64.tar.gz"),
			release: NewRelease("v4.1.0"),
			url:     NewURL("https://github.com/viaduct-ai/kustomize-sops/releases/download/v4.1.0/ksops_4.1.0_Linux_x86_64.tar.gz"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := require.New(t)
			url, err := tt.tmpl.RenderWithRelease(tt.release)
			assert.NoError(err)
			assert.Equal(tt.url, url)
		})
	}
}
