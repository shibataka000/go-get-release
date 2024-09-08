package github

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestApplicationService(t *testing.T) {
	tests := []struct {
		name string

		repoFullName       string
		tag                string
		assetPatterns      []string
		execBinaryPatterns []string

		asset      Asset
		execBinary ExecBinary
	}{
		{},
	}

	ctx := context.Background()
	app := NewApplicationService(
		NewAssetRepository(ctx, githubToken),
		NewExecBinaryRepository(),
	)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require := require.New(t)

			asset, err := app.FindAsset(ctx, tt.repoFullName, tt.tag, tt.assetPatterns)
			require.NoError(err)
			require.Equal(tt.asset, asset)

			execBinary, err := app.FindExecBinary(asset, tt.assetPatterns, tt.execBinaryPatterns)
			require.NoError(err)
			require.Equal(tt.execBinary, execBinary)

			file, err := os.CreateTemp("", "")
		})
	}
}
