package github

import (
	"fmt"
)

func getIstioAsset(c *client, repo *repository, release *release, goos, goarch string) (Asset, error) {
	name := ""
	switch goos {
	case "linux":
		name = fmt.Sprintf("istio-%s-linux.tar.gz", release.tag)
	case "windows":
		name = fmt.Sprintf("istio-%s-win.zip", release.tag)
	case "darwin":
		name = fmt.Sprintf("istio-%s-osx.tar.gz", release.tag)
	default:
		return nil, fmt.Errorf("Fail to guess asset name in istio. Unexpected GOOS value: %s", goos)
	}

	downloadURL := fmt.Sprintf("https://github.com/%s/%s/releases/download/%s/%s", repo.owner, repo.name, release.tag, name)

	binaryName := "istioctl"
	if goos == "windows" {
		binaryName += ".exe"
	}

	return &asset{
		name:        name,
		downloadURL: downloadURL,
		binaryName:  binaryName,
	}, nil
}
