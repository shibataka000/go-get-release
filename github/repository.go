package github

import (
	"context"

	"github.com/google/go-github/v48/github"
)

// Repository represents a GitHub repository.
type Repository struct {
	Owner string `yaml:"owner"`
	Name  string `yaml:"name"`
}

// newRepository return new Repository object.
func newRepository(owner string, name string) Repository {
	return Repository{
		Owner: owner,
		Name:  name,
	}
}

// Repository is repository for GitHub repository.
type RepositoryRepository struct {
	client *github.Client
}

// NewRepositoryRepository return new RepositoryRepository object.
func NewRepositoryRepository(ctx context.Context, token string) *RepositoryRepository {
	return &RepositoryRepository{
		client: newGitHubClient(ctx, token),
	}
}

// search repository.
func (r *RepositoryRepository) search(ctx context.Context, query string) (Repository, error) {
	result, _, err := r.client.Search.Repositories(ctx, query, &github.SearchOptions{})
	if err != nil {
		return Repository{}, err
	}
	repos := result.Repositories
	if len(repos) == 0 {
		return Repository{}, &RepositoryNotFoundError{}
	}
	repo := repos[0]
	return newRepository(repo.GetOwner().GetLogin(), repo.GetName()), nil
}
