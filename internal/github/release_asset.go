package github

import (
	"bytes"
	"fmt"
	"strings"
	"text/template"

	"github.com/google/go-github/github"
)

func (r *release) Asset(goos, goarch string) (Asset, error) {
	if isSpecialAsset(r.repo.owner, r.repo.name) {
		return r.getSpecialAsset(goos, goarch)
	}
	return r.getGeneralAsset(goos, goarch)
}

func (r *release) getSpecialAsset(goos, goarch string) (Asset, error) {
	repo := r.repo

	key := fmt.Sprintf("%s/%s", repo.owner, repo.name)
	assetMap, ok := specialAssetMap[key]
	if !ok {
		return nil, fmt.Errorf("%s is not found in specialAssetMap", key)
	}

	var assetTemplate *asset
	key = fmt.Sprintf("%s/%s", goos, goarch)
	if value, ok := assetMap["default"]; ok {
		assetTemplate = value
	}
	if value, ok := assetMap[key]; ok {
		assetTemplate = value
	}
	if assetTemplate == nil {
		return nil, fmt.Errorf("unsupported GOOS and GOARCH in this repository: %s", key)
	}

	version := strings.TrimLeft(r.tag, "v")

	tmpl, err := template.New("downloadURL").Parse(assetTemplate.downloadURL)
	if err != nil {
		return nil, err
	}
	buf := new(bytes.Buffer)
	err = tmpl.Execute(buf, struct {
		Owner   string
		Repo    string
		Tag     string
		Version string
		Goos    string
		Goarch  string
	}{
		Owner:   repo.owner,
		Repo:    repo.name,
		Tag:     r.tag,
		Version: version,
		Goos:    goos,
		Goarch:  goarch,
	})
	if err != nil {
		return nil, err
	}
	downloadURL := buf.String()

	binaryName := assetTemplate.binaryName

	return &asset{
		downloadURL: downloadURL,
		binaryName:  binaryName,
	}, nil
}

func (r *release) getGeneralAsset(goos, goarch string) (Asset, error) {
	c := r.client
	repo := r.repo

	assets, _, err := c.client.Repositories.ListReleaseAssets(c.ctx, repo.owner, repo.name, r.id, &github.ListOptions{})
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
		return nil, fmt.Errorf("no asset found")
	} else if len(filteredAssets) >= 2 {
		assetNames := []string{}
		for _, asset := range filteredAssets {
			assetNames = append(assetNames, asset.GetName())
		}
		return nil, fmt.Errorf("too many assets found: %v", strings.Join(assetNames, ", "))
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

func isSpecialAsset(owner, repo string) bool {
	key := fmt.Sprintf("%s/%s", owner, repo)
	_, ok := specialAssetMap[key]
	return ok
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
		return "", fmt.Errorf("fail to guess GOOS by asset: %s", asset)
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
		return "", fmt.Errorf("fail to guess GOARCH by asset: %s", asset)
	}
}
