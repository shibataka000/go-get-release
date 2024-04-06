package github

import (
	"fmt"
	"regexp"
)

// SearchQuery.
type SearchQuery struct {
	Repository Repository
	Release    Release
}

// NewSearchQuery return new search query instance.
func NewSearchQuery(repo Repository, release Release) SearchQuery {
	return SearchQuery{
		Repository: repo,
		Release:    release,
	}
}

// ParseSearchQuery parse search query string and return search query instance.
func ParseSearchQuery(query string) (SearchQuery, error) {
	re := regexp.MustCompile(`(([^/=]+)/)?([^/=]+)(=([^/=]+))?`)
	submatch := re.FindStringSubmatch(query)
	if submatch == nil || len(submatch) != 6 {
		return SearchQuery{}, fmt.Errorf("%s is invalid query", query)
	}
	return NewSearchQuery(NewRepository(submatch[2], submatch[3]), NewRelease(submatch[5])), nil
}

// HasRepository return true if search query specify repository's owner and name.
func (q SearchQuery) HasRepository() bool {
	return q.Repository.Owner != "" && q.Repository.Name != ""
}

// HasRelease return true if search query specify release tag.
func (q SearchQuery) HasRelease() bool {
	return q.Release.Tag != ""
}

// SearchRepositoryQuery return query string to search repository in GitHub.
func (q SearchQuery) SearchRepositoryQuery() string {
	return q.Repository.Name
}
