package github

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/google/go-github/github"
)

// Release in GitHub repository
type Release interface {
	Tag() string
	Assets() ([]Asset, error)
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

func (r *release) Assets() ([]Asset, error) {
	if r.repo.Owner() == "hashicorp" && r.repo.Name() == "terraform" {
		return r.terraformAssets()
	} else if r.repo.Owner() == "helm" && r.repo.Name() == "helm" {
		return r.helmAssets()
	}

	c := r.client
	assets, _, err := c.client.Repositories.ListReleaseAssets(c.ctx, r.repo.owner, r.repo.name, r.id, &github.ListOptions{
		PerPage: 100,
	})
	if err != nil {
		return nil, err
	}

	result := []Asset{}
	for _, a := range assets {
		result = append(result, &asset{
			client:      r.client,
			repo:        r.repo,
			release:     r,
			downloadURL: a.GetBrowserDownloadURL(),
		})
	}
	return result, nil
}

func (r *release) terraformAssets() ([]Asset, error) {
	version := strings.TrimLeft(r.tag, "v")
	url := fmt.Sprintf("https://releases.hashicorp.com/terraform/%s", version)

	doc, err := newGoqueryDocument(url)
	if err != nil {
		return nil, err
	}

	assetNames := []string{}
	doc.Find("li").Each(func(i int, s *goquery.Selection) {
		assetName := s.Find("a").Text()
		if assetName != "../" {
			assetNames = append(assetNames, assetName)
		}
	})

	assets := []Asset{}
	for _, assetName := range assetNames {
		downloadURL := fmt.Sprintf("%s/%s", url, assetName) // todo: fix me
		assets = append(assets, &asset{
			client:      r.client,
			repo:        r.repo,
			release:     r,
			downloadURL: downloadURL,
		})
	}
	return assets, nil
}

func newGoqueryDocument(url string) (*goquery.Document, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		return nil, fmt.Errorf("status code error: %d %s", res.StatusCode, res.Status)
	}
	return goquery.NewDocumentFromReader(res.Body)
}

func (r *release) helmAssets() ([]Asset, error) {
	return nil, nil
}

func (r *release) Asset(name string) (Asset, error) {
	assets, err := r.Assets()
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
		for _, a := range assets {
			assetNames = append(assetNames, a.Name())
		}
		return nil, fmt.Errorf("too many assets found: %v", strings.Join(assetNames, ", "))
	}

	return assets[0], nil
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

func (r *release) listAssetsFilteredByPlatform(goos, goarch string) ([]Asset, error) {
	var result []Asset
	assets, err := r.Assets()
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
