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

// Repository is repository for a GitHub repository.
type RepositoryRepository struct {
	client  *github.Client
	factory *RepositoryFactory
}

// NewRepositoryRepository return new RepositoryRepository object.
func NewRepositoryRepository(ctx context.Context, token string, factory *RepositoryFactory) *RepositoryRepository {
	return &RepositoryRepository{
		client:  newGitHubClient(ctx, token),
		factory: factory,
	}
}

// search GitHub repository.
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
	return r.factory.new(repo.GetOwner().GetLogin(), repo.GetName()), nil
}

// RepositoryFactory is factory to create new GitHub repository object.
type RepositoryFactory struct{}

// NewRepositoryFactory return new RepositoryFactory object.
func NewRepositoryFactory() *RepositoryFactory {
	return &RepositoryFactory{}
}

// new GitHub repository object.
func (f *RepositoryFactory) new(owner string, name string) Repository {
	return Repository{
		Owner: owner,
		Name:  name,
	}
}
