package github

import "os"

// githubTokenForTest is authentication token for GitHub API requests. This can be used for test only.
var githubTokenForTest = os.Getenv("GITHUB_TOKEN")

func must[E any](e E, err error) E {
	if err != nil {
		panic(err)
	}
	return e
}

// mustNewPatternFromString is like [newPatternFromString] but panics if arguments cannot be parsed.
func mustNewPatternFromString(asset string, execBinary string) Pattern {
	return must(newPatternFromString(asset, execBinary))
}
