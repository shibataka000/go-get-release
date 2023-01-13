package github

import (
	"context"
	"net/http"

	"github.com/google/go-github/v48/github"
	"golang.org/x/oauth2"
)

// Repository for GitHub repository.
type Repository struct {
	client *github.Client
}

// New client for GitHub.
func New(ctx context.Context, token string) *Repository {
	var httpClient *http.Client
	if token != "" {
		tokenSource := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})
		httpClient = oauth2.NewClient(ctx, tokenSource)
	}
	client := github.NewClient(httpClient)
	return &Repository{
		client: client,
	}
}
