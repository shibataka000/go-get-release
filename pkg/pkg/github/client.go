package github

import (
	"context"
	"fmt"

	g "github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

// Client to fetch data from GitHub
type Client interface {
	GetRepository(owner, repo string) Repository
	FindRepository(keyword string) (Repository, error)
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

func (c *client) GetRepository(owner, repo string) Repository {
	return &repository{
		client: c,
		owner:  owner,
		name:   repo,
	}
}

func (c *client) FindRepository(keyword string) (Repository, error) {
	result, _, err := c.client.Search.Repositories(c.ctx, keyword, &g.SearchOptions{})
	if err != nil {
		return nil, err
	}
	if len(result.Repositories) == 0 {
		return nil, fmt.Errorf("No repository found: %s", keyword)
	}
	return c.GetRepository(result.Repositories[0].GetOwner().GetLogin(), result.Repositories[0].GetName()), nil
}
