package github

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParseSearchQuery(t *testing.T) {
	tests := []struct {
		name     string
		queryStr string
		query    SearchQuery
	}{
		{
			name:     "shibataka000/go-get-release=v0.0.1",
			queryStr: "shibataka000/go-get-release=v0.0.1",
			query:    NewSearchQuery(NewRepository("shibataka000", "go-get-release"), NewRelease("v0.0.1")),
		},
		{
			name:     "go-get-release",
			queryStr: "go-get-release",
			query:    NewSearchQuery(NewRepository("", "go-get-release"), NewRelease("")),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require := require.New(t)
			query, err := ParseSearchQuery(tt.queryStr)
			require.NoError(err)
			require.Equal(tt.query, query)
		})
	}
}

func TestSearchQueryHasOwner(t *testing.T) {
	tests := []struct {
		name     string
		query    SearchQuery
		hasOwner bool
	}{
		{
			name:     "shibataka000/go-get-release=v0.0.1",
			query:    NewSearchQuery(NewRepository("shibataka000", "go-get-release"), NewRelease("v0.0.1")),
			hasOwner: true,
		},
		{
			name:     "go-get-release",
			query:    NewSearchQuery(NewRepository("", "go-get-release"), NewRelease("")),
			hasOwner: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require := require.New(t)
			require.Equal(tt.hasOwner, tt.query.HasRepository())
		})
	}
}

func TestSearchQueryHasTag(t *testing.T) {
	tests := []struct {
		name   string
		query  SearchQuery
		hasTag bool
	}{
		{
			name:   "shibataka000/go-get-release=v0.0.1",
			query:  NewSearchQuery(NewRepository("shibataka000", "go-get-release"), NewRelease("v0.0.1")),
			hasTag: true,
		},
		{
			name:   "go-get-release",
			query:  NewSearchQuery(NewRepository("", "go-get-release"), NewRelease("")),
			hasTag: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require := require.New(t)
			require.Equal(tt.hasTag, tt.query.HasRelease())
		})
	}
}
