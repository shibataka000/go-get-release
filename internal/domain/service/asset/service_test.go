package asset

import (
	"fmt"
	"testing"

	"github.com/shibataka000/go-get-release/internal/domain/model/asset"
	"github.com/shibataka000/go-get-release/internal/domain/model/platform"
	specification "github.com/shibataka000/go-get-release/internal/domain/specification/asset"
	"github.com/stretchr/testify/require"
)

func TestFileMetadata(t *testing.T) {
	tests := []struct {
		files []asset.Metadata
		spec  specification.FindMetadata
		found asset.Metadata
	}{
		{
			files: []asset.Metadata{
				asset.NewMetadata("https://github.com/cli/cli/releases/download/v2.21.0/gh_2.21.0_linux_amd64.tar.gz"),
				asset.NewMetadata("https://github.com/cli/cli/releases/download/v2.21.0/gh_2.21.0_windows_amd64.zip"),
			},
			spec:  specification.NewFindMetadataByPlatform(platform.New("linux", "amd64")),
			found: asset.NewMetadata("https://github.com/cli/cli/releases/download/v2.21.0/gh_2.21.0_linux_amd64.tar.gz"),
		},
	}

	svc := New()

	for i, tt := range tests {
		name := fmt.Sprint(i)
		t.Run(name, func(t *testing.T) {
			assert := require.New(t)
			file, err := svc.FindMetadata(tt.files, tt.spec)
			assert.NoError(err)
			assert.Equal(tt.found, file)
		})
	}
}
