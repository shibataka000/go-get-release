package github

// Release in GitHub repository
type Release interface {
	GetAsset(goos, goarch string) (Asset, error)
	Tag() string
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
