package github

import "fmt"

func getDockerComposeAsset(c *client, repo *repository, release *release, goos, goarch string) (Asset, error) {
	name := ""
	if goos == "darwin" {
		name = "docker-compose-Darwin-x86_64"
	} else {
		a, err := getGeneralAsset(c, repo, release, goos, goarch)
		if err != nil {
			return nil, err
		}
		name = a.Name()
	}

	downloadURL := fmt.Sprintf("https://github.com/docker/compose/releases/download/%s/%s", release.tag, name)

	binaryName := "docker-compose"
	if goos == "windows" {
		binaryName += ".exe"
	}
	return &asset{
		name:        name,
		downloadURL: downloadURL,
		binaryName:  binaryName,
	}, nil
}
