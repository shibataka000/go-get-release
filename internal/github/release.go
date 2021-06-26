package github

import "github.com/google/go-github/github"

// Release in GitHub repository
type Release interface {
	Tag() string
	Assets() ([]Asset, error)
}

type release struct {
	client *client
	repo   *repository
	tag    string
	id     int64
}

func (r *release) Tag() string {
	return r.tag
}

func (r *release) Assets() ([]Asset, error) {
	c := r.client
	repo := r.repo
	githubAssets, _, err := c.client.Repositories.ListReleaseAssets(c.ctx, repo.owner, repo.name, r.id, &github.ListOptions{
		PerPage: 100,
	})
	if err != nil {
		return nil, err
	}
	assets := []Asset{}
	for _, a := range githubAssets {
		assets = append(assets, &asset{
			name:        a.GetName(),
			downloadURL: a.GetBrowserDownloadURL(),
		})
	}
	return assets, nil
}
