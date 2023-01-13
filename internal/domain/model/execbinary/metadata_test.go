package execbinary

import (
	"testing"

	"github.com/shibataka000/go-get-release/internal/domain/model/platform"
	"github.com/stretchr/testify/require"
)

func TestMetadataName(t *testing.T) {
	tests := []struct {
		meta           Metadata
		execBinaryName string
	}{
		{
			meta:           NewMetadata("gh", platform.New("linux", "amd64")),
			execBinaryName: "gh",
		},
		{
			meta:           NewMetadata("gh", platform.New("windows", "amd64")),
			execBinaryName: "gh.exe",
		},
	}

	for _, tt := range tests {
		name := tt.execBinaryName
		t.Run(name, func(t *testing.T) {
			assert := require.New(t)
			execBinaryName := tt.meta.Name()
			assert.Equal(tt.execBinaryName, execBinaryName)
		})
	}
}
