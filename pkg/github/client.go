package github

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"path/filepath"
	"sort"
	"strings"
	"text/template"

	"github.com/google/go-github/v48/github"
	"github.com/shibataka000/go-get-release/pkg/archive"
	"golang.org/x/oauth2"
)

// Client to fetch data from GitHub.
type Client interface {
	FindAsset(ctx context.Context, keyword, tag, goos, goarch string) (Asset, error)
}

type client struct {
	client *github.Client
}

// NewClient return new GitHub client.
func NewClient(ctx context.Context, token string) (Client, error) {
	var httpClient *http.Client
	if token != "" {
		tokenSource := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})
		httpClient = oauth2.NewClient(ctx, tokenSource)
	}
	c := github.NewClient(httpClient)

	return &client{
		client: c,
	}, nil
}

func (c *client) FindAsset(ctx context.Context, keyword, tag, goos, goarch string) (Asset, error) {
	var empty Asset
	var err error

	// Find repository.
	var repository *github.Repository
	if strings.Contains(keyword, "/") {
		splitted := strings.Split(keyword, "/")
		repository, _, err = c.client.Repositories.Get(ctx, splitted[0], splitted[1])
		if err != nil {
			return empty, err
		}
	} else {
		result, _, err := c.client.Search.Repositories(ctx, keyword, &github.SearchOptions{})
		if err != nil {
			return empty, err
		}
		if len(result.Repositories) == 0 {
			return empty, fmt.Errorf("no repository found")
		}
		repository = result.Repositories[0]
	}
	owner := repository.GetOwner().GetLogin()
	repo := repository.GetName()

	// Find release.
	var release *github.RepositoryRelease
	if tag == "" || tag == "latest" {
		release, _, err = c.client.Repositories.GetLatestRelease(ctx, owner, repo)
	} else {
		release, _, err = c.client.Repositories.GetReleaseByTag(ctx, owner, repo, tag)
	}
	if err != nil {
		return empty, err
	}

	// Return asset if registered.
	for _, asset := range registeredAsset {
		if asset.Owner == owner && asset.Repo == repo && asset.Goos == goos && asset.Goarch == goarch {
			buf := new(bytes.Buffer)
			tmpl, err := template.New("downloadURL").Parse(asset.DownloadURLTemplate)
			if err != nil {
				return empty, err
			}
			err = tmpl.Execute(buf, struct {
				Tag     string
				Version string
			}{
				Tag:     release.GetTagName(),
				Version: strings.TrimLeft(release.GetTagName(), "v"),
			})
			if err != nil {
				return empty, err
			}
			return Asset{
				Owner:       owner,
				Repo:        repo,
				Tag:         release.GetTagName(),
				DownloadURL: buf.String(),
				BinaryName:  asset.BinaryName,
				Goos:        goos,
				Goarch:      goarch,
			}, nil
		}
	}

	// Find assets.
	assets, _, err := c.client.Repositories.ListReleaseAssets(ctx, owner, repo, release.GetID(), &github.ListOptions{PerPage: 100})
	if err != nil {
		return empty, err
	}
	filtered := []*github.ReleaseAsset{}
	for _, asset := range assets {
		assetName := asset.GetName()
		if !(archive.IsArchived(assetName) || archive.IsCompressed(assetName) || filepath.Ext(assetName) == ".exe" || filepath.Ext(assetName) == "") {
			continue
		}
		assetGoos, assetGoarch, err := findPlatform(assetName)
		if err != nil {
			continue
		}
		if goos == assetGoos && goarch == assetGoarch {
			filtered = append(filtered, asset)
		}
	}
	if len(filtered) == 0 {
		return empty, fmt.Errorf("no asset found")
	} else if len(filtered) >= 2 {
		assetNames := []string{}
		for _, asset := range filtered {
			assetNames = append(assetNames, asset.GetName())
		}
		return empty, fmt.Errorf("too many assets matched, %v", assetNames)
	}
	asset := filtered[0]

	// Binary name.
	binaryName := repo
	if goos == "windows" {
		binaryName += ".exe"
	}

	return Asset{
		Owner:       owner,
		Repo:        repo,
		Tag:         release.GetTagName(),
		DownloadURL: asset.GetBrowserDownloadURL(),
		BinaryName:  binaryName,
		Goos:        goos,
		Goarch:      goarch,
	}, nil
}

// findPlatform find golang platform (GOOS/GOARCH) from platform map based on asset name.
func findPlatform(name string) (string, string, error) {
	goos, err := findPlatformHelper(name, goosMap)
	if err != nil {
		return "", "", fmt.Errorf("fail to guess GOOS from asset name: %s", name)
	}

	goarch, err := findPlatformHelper(name, goarchMap)
	if err != nil {
		goarch = "amd64"
	}

	return goos, goarch, nil
}

func findPlatformHelper(name string, platform map[string][]string) (string, error) {
	reversed := map[string]string{}
	for key, values := range platform {
		for _, value := range values {
			reversed[value] = key
		}
	}
	keys := []string{}
	for key := range reversed {
		keys = append(keys, key)
	}
	sort.Slice(keys, func(i, j int) bool { return len(keys[i]) > len(keys[j]) })

	lowerName := strings.ToLower(name)
	for _, key := range keys {
		if strings.Contains(lowerName, key) {
			if v, ok := reversed[key]; ok {
				return v, nil
			}
			return "", fmt.Errorf("index error: %s is not in %v", key, reversed)
		}
	}
	return "", fmt.Errorf("fail to guess go platform from name: %s", name)
}
