package github

import (
	"fmt"
	"strings"
)

func getTerraformAsset(c *client, repo *repository, release *release, goos, goarch string) (Asset, error) {
	version := strings.TrimLeft(release.tag, "v")
	name := fmt.Sprintf("terraform_%s_%s_%s.zip", version, goos, goarch)

	downloadURL := fmt.Sprintf("https://releases.hashicorp.com/terraform/%s/%s", version, name)

	binaryName := "terraform"
	if goos == "windows" {
		binaryName += ".exe"
	}

	return &asset{
		name:        name,
		downloadURL: downloadURL,
		binaryName:  binaryName,
	}, nil
}
