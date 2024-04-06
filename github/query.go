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

// ParseQuery parse search query string and return search query instance.
func ParseQuery(query string) (SearchQuery, error) {
	re := regexp.MustCompile(`(([^/=]+)/)?([^/=]+)(=([^/=]+))?`)
	submatch := re.FindStringSubmatch(query)
	if submatch == nil || len(submatch) != 6 {
		return SearchQuery{}, fmt.Errorf("%s is invalid query", query)
	}
	return NewSearchQuery(NewRepository(submatch[2], submatch[3]), NewRelease(submatch[5])), nil
}

// HasOwner return true if search query has repository owner.
func (q SearchQuery) HasOwner() bool {
	return q.Repository.Owner != ""
}

// HasTag return true if search query has tag.
func (q SearchQuery) HasTag() bool {
	return q.Release.Tag != ""
}
