package pkg

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParseQuery(t *testing.T) {
	tests := []struct {
		name     string
		queryStr string
		query    Query
	}{
		{
			name:     "shibataka000/go-get-release=v0.0.1",
			queryStr: "shibataka000/go-get-release=v0.0.1",
			query:    NewQuery(NewGitHubRepository("shibataka000", "go-get-release"), "v0.0.1"),
		},
		{
			name:     "go-get-release",
			queryStr: "go-get-release",
			query:    NewQuery(NewGitHubRepository("", "go-get-release"), ""),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := require.New(t)
			query, err := ParseQuery(tt.queryStr)
			assert.NoError(err)
			assert.Equal(tt.query, query)
		})
	}
}

func TestQuerySearchRepositoryQuery(t *testing.T) {
	tests := []struct {
		name                  string
		query                 Query
		searchRepositoryQuery string
	}{
		{
			name:                  "shibataka000/go-get-release=v0.0.1",
			query:                 NewQuery(NewGitHubRepository("shibataka000", "go-get-release"), "v0.0.1"),
			searchRepositoryQuery: "shibataka000/go-get-release",
		},
		{
			name:                  "go-get-release",
			query:                 NewQuery(NewGitHubRepository("", "go-get-release"), ""),
			searchRepositoryQuery: "go-get-release",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := require.New(t)
			assert.Equal(tt.searchRepositoryQuery, tt.query.SearchRepositoryQuery())
		})
	}
}

func TestQueryHasFullRepositoryName(t *testing.T) {
	tests := []struct {
		name                  string
		query                 Query
		hasFullRepositoryName bool
	}{
		{
			name:                  "shibataka000/go-get-release=v0.0.1",
			query:                 NewQuery(NewGitHubRepository("shibataka000", "go-get-release"), "v0.0.1"),
			hasFullRepositoryName: true,
		},
		{
			name:                  "go-get-release",
			query:                 NewQuery(NewGitHubRepository("", "go-get-release"), ""),
			hasFullRepositoryName: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := require.New(t)
			assert.Equal(tt.hasFullRepositoryName, tt.query.HasFullRepositoryName())
		})
	}
}

func TestQueryHasTag(t *testing.T) {
	tests := []struct {
		name   string
		query  Query
		hasTag bool
	}{
		{
			name:   "shibataka000/go-get-release=v0.0.1",
			query:  NewQuery(NewGitHubRepository("shibataka000", "go-get-release"), "v0.0.1"),
			hasTag: true,
		},
		{
			name:   "go-get-release",
			query:  NewQuery(NewGitHubRepository("", "go-get-release"), ""),
			hasTag: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := require.New(t)
			assert.Equal(tt.hasTag, tt.query.HasFullRepositoryName())
		})
	}
}
