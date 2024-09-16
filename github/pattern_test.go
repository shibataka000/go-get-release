package github

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPatternMatch(t *testing.T) {
	tests := []struct {
		name    string
		pattern Pattern
		asset   Asset
		match   bool
	}{
		{
			name:    "gh_2.52.0_linux_amd64.tar.gz",
			pattern: mustNewPatternFromString("gh_.*_linux_amd64.tar.gz", "gh"),
			asset:   newAsset(0, "gh_2.52.0_linux_amd64.tar.gz"),
			match:   true,
		},
		{
			name:    "gh_2.52.0_linux_arm64.tar.gz",
			pattern: mustNewPatternFromString("gh_.*_linux_amd64.tar.gz", "gh"),
			asset:   newAsset(0, "gh_2.52.0_linux_arm64.tar.gz"),
			match:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require := require.New(t)
			require.Equal(tt.match, tt.pattern.match(tt.asset))
		})
	}
}

func TestPatternApply(t *testing.T) {
	// todo: implement this.
}

// mustNewPatternFromString is like [newPatternFromString] but panics if arguments cannot be parsed.
func mustNewPatternFromString(asset string, execBinary string) Pattern {
	p, err := newPatternFromString(asset, execBinary)
	if err != nil {
		panic(err)
	}
	return p
}
