package github

import (
	"bytes"
	"fmt"
	"html/template"
	"strings"

	"github.com/google/go-github/github"
)

// Release in GitHub repository
type Release interface {
	Tag() string
	Asset(string) (Asset, error)
	FindAssetByPlatform(goos string, goarch string) (Asset, error)
}

type release struct {
	client *client
	repo   *repository
	tag    string
	id     int64
}

func (r *release) Tag() string {
	return r.tag
}

func (r *release) Asset(name string) (Asset, error) {
	assets, err := r.assets()
	if err != nil {
		return nil, err
	}
	for _, a := range assets {
		if a.Name() == name {
			return a, nil
		}
	}
	return nil, fmt.Errorf("no asset found: %s", name)
}

func (r *release) FindAssetByPlatform(goos, goarch string) (Asset, error) {
	assetName, err := r.renderAssetName(goos, goarch)
	if err == nil {
		return r.Asset(assetName)
	}

	assets, err := r.listAssetsFilteredByPlatform(goos, goarch)
	if err != nil {
		return nil, err
	}

	if len(assets) == 0 {
		return nil, fmt.Errorf("no asset found")
	} else if len(assets) >= 2 {
		assetNames := []string{}
		for _, asset := range assets {
			assetNames = append(assetNames, asset.Name())
		}
		return nil, fmt.Errorf("too many assets found: %v", strings.Join(assetNames, ", "))
	}

	return assets[0], nil
}

func (r *release) assets() ([]Asset, error) {
	c := r.client
	repo := r.repo
	githubAssets, _, err := c.client.Repositories.ListReleaseAssets(c.ctx, repo.owner, repo.name, r.id, &github.ListOptions{
		PerPage: 100,
	})
	if err != nil {
		return nil, err
	}
	assets := []Asset{}
	for _, a := range githubAssets {
		assets = append(assets, &asset{
			name:        a.GetName(),
			downloadURL: a.GetBrowserDownloadURL(),
		})
	}
	return assets, nil
}

func (r *release) listAssetsFilteredByPlatform(goos, goarch string) ([]Asset, error) {
	var result []Asset
	assets, err := r.assets()
	if err != nil {
		return nil, err
	}
	for _, a := range assets {
		if !a.IsReleaseBinary() {
			continue
		}
		assetGoos, err := a.Goos()
		if err != nil {
			continue
		}
		assetGoarch, err := a.Goarch()
		if err != nil {
			continue
		}
		if assetGoos == goos && assetGoarch == goarch {
			result = append(result, a)
		}
	}
	return result, nil
}

func (r *release) renderAssetName(goos, goarch string) (string, error) {
	key := fmt.Sprintf("%s/%s", r.repo.Owner(), r.repo.Name())
	tmplMap, ok := assetNameMap[key]
	if !ok {
		return "", fmt.Errorf("asset name is not predefined: %s", key)
	}

	key = fmt.Sprintf("%s/%s", goos, goarch)
	tmplStr := ""
	if s, ok := tmplMap["default"]; ok {
		tmplStr = s
	}
	if s, ok := tmplMap[key]; ok {
		tmplStr = s
	}
	if tmplStr == "" {
		return "", fmt.Errorf("unsupported GOOS and GOARCH: %s", key)
	}

	version := strings.TrimLeft(r.tag, "v")

	buf := new(bytes.Buffer)
	tmpl, err := template.New("assetName").Parse(tmplStr)
	if err != nil {
		return "", err
	}
	err = tmpl.Execute(buf, struct {
		Tag     string
		Version string
		Goos    string
		Goarch  string
	}{
		Tag:     r.tag,
		Version: version,
		Goos:    goos,
		Goarch:  goarch,
	})
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}
