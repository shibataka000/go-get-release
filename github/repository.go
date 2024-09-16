package github

import (
	"fmt"
	"strings"
)

// Repository represents a GitHub repository.
type Repository struct {
	owner string
	name  string
}

// newRepository returns a new [Repository] object.
func newRepository(owner string, name string) Repository {
	return Repository{
		owner: owner,
		name:  name,
	}
}

// newRepositoryFromFullName returns a new [Repository] object from repository full name.
// Repository full name should be 'OWNER/REPO' format.
func newRepositoryFromFullName(fullName string) (Repository, error) {
	s := strings.Split(fullName, "/")
	if len(s) != 2 {
		return Repository{}, fmt.Errorf("%w: %s", ErrInvalidRepositoryFullName, fullName)
	}
	return newRepository(s[0], s[1]), nil
}
