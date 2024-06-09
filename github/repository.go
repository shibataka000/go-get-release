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
		return Repository{}, &InvalidRepositoryFullNameFormatError{fullName}
	}
	return newRepository(s[0], s[1]), nil
}

// fullName returns a repository full name. It is 'OWNER/REPO' format.
func (r Repository) fullName() string {
	return fmt.Sprintf("%s/%s", r.owner, r.name)
}
