package pkg

import (
	"fmt"
	"regexp"
)

// Query to search package.
type Query struct {
	Repository GitHubRepository
	Tag        string
}

// NewQuery return new query instance to search package.
func NewQuery(repo GitHubRepository, tag string) Query {
	return Query{
		Repository: repo,
		Tag:        tag,
	}
}

// ParseQuery parse query string and return query instance.
func ParseQuery(query string) (Query, error) {
	re := regexp.MustCompile(`(([^/=]+)/)?([^/=]+)(=([^/=]+))?`)
	submatch := re.FindStringSubmatch(query)
	if submatch == nil || len(submatch) != 6 {
		return Query{}, fmt.Errorf("%s is invalid query", query)
	}
	return NewQuery(NewGitHubRepository(submatch[2], submatch[3]), submatch[5]), nil
}

// SearchRepositoryQuery return query string to search GitHub repository by SearchRepository function.
func (q Query) SearchRepositoryQuery() string {
	if q.Repository.Owner == "" {
		return q.Repository.Name
	}
	return q.Repository.FullName()
}

// HasFullRepositoryName return true if query string has both of owner and name of GitHub repository.
func (q Query) HasFullRepositoryName() bool {
	return q.Repository.Owner != "" && q.Repository.Name != ""
}

// HasTag return true if query string has GitHub release tag.
func (q Query) HasTag() bool {
	return q.Tag != ""
}
