package github

import "strings"

// Repository represents a GitHub repository.
type Repository struct {
	Owner string `yaml:"owner"`
	Name  string `yaml:"name"`
}

// newRepository returns a new GitHub repository object.
func newRepository(owner string, name string) Repository {
	return Repository{
		Owner: owner,
		Name:  name,
	}
}

// newRepositoryFromFullName returns a new GitHub repository object from repository full name.
// Repository full name must be 'OWNER/REPO' format.
func newRepositoryFromFullName(fullName string) (Repository, error) {
	s := strings.Split(fullName, "/")
	if len(s) != 2 {
		return Repository{}, nil
	}
	return newRepository(s[0], s[1]), nil
}
