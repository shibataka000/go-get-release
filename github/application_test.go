package github

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestApplicationServiceSearch(t *testing.T) {
	require := require.New(t)

	tests, err := readAssetTestData(t)
	require.NoError(err)

	ctx := context.Background()
	app := NewApplicationService(
		NewAssetRepository(ctx, os.Getenv("GITHUB_TOKEN")),
	)

	for _, tt := range tests {
		name := tt.asset.DownloadURL.String()
		t.Run(name, func(_ *testing.T) {
			repoFullName := fmt.Sprintf("%s/%s", tt.repo.owner, tt.repo.name)
			asset, err := app.FindAsset(ctx, repoFullName, tt.release.tag, tt.os, tt.arch)
			require.NoError(err)
			require.Equal(tt.asset, asset)
		})
	}
}
