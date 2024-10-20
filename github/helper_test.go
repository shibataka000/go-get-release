package github

import "os"

// githubTokenForTest is authentication token for GitHub API requests. This can be used for test only.
var githubTokenForTest = os.Getenv("GITHUB_TOKEN")

// must is a helper that wraps a call to a function returning (E, error) and panics if the error is non-nil.
// This is intended for use in variable initializations.
func must[E any](e E, err error) E {
	if err != nil {
		panic(err)
	}
	return e
}
