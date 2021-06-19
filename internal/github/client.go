package github

import (
	"context"
	"fmt"

	g "github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

// Client to fetch data from GitHub
type Client interface {
	Repository(owner, repo string) (Repository, error)
	FindRepository(keyword string) (Repository, error)
	SearchRepositories(keyword string) ([]Repository, error)
}

type client struct {
	client *g.Client
	ctx    context.Context
}

// NewClient return new GitHub client
func NewClient(token string) (Client, error) {
	ctx := context.Background()

	var c *g.Client
	if token == "" {
		c = g.NewClient(nil)
	} else {
		ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})
		tc := oauth2.NewClient(ctx, ts)
		c = g.NewClient(tc)
	}

	return &client{
		client: c,
		ctx:    ctx,
	}, nil

}

func (c *client) Repository(owner, repo string) (Repository, error) {
	result, _, err := c.client.Repositories.Get(c.ctx, owner, repo)
	if err != nil {
		return nil, err
	}
	return &repository{
		client: c,
		owner:  result.GetOwner().GetLogin(),
		name:   result.GetName(),
	}, nil
}

func (c *client) FindRepository(keyword string) (Repository, error) {
	repos, err := c.SearchRepositories(keyword)
	if err != nil {
		return nil, err
	}
	if len(repos) == 0 {
		return nil, fmt.Errorf("no repository found: %s", keyword)
	}
	return repos[0], nil
}

func (c *client) SearchRepositories(keyword string) ([]Repository, error) {
	result, _, err := c.client.Search.Repositories(c.ctx, keyword, &g.SearchOptions{})
	if err != nil {
		return nil, err
	}
	repos := []Repository{}
	for _, repo := range result.Repositories {
		repos = append(repos, &repository{
			client: c,
			owner:  repo.GetOwner().GetLogin(),
			name:   repo.GetName(),
		})
	}
	return repos, nil
}
