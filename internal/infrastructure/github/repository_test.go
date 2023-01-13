package github

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/shibataka000/go-get-release/internal/domain/model/repository"
	"github.com/stretchr/testify/require"
)

func TestSearchRepository(t *testing.T) {
	tests := []struct {
		query      string
		repository repository.Repository
	}{
		{
			query:      "terraform",
			repository: repository.New("hashicorp", "terraform"),
		},
	}

	ctx := context.Background()
	token := os.Getenv("GITHUB_TOKEN")
	repositoryRepository := New(ctx, token)

	for _, tt := range tests {
		name := tt.query
		t.Run(name, func(t *testing.T) {
			assert := require.New(t)
			repo, err := repositoryRepository.SearchRepository(ctx, tt.query)
			assert.NoError(err)
			assert.Equal(tt.repository, repo)
		})
	}
}

func TestFindRepository(t *testing.T) {
	tests := []struct {
		owner      string
		name       string
		repository repository.Repository
	}{
		{
			owner:      "hashicorp",
			name:       "terraform",
			repository: repository.New("hashicorp", "terraform"),
		},
	}

	ctx := context.Background()
	token := os.Getenv("GITHUB_TOKEN")
	repositoryRepository := New(ctx, token)

	for _, tt := range tests {
		name := fmt.Sprintf("%s/%s", tt.owner, tt.name)
		t.Run(name, func(t *testing.T) {
			assert := require.New(t)
			repo, err := repositoryRepository.FindRepository(ctx, tt.owner, tt.name)
			assert.NoError(err)
			assert.Equal(tt.repository, repo)
		})
	}
}
