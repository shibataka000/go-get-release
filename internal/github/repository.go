package github

import "github.com/google/go-github/github"

// Repository in GitHub
type Repository interface {
	Owner() string
	Name() string
	Description() string
	LatestRelease() (Release, error)
	Release(tag string) (Release, error)
	ListRelease(n int) ([]Release, error)
}

type repository struct {
	client      *client
	owner       string
	name        string
	description string
}

func (r *repository) Owner() string {
	return r.owner
}

func (r *repository) Name() string {
	return r.name
}

func (r *repository) Description() string {
	return r.description
}

func (r *repository) LatestRelease() (Release, error) {
	c := r.client
	rel, _, err := c.client.Repositories.GetLatestRelease(c.ctx, r.owner, r.name)
	if err != nil {
		return nil, err
	}

	return &release{
		client:      r.client,
		repo:        r,
		id:          rel.GetID(),
		tag:         rel.GetTagName(),
		publishedAt: rel.PublishedAt.Time,
	}, nil
}

func (r *repository) Release(tag string) (Release, error) {
	c := r.client
	rel, _, err := c.client.Repositories.GetReleaseByTag(c.ctx, r.owner, r.name, tag)
	if err != nil {
		return nil, err
	}
	return &release{
		client:      r.client,
		repo:        r,
		id:          rel.GetID(),
		tag:         tag,
		publishedAt: rel.PublishedAt.Time,
	}, nil
}

func (r *repository) ListRelease(n int) ([]Release, error) {
	c := r.client
	result, _, err := c.client.Repositories.ListReleases(c.ctx, r.owner, r.name, &github.ListOptions{
		PerPage: 100,
	})
	if err != nil {
		return nil, err
	}
	releases := []Release{}
	for _, rel := range result {
		if rel.GetPrerelease() {
			continue
		}
		releases = append(releases, &release{
			client:      r.client,
			repo:        r,
			id:          rel.GetID(),
			tag:         rel.GetTagName(),
			publishedAt: rel.PublishedAt.Time,
		})
		if len(releases) >= n {
			break
		}
	}
	return releases, nil
}
