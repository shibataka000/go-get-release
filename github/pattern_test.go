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
			name:    "Match",
			pattern: mustNewPatternFromString("gh_.*_linux_amd64.tar.gz", "gh"),
			asset:   newAsset(0, "gh_2.52.0_linux_amd64.tar.gz"),
			match:   true,
		},
		{
			name:    "NotMatch",
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

func TestPatternPriority(t *testing.T) {
	tests := []struct {
		name     string
		pattern  Pattern
		priority int
	}{
		{
			name:     "gh_2.52.0_linux_amd64.tar.gz",
			pattern:  mustNewPatternFromString("gh_2.52.0_linux_amd64.tar.gz", "gh"),
			priority: len("gh_2"),
		},
		{
			name:     "gh_.*_linux_amd64.tar.gz",
			pattern:  mustNewPatternFromString("gh_.*_linux_amd64.tar.gz", "gh"),
			priority: len("gh_"),
		},
		{
			name:     ".*_linux_amd64.tar.gz",
			pattern:  mustNewPatternFromString(".*_linux_amd64.tar.gz", "gh"),
			priority: len(""),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require := require.New(t)
			priority := tt.pattern.priority()
			require.Equal(tt.priority, priority)
		})
	}
}

func TestPatternExecute(t *testing.T) {
	tests := []struct {
		name       string
		pattern    Pattern
		asset      Asset
		execBinary ExecBinary
	}{
		{
			name:       "NamedCapturingGroup",
			pattern:    mustNewPatternFromString(`(?P<name>\w+)_[\d\.]+_linux_amd64.tar.gz`, "{{.name}}"),
			asset:      newAsset(0, "gh_2.52.0_linux_amd64.tar.gz"),
			execBinary: newExecBinary("gh"),
		},
		{
			name:       "CapturingGroup",
			pattern:    mustNewPatternFromString(`(\w+)_[\d\.]+_linux_amd64.tar.gz`, `{{index . "1"}}`),
			asset:      newAsset(0, "gh_2.52.0_linux_amd64.tar.gz"),
			execBinary: newExecBinary("gh"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require := require.New(t)
			execBinary, err := tt.pattern.execute(tt.asset)
			require.NoError(err)
			require.Equal(tt.execBinary, execBinary)
		})
	}
}

func TestFindAssetAndPattern(t *testing.T) {
	tests := []struct {
		name     string
		assets   AssetList
		patterns PatternList
		asset    Asset
		pattern  Pattern
		err      error
	}{
		{
			name:    "gh_2.52.0_linux_amd64.tar.gz",
			pattern: mustNewPatternFromString(`(?P<name>\w+)_[\d\.]+_linux_amd64.tar.gz`, "{{.name}}"),
			asset:   newAsset(0, "gh_2.52.0_linux_amd64.tar.gz"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require := require.New(t)
			asset, pattern, err := find(tt.assets, tt.patterns)
			if tt.err == nil {
				require.NoError(err)
				require.Equal(tt.asset, asset)
				require.Equal(tt.pattern, pattern)
			} else {
				require.ErrorIs(tt.err, err)
			}
		})
	}
}
