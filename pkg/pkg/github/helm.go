package github

import (
	"fmt"
)

func getHelmAsset(c *client, repo *repository, release *release, goos, goarch string) (Asset, error) {
	name := ""
	if goos == "windows" {
		name = fmt.Sprintf("helm-%s-%s-%s.zip", release.tag, goos, goarch)
	} else {
		name = fmt.Sprintf("helm-%s-%s-%s.tar.gz", release.tag, goos, goarch)
	}

	downloadURL := fmt.Sprintf("https://get.helm.sh/%s", name)

	binaryName := "helm"
	if goos == "windows" {
		binaryName += ".exe"
	}

	return &asset{
		name:        name,
		downloadURL: downloadURL,
		binaryName:  binaryName,
	}, nil
}
