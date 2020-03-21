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

func (r *release) GetAsset(goos, goarch string) (Asset, error) {
	switch {
	case r.repo.owner == "docker" && r.repo.name == "compose":
		return getDockerComposeAsset(r.client, r.repo, r, goos, goarch)
	case r.repo.owner == "docker" && r.repo.name == "machine":
		return getDockerMachineAsset(r.client, r.repo, r, goos, goarch)
	case r.repo.owner == "helm" && r.repo.name == "helm":
		return getHelmAsset(r.client, r.repo, r, goos, goarch)
	case r.repo.owner == "istio" && r.repo.name == "istio":
		return getIstioAsset(r.client, r.repo, r, goos, goarch)
	case r.repo.owner == "hashicorp" && r.repo.name == "terraform":
		return getTerraformAsset(r.client, r.repo, r, goos, goarch)
	default:
		return getGeneralAsset(r.client, r.repo, r, goos, goarch)
	}
}

func (r *release) Tag() string {
	return r.tag
}
