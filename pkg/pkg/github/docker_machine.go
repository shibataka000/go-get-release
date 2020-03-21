package github

func getDockerMachineAsset(c *client, repo *repository, release *release, goos, goarch string) (Asset, error) {
	a, err := getGeneralAsset(c, repo, release, goos, goarch)
	if err != nil {
		return nil, err
	}

	binaryName := "docker-machine"
	if goos == "windows" {
		binaryName += ".exe"
	}

	return &asset{
		name:        a.Name(),
		downloadURL: a.DownloadURL(),
		binaryName:  binaryName,
	}, nil
}
