package github

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/google/go-github/github"
)

// AssetLister list assets
type AssetLister interface {
	List() ([]Asset, error)
}

// newAssetLister return assetLister
func newAssetLister(r *release) AssetLister {
	switch {
	case r.repo.Owner() == "hashicorp" && r.repo.Name() == "terraform":
		return &terraformAssetLister{
			client:  r.client,
			repo:    r.repo,
			release: r,
		}
	case r.repo.Owner() == "helm" && r.repo.Name() == "helm":
		return &helmAssetLister{
			client:  r.client,
			repo:    r.repo,
			release: r,
		}
	case r.repo.Owner() == "gravitational" && r.repo.Name() == "teleport":
		return &teleportAssetLister{
			client:  r.client,
			repo:    r.repo,
			release: r,
		}
	default:
		return &githubAssetLister{
			client:  r.client,
			repo:    r.repo,
			release: r,
		}
	}
}

// githubAssetLister is asset lister for GitHub release
type githubAssetLister struct {
	client  *client
	repo    *repository
	release *release
}

// List assets in GitHub release
func (l *githubAssetLister) List() ([]Asset, error) {
	c := l.client
	assets, _, err := c.client.Repositories.ListReleaseAssets(c.ctx, l.repo.Owner(), l.repo.Name(), l.release.id, &github.ListOptions{
		PerPage: 100,
	})
	if err != nil {
		return nil, err
	}

	result := []Asset{}
	for _, a := range assets {
		result = append(result, &asset{
			client:      l.client,
			repo:        l.repo,
			release:     l.release,
			downloadURL: a.GetBrowserDownloadURL(),
		})
	}
	return result, nil
}

// terraformAssetLister is asset lister for hashicorp/terraform
type terraformAssetLister struct {
	client  *client
	repo    *repository
	release *release
}

// List assets in hashicorp/terraform
func (l *terraformAssetLister) List() ([]Asset, error) {
	baseURL, err := url.Parse("https://releases.hashicorp.com")
	if err != nil {
		return nil, err
	}
	baseURL.Path = path.Join(baseURL.Path, "terraform", l.release.Version())

	res, err := http.Get(baseURL.String())
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return nil, fmt.Errorf("status code error: %d %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
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

	results := []Asset{}
	for _, assetName := range assetNames {
		downloadURL, err := url.Parse(baseURL.String())
		if err != nil {
			return nil, err
		}
		downloadURL.Path = path.Join(downloadURL.Path, assetName)
		results = append(results, &asset{
			client:      l.client,
			repo:        l.repo,
			release:     l.release,
			downloadURL: downloadURL.String(),
		})
	}
	return results, nil
}

// helmAssetLister is asset lister for helm/helm
type helmAssetLister struct {
	client  *client
	repo    *repository
	release *release
}

// List assets in helm/helm
func (l *helmAssetLister) List() ([]Asset, error) {
	baseURL, err := url.Parse("https://get.helm.sh")
	if err != nil {
		return nil, err
	}

	gLister := &githubAssetLister{
		client:  l.client,
		repo:    l.repo,
		release: l.release,
	}
	assets, err := gLister.List()
	if err != nil {
		return nil, err
	}

	result := []Asset{}
	for _, a := range assets {
		assetName := strings.TrimRight(a.Name(), ".asc")
		if strings.Contains(assetName, "sha256") {
			continue
		}
		downloadURL, err := url.Parse(baseURL.String())
		if err != nil {
			return nil, err
		}
		downloadURL.Path = path.Join(downloadURL.Path, assetName)
		result = append(result, &asset{
			client:      l.client,
			repo:        l.repo,
			release:     l.release,
			downloadURL: downloadURL.String(),
		})
	}
	return result, nil
}

// teleportAssetLister is asset lister for gravitational/teleport
type teleportAssetLister struct {
	client  *client
	repo    *repository
	release *release
}

// TeleportReleasesResponse is response from https://dashboard.gravitational.com/webapi/releases-oss
type TeleportReleasesResponse struct {
	Next  int               `json:"next"`
	Last  int               `json:"last"`
	Items []TeleportRelease `json:"items"`
}

// TeleportRelease is part of response from https://dashboard.gravitational.com/webapi/releases-oss
type TeleportRelease struct {
	Version   string          `json:"version"`
	Downloads []TeleportAsset `json:"downloads"`
}

// TeleportAsset is part of response from https://dashboard.gravitational.com/webapi/releases-oss
type TeleportAsset struct {
	URL string `json:"url"`
}

// List assets in gravitational/teleport
func (l *teleportAssetLister) List() ([]Asset, error) {
	var teleportRelease TeleportRelease
	found := false

	page := 0
	for {
		req, err := http.NewRequest("GET", "https://dashboard.gravitational.com/webapi/releases-oss", nil)
		if err != nil {
			return nil, err
		}

		params := req.URL.Query()
		params.Add("product", "teleport")
		params.Add("page", fmt.Sprintf("%d", page))
		req.URL.RawQuery = params.Encode()

		client := &http.Client{}
		res, err := client.Do(req)
		if err != nil {
			return nil, err
		}
		defer res.Body.Close()

		if res.StatusCode != 200 {
			return nil, fmt.Errorf("status code error: %d %s", res.StatusCode, res.Status)
		}

		var response TeleportReleasesResponse
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(body, &response)
		if err != nil {
			return nil, err
		}

		for _, tr := range response.Items {
			if tr.Version == l.release.Tag() {
				teleportRelease = tr
				found = true
				break
			}
		}

		if found || response.Last == page {
			break
		}

		page = response.Next
	}

	if !found {
		return nil, fmt.Errorf("no teleport release found")
	}

	result := []Asset{}
	for _, a := range teleportRelease.Downloads {
		result = append(result, &asset{
			client:      l.client,
			repo:        l.repo,
			release:     l.release,
			downloadURL: a.URL,
		})
	}
	return result, nil
}
