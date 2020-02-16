package pkg

import (
	"context"
	"fmt"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

type pkgInfo struct {
	owner       string
	repo        string
	tag         string
	asset       string
	downloadURL string
	binary      string
	isArchived  bool
}

type githubClient struct {
	client *github.Client
	ctx    context.Context
}

func getPkgInfo(pkgName string, option *Option) (*pkgInfo, error) {
	c, err := newGithubClient(option)
	if err != nil {
		return nil, err
	}

	owner, repo, tag, err := parsePkgName(pkgName)
	if err != nil {
		return nil, err
	}

	if owner == "" {
		owner, repo, err = c.searchRepository(repo)
		if err != nil {
			return nil, err
		}
	}

	if tag == "" {
		tag, err = c.getLatestTag(owner, repo)
		if err != nil {
			return nil, err
		}
	}

	asset, downloadURL, err := c.getAsset(owner, repo, tag, option.OS, option.Arch)
	if err != nil {
		return nil, err
	}

	return &pkgInfo{
		owner:       owner,
		repo:        repo,
		tag:         tag,
		asset:       asset,
		downloadURL: downloadURL,
		binary:      getBinaryName(owner, repo, option.OS),
		isArchived:  isArchived(asset),
	}, nil
}

func parsePkgName(pkgName string) (string, string, string, error) {
	re := regexp.MustCompile(`(([^/=]+)/)?([^/=]+)(=([^=]+))?`)
	if re.MatchString(pkgName) {
		match := re.FindStringSubmatch(pkgName)
		return match[2], match[3], match[5], nil
	}
	return "", "", "", fmt.Errorf("Parsing pkgName failed: %s\npkgName should be \"owner/repo=tag\" format", pkgName)
}

func newGithubClient(option *Option) (*githubClient, error) {
	ctx := context.Background()

	var client *github.Client
	if option.GithubToken == "" {
		client = github.NewClient(nil)
	} else {
		ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: option.GithubToken})
		tc := oauth2.NewClient(ctx, ts)
		client = github.NewClient(tc)
	}

	return &githubClient{
		client: client,
		ctx:    ctx,
	}, nil
}

func (c *githubClient) searchRepository(repo string) (string, string, error) {
	switch repo {
	default:
		result, _, err := c.client.Search.Repositories(c.ctx, repo, &github.SearchOptions{})
		if err != nil {
			return "", repo, err
		}
		if len(result.Repositories) == 0 {
			return "", repo, fmt.Errorf("No repository found: %s", repo)
		}
		return result.Repositories[0].GetOwner().GetLogin(), result.Repositories[0].GetName(), nil
	}
}

func (c *githubClient) getLatestTag(owner, repo string) (string, error) {
	release, _, err := c.client.Repositories.GetLatestRelease(c.ctx, owner, repo)
	if err != nil {
		return "", err
	}
	return release.GetTagName(), nil
}

func (c *githubClient) getAsset(owner, repo, tag, goos, goarch string) (string, string, error) {
	switch {
	case owner == "hashicorp" && repo == "terraform":
		version := strings.TrimLeft(tag, "v")
		asset := fmt.Sprintf("terraform_%s_%s_%s.zip", version, goos, goarch)
		downloadURL := fmt.Sprintf("https://releases.hashicorp.com/terraform/%s/%s", version, asset)
		return asset, downloadURL, nil

	case owner == "helm" && repo == "helm":
		asset := fmt.Sprintf("helm-%s-%s-%s.tar.gz", tag, goos, goarch)
		if goos == "windows" {
			asset = fmt.Sprintf("helm-%s-%s-%s.zip", tag, goos, goarch)
		}
		downloadURL := fmt.Sprintf("https://get.helm.sh/%s", asset)
		return asset, downloadURL, nil

	case owner == "docker" && repo == "compose" && goos == "darwin":
		asset := "docker-compose-Darwin-x86_64"
		downloadURL := fmt.Sprintf("https://github.com/docker/compose/releases/download/%s/%s", tag, asset)
		return asset, downloadURL, nil

	default:
		release, _, err := c.client.Repositories.GetReleaseByTag(c.ctx, owner, repo, tag)
		if err != nil {
			return "", "", err
		}
		assets, _, err := c.client.Repositories.ListReleaseAssets(c.ctx, owner, repo, release.GetID(), &github.ListOptions{})

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
			return "", "", fmt.Errorf("No asset found")
		} else if len(filteredAssets) >= 2 {
			assetNames := []string{}
			for _, asset := range filteredAssets {
				assetNames = append(assetNames, asset.GetName())
			}
			return "", "", fmt.Errorf("Too many assets found: %v", strings.Join(assetNames, ", "))
		}
		return filteredAssets[0].GetName(), filteredAssets[0].GetBrowserDownloadURL(), nil
	}
}

func getBinaryName(owner, repo, goos string) string {
	ext := ""
	if goos == "windows" {
		ext = ".exe"
	}
	switch {
	case owner == "docker" && repo == "machine":
		return "docker-machine" + ext
	case owner == "docker" && repo == "compose":
		return "docker-compose" + ext
	case owner == "istio" && repo == "istio":
		return "istioctl" + ext
	default:
		return repo + ext
	}
}

func isArchived(asset string) bool {
	archivedExt := []string{".tar", ".gz", ".tgz", ".bz2", ".tbz", ".Z", ".zip", ".bz2", ".lzh", ".7z", ".gz", ".rar", ".cab", ".afz"}
	for _, ext := range archivedExt {
		if filepath.Ext(asset) == ext {
			return true
		}
	}
	return false
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
	case strings.Contains(s, "win"):
		return "windows", nil
	case strings.Contains(s, "osx"):
		return "darwin", nil
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
	case strings.HasPrefix(asset, "istio-"):
		return "amd64", nil
	default:
		return "", fmt.Errorf("Fail to guess GOARCH by asset: %s", asset)
	}
}
