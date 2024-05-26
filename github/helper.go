package github

import (
	"context"
	"net/http"

	"github.com/google/go-github/v48/github"
	"golang.org/x/oauth2"
)

// newGitHubClient returns a new client for GitHub API.
func newGitHubClient(ctx context.Context, token string) *github.Client {
	var httpClient *http.Client
	if token != "" {
		tokenSource := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})
		httpClient = oauth2.NewClient(ctx, tokenSource)
	}
	return github.NewClient(httpClient)
}
