package asset

import (
	"fmt"
	"testing"

	"github.com/shibataka000/go-get-release/internal/domain/model/platform"
	"github.com/shibataka000/go-get-release/internal/domain/model/release"
	"github.com/stretchr/testify/require"
)

func TestNewMetadataFromTemplate(t *testing.T) {
	tests := []struct {
		downloadURLTemplate string
		release             Release
		meta                Metadata
	}{
		{
			downloadURLTemplate: "https://github.com/cli/cli/releases/download/{{.Tag}}/gh_{{.SemVer}}_linux_amd64.tar.gz",
			release:             release.New(0, "v2.21.0"),
			meta:                NewMetadata("https://github.com/cli/cli/releases/download/v2.21.0/gh_2.21.0_linux_amd64.tar.gz"),
		},
	}

	for _, tt := range tests {
		name := tt.downloadURLTemplate
		t.Run(name, func(t *testing.T) {
			assert := require.New(t)
			meta, err := NewMetadataFromTemplate(tt.downloadURLTemplate, tt.release)
			assert.NoError(err)
			assert.Equal(tt.meta, meta)
		})
	}
}

func TestMetadataName(t *testing.T) {
	tests := []struct {
		meta Metadata
		name string
	}{
		{
			meta: NewMetadata("https://github.com/cli/cli/releases/download/v2.21.0/gh_2.21.0_linux_amd64.tar.gz"),
			name: "gh_2.21.0_linux_amd64.tar.gz",
		},
	}

	for _, tt := range tests {
		name := tt.name
		t.Run(name, func(t *testing.T) {
			assert := require.New(t)
			assert.Equal(tt.name, tt.meta.Name())
		})
	}
}

func TestMetadataExt(t *testing.T) {
	tests := []struct {
		meta Metadata
		ext  string
	}{
		{
			meta: NewMetadata("https://github.com/cli/cli/releases/download/v2.21.0/gh_2.21.0_linux_amd64.tar.gz"),
			ext:  ".tar.gz",
		},
		{
			meta: NewMetadata("https://github.com/cli/cli/releases/download/v2.21.0/gh_2.21.0_checksums.txt"),
			ext:  ".txt",
		},
	}

	for i, tt := range tests {
		name := fmt.Sprint(i)
		t.Run(name, func(t *testing.T) {
			assert := require.New(t)
			assert.Equal(tt.ext, tt.meta.ext())
		})
	}
}

func TestMetadataFormat(t *testing.T) {
	tests := []struct {
		meta         Metadata
		isExecBinary bool
		isArchived   bool
		isCompressed bool
	}{
		{
			meta:         NewMetadata("https://github.com/cli/cli/releases/download/v2.21.0/gh_2.21.0_linux_amd64.tar.gz"),
			isExecBinary: false,
			isArchived:   true,
			isCompressed: true,
		},
		{
			meta:         NewMetadata("https://github.com/cli/cli/releases/download/v2.21.0/gh_2.21.0_checksums.txt"),
			isExecBinary: false,
			isArchived:   false,
			isCompressed: false,
		},
	}

	for i, tt := range tests {
		name := fmt.Sprint(i)
		t.Run(name, func(t *testing.T) {
			assert := require.New(t)
			assert.Equal(tt.isExecBinary, tt.meta.IsExecBinary())
			assert.Equal(tt.isArchived, tt.meta.IsArchived())
			assert.Equal(tt.isCompressed, tt.meta.IsCompressed())
		})
	}
}

func TestMetadataPlatform(t *testing.T) {
	tests := []struct {
		meta     Metadata
		platform platform.Platform
	}{
		{
			meta:     NewMetadata("https://github.com/cli/cli/releases/download/v2.21.0/gh_2.21.0_linux_amd64.tar.gz"),
			platform: platform.New("linux", "amd64"),
		},
	}

	for i, tt := range tests {
		name := fmt.Sprint(i)
		t.Run(name, func(t *testing.T) {
			assert := require.New(t)
			platform, err := tt.meta.Platform()
			assert.NoError(err)
			assert.Equal(tt.platform, platform)
		})
	}
}
