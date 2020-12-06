package github

import (
	"fmt"
	"strings"

	"github.com/google/go-github/github"
)

func getGeneralAsset(c *client, repo *repository, release *release, goos, goarch string) (Asset, error) {
	assets, _, err := c.client.Repositories.ListReleaseAssets(c.ctx, repo.owner, repo.name, release.id, &github.ListOptions{})
	if err != nil {
		return nil, err
	}

	filteredAssets := []*github.ReleaseAsset{}
	for _, asset := range assets {
		if strings.HasSuffix(asset.GetName(), ".sha256") {
			continue
		}
		goosInAsset, err := getGoosByAsset(asset.GetName())
		if err != nil {
			continue
		}
		goarchInAsset, err := getGoarchByAsset(asset.GetName())
		if err != nil {
			continue
		}
		if goos == goosInAsset && goarch == goarchInAsset {
			filteredAssets = append(filteredAssets, asset)
		}
	}

	if len(filteredAssets) == 0 {
		return nil, fmt.Errorf("No asset found")
	} else if len(filteredAssets) >= 2 {
		assetNames := []string{}
		for _, asset := range filteredAssets {
			assetNames = append(assetNames, asset.GetName())
		}
		return nil, fmt.Errorf("Too many assets found: %v", strings.Join(assetNames, ", "))
	}

	downloadURL := filteredAssets[0].GetBrowserDownloadURL()

	binaryName := repo.name
	if goos == "windows" {
		binaryName += ".exe"
	}

	return &asset{
		downloadURL: downloadURL,
		binaryName:  binaryName,
	}, nil
}

func getGoosByAsset(asset string) (string, error) {
	s := strings.ToLower(asset)
	switch {
	case strings.Contains(s, "linux"):
		return "linux", nil
	case strings.Contains(s, "windows"):
		return "windows", nil
	case strings.Contains(s, "darwin"):
		return "darwin", nil
	case strings.Contains(s, "osx"):
		return "darwin", nil
	case strings.Contains(s, "win"):
		return "windows", nil
	default:
		return "", fmt.Errorf("Fail to guess GOOS by asset: %s", asset)
	}
}

func getGoarchByAsset(asset string) (string, error) {
	s := strings.ToLower(asset)
	switch {
	case strings.Contains(s, "amd64"):
		return "amd64", nil
	case strings.Contains(s, "x86_64"):
		return "amd64", nil
	default:
		return "", fmt.Errorf("Fail to guess GOARCH by asset: %s", asset)
	}
}
