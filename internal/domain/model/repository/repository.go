package repository

// Repository in GitHub
type Repository struct {
	owner string
	name  string
}

// New repository instance.
func New(owner string, name string) Repository {
	return Repository{
		owner: owner,
		name:  name,
	}
}

// Owner return repository owner.
func (r Repository) Owner() string {
	return r.owner
}

// Name return repository name.
func (r Repository) Name() string {
	return r.name
}
