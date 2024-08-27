package github

import (
	"strings"
)

// Repository represents a GitHub repository.
type Repository struct {
	owner string
	name  string
}

// newRepository returns a new GitHub repository object.
func newRepository(owner string, name string) Repository {
	return Repository{
		owner: owner,
		name:  name,
	}
}

// newRepositoryFromFullName returns a new GitHub repository object from repository full name.
// Repository full name must be 'OWNER/REPO' format.
func newRepositoryFromFullName(fullName string) (Repository, error) {
	s := strings.Split(fullName, "/")
	if len(s) != 2 {
		return Repository{}, &InvalidRepositoryFullNameError{fullName}
	}
	return newRepository(s[0], s[1]), nil
}
