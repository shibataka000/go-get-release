package github

import (
	"context"
	"fmt"

	"github.com/google/go-github/v48/github"
	"github.com/shibataka000/go-get-release/internal/domain/model/repository"
)

// SearchRepository search GitHub repository.
func (r *Repository) SearchRepository(ctx context.Context, query string) (repository.Repository, error) {
	result, _, err := r.client.Search.Repositories(ctx, query, &github.SearchOptions{})
	if err != nil {
		return repository.Repository{}, err
	}
	repos := result.Repositories
	if len(repos) == 0 {
		return repository.Repository{}, fmt.Errorf("no repository was found")
	}
	repo := repos[0]
	return repository.New(repo.GetOwner().GetLogin(), repo.GetName()), nil
}

// FindRepository find GitHub repository.
func (r *Repository) FindRepository(ctx context.Context, owner string, name string) (repository.Repository, error) {
	repo, _, err := r.client.Repositories.Get(ctx, owner, name)
	if err != nil {
		return repository.Repository{}, err
	}
	return repository.New(repo.GetOwner().GetLogin(), repo.GetName()), nil
}
