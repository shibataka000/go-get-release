package github

// Repository in GitHub
type Repository interface {
	Owner() string
	Name() string
	GetLatestRelease() (Release, error)
	GetRelease(tag string) (Release, error)
}

type repository struct {
	client *client
	owner  string
	name   string
}

func (r *repository) Owner() string {
	return r.owner
}

func (r *repository) Name() string {
	return r.name
}

func (r *repository) GetLatestRelease() (Release, error) {
	c := r.client
	rel, _, err := c.client.Repositories.GetLatestRelease(c.ctx, r.owner, r.name)
	if err != nil {
		return nil, err
	}
	return &release{
		client: r.client,
		repo:   r,
		id:     rel.GetID(),
		tag:    rel.GetTagName(),
	}, nil
}

func (r *repository) GetRelease(tag string) (Release, error) {
	c := r.client
	rel, _, err := c.client.Repositories.GetReleaseByTag(c.ctx, r.owner, r.name, tag)
	if err != nil {
		return nil, err
	}
	return &release{
		client: r.client,
		repo:   r,
		id:     rel.GetID(),
		tag:    tag,
	}, nil
}
