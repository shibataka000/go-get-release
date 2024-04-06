package github

// Repository is repository in GitHub.
type Repository struct {
	Owner string
	Name  string
}

// NewRepository return new repository instance.
func NewRepository(owner string, name string) Repository {
	return Repository{
		Owner: owner,
		Name:  name,
	}
}

// Equal return true if two repositories are same.
func (r Repository) Equal(other Repository) bool {
	return r.Owner == other.Owner && r.Name == other.Name
}
