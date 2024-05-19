package github

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSearchRepository(t *testing.T) {
	tests := []struct {
		name  string
		query string
		repo  Repository
	}{
		{
			name:  "terraform",
			query: "terraform",
			repo:  newRepository("hashicorp", "terraform"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require := require.New(t)
			ctx := context.Background()
			repository := NewRepositoryRepository(ctx, os.Getenv("GITHUB_TOKEN"))
			repo, err := repository.search(ctx, "terraform")
			require.NoError(err)
			require.Equal(tt.repo, repo)
		})
	}
}
