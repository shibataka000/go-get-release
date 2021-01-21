package github

func (r *release) GetAsset(goos, goarch string) (Asset, error) {
	if isSpecialAsset(r.repo.owner, r.repo.name) {
		return r.getSpecialAsset(goos, goarch)
	}
	return r.getGeneralAsset(goos, goarch)
}
