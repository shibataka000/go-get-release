package github

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestApplicationServiceSearch(t *testing.T) {
	tests, err := readAssetTestCase(t)
	require.NoError(t, err)

	ctx := context.Background()
	app := NewApplicationService(
		NewAssetRepository(ctx, os.Getenv("GITHUB_TOKEN")),
	)

	for _, tt := range tests {
		if !tt.HasExecBinary {
			continue
		}
		name := tt.AssetDownloadURL
		t.Run(name, func(_ *testing.T) {
			require := require.New(t)

			except, err := tt.asset()
			require.NoError(err)

			repoFullName := fmt.Sprintf("%s/%s", tt.Owner, tt.Repo)
			actual, err := app.FindAsset(ctx, repoFullName, tt.Tag, tt.OS, tt.Arch)
			require.NoError(err)

			require.Equal(except, actual)
		})
	}
}
