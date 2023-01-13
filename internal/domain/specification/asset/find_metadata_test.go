package asset

import (
	"fmt"
	"testing"

	"github.com/shibataka000/go-get-release/internal/domain/model/asset"
	"github.com/shibataka000/go-get-release/internal/domain/model/platform"
	"github.com/stretchr/testify/require"
)

func TestFindMetadataSpecificationIsSatisfied(t *testing.T) {
	tests := []struct {
		spec      FindMetadataByPlatform
		meta      asset.Metadata
		satisfied bool
	}{
		{
			spec:      NewFindMetadataByPlatform(platform.New("linux", "amd64")),
			meta:      asset.NewMetadata("https://github.com/cli/cli/releases/download/v2.21.0/gh_2.21.0_linux_amd64.tar.gz"),
			satisfied: true,
		},
		{
			spec:      NewFindMetadataByPlatform(platform.New("linux", "amd64")),
			meta:      asset.NewMetadata("https://github.com/cli/cli/releases/download/v2.21.0/gh_2.21.0_windows_amd64.zip"),
			satisfied: false,
		},
	}

	for i, tt := range tests {
		name := fmt.Sprint(i)
		t.Run(name, func(t *testing.T) {
			assert := require.New(t)
			satisfied, err := tt.spec.IsSatisfied(tt.meta)
			assert.NoError(err)
			assert.Equal(tt.satisfied, satisfied)
		})
	}
}
