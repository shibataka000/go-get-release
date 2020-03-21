package github

import (
	"fmt"
	"strings"
)

func getProtobufAsset(c *client, repo *repository, release *release, goos, goarch string) (Asset, error) {
	version := strings.TrimLeft(release.tag, "v")
	name := ""
	switch goos {
	case "linux":
		name = fmt.Sprintf("protoc-%s-linux-x86_64.zip", version)
	case "windows":
		name = fmt.Sprintf("protoc-%s-win64.zip", version)
	case "darwin":
		name = fmt.Sprintf("protoc-%s-osx-x86_64.zip", version)
	default:
		return nil, fmt.Errorf("Fail to guess asset name in istio. Unexpected GOOS value: %s", goos)
	}

	downloadURL := fmt.Sprintf("https://github.com/%s/%s/releases/download/%s/%s", repo.owner, repo.name, release.tag, name)

	binaryName := "protoc"
	if goos == "windows" {
		binaryName += ".exe"
	}

	return &asset{
		name:        name,
		downloadURL: downloadURL,
		binaryName:  binaryName,
	}, nil

}
