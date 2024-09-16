package github

import (
	"os"
	"testing"
)

func TestFind(t *testing.T) {
	// todo: implement this.
}

func TestInstall(t *testing.T) {
	// todo: implement this.
}

// githubTokenForTest is authentication token for GitHub API requests. This can be used for test only.
var githubTokenForTest = os.Getenv("GITHUB_TOKEN")
