package pkg

import (
	"bytes"
	"fmt"
	"path"
	"path/filepath"
	"regexp"
	"strings"
	"text/template"

	"github.com/shibataka000/go-get-release/internal/github"
)

// FindInput is input to Find function
type FindInput struct {
	Name        string
	GithubToken string
	Goos        string
	Goarch      string
}

// FindOutput is output by Find function. This have information about golang release binary
type FindOutput struct {
	Owner       string
	Repo        string
	Tag         string
	Asset       string
	DownloadURL string
	BinaryName  string
	IsArchived  bool
}

// Find golang release binary and return its information
func Find(input *FindInput) (*FindOutput, error) {
	owner, repoName, tag, err := parse(input.Name)
	if err != nil {
		return nil, err
	}

	client, err := github.NewClient(input.GithubToken)
	if err != nil {
		return nil, err
	}

	var repo github.Repository
	if owner != "" {
		repo, err = client.Repository(owner, repoName)
	} else {
		repo, err = client.FindRepository(repoName)
	}
	if err != nil {
		return nil, err
	}

	var release github.Release
	if tag != "" {
		release, err = repo.Release(tag)
	} else {
		release, err = repo.LatestRelease()
	}
	if err != nil {
		return nil, err
	}

	downloadURL := ""
	if repo.Owner() == "hashicorp" && repo.Name() == "terraform" {
		version := strings.TrimLeft(release.Tag(), "v")
		downloadURL = fmt.Sprintf("https://releases.hashicorp.com/terraform/%s/terraform_%s_%s_%s.zip", version, version, input.Goos, input.Goarch)
	} else if repo.Owner() == "helm" && repo.Name() == "helm" && input.Goos == "windows" {
		downloadURL = fmt.Sprintf("https://get.helm.sh/helm-%s-%s-%s.zip", release.Tag(), input.Goos, input.Goarch)
	} else if repo.Owner() == "helm" && repo.Name() == "helm" {
		downloadURL = fmt.Sprintf("https://get.helm.sh/helm-%s-%s-%s.tar.gz", release.Tag(), input.Goos, input.Goarch)
	} else {
		assets, err := release.Assets()
		if err != nil {
			return nil, err
		}
		asset, err := findAsset(repo, release, assets, input.Goos, input.Goarch)
		if err != nil {
			return nil, err
		}
		downloadURL = asset.DownloadURL()
	}

	_, assetName := path.Split(downloadURL)

	binaryName, err := getBinaryName(repo.Owner(), repo.Name(), input.Goos)
	if err != nil {
		return nil, err
	}

	return &FindOutput{
		Owner:       repo.Owner(),
		Repo:        repo.Name(),
		Tag:         release.Tag(),
		Asset:       assetName,
		DownloadURL: downloadURL,
		BinaryName:  binaryName,
		IsArchived:  isArchived(assetName),
	}, nil
}

func parse(name string) (string, string, string, error) {
	re := regexp.MustCompile(`(([^/=]+)/)?([^/=]+)(=([^=]+))?`)
	if re.MatchString(name) {
		match := re.FindStringSubmatch(name)
		return match[2], match[3], match[5], nil
	}
	return "", "", "", fmt.Errorf("parsing package name failed: %s\npackage name should be \"owner/repo=tag\" format", name)
}

func findAsset(repo github.Repository, release github.Release, assets []github.Asset, goos, goarch string) (github.Asset, error) {
	filteredAssets := []github.Asset{}
	for _, asset := range assets {
		if isAssetNamePredefined(repo.Owner(), repo.Name()) {
			assetName, err := getPredefinedAssetName(repo.Owner(), repo.Name(), release.Tag(), goos, goarch)
			if err != nil {
				return nil, err
			}
			if asset.Name() == assetName {
				filteredAssets = append(filteredAssets, asset)
			}
		} else {
			if isNotReleaseBinary(asset.Name()) {
				continue
			}
			assetGoos, err := guessGoos(asset.Name())
			if err != nil {
				continue
			}
			assetGoarch, err := guessGoarch(asset.Name())
			if err != nil {
				continue
			}
			if assetGoos == goos && assetGoarch == goarch {
				filteredAssets = append(filteredAssets, asset)
			}
		}
	}

	if len(filteredAssets) == 0 {
		return nil, fmt.Errorf("no asset found")
	} else if len(filteredAssets) >= 2 {
		assetNames := []string{}
		for _, asset := range filteredAssets {
			assetNames = append(assetNames, asset.Name())
		}
		return nil, fmt.Errorf("too many assets found: %v", strings.Join(assetNames, ", "))
	}

	return filteredAssets[0], nil
}

func isAssetNamePredefined(owner, repo string) bool {
	key := fmt.Sprintf("%s/%s", owner, repo)
	_, ok := assetMap[key]
	return ok
}

func getPredefinedAssetName(owner, repo, tag, goos, goarch string) (string, error) {
	key := fmt.Sprintf("%s/%s", owner, repo)
	tmplMap, ok := assetMap[key]
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

	version := strings.TrimLeft(tag, "v")

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
		Tag:     tag,
		Version: version,
		Goos:    goos,
		Goarch:  goarch,
	})
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}

func isNotReleaseBinary(asset string) bool {
	return hasExt(asset, []string{".sha256", ".deb", ".rpm", ".msi"})
}

func guessGoos(asset string) (string, error) {
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
	case strings.Contains(s, "macos"):
		return "darwin", nil
	case strings.Contains(s, "win"):
		return "windows", nil
	default:
		return "", fmt.Errorf("fail to guess GOOS from asset: %s", asset)
	}
}

func guessGoarch(asset string) (string, error) {
	s := strings.ToLower(asset)
	switch {
	case strings.Contains(s, "amd64"):
		return "amd64", nil
	case strings.Contains(s, "x86_64"):
		return "amd64", nil
	default:
		return "", fmt.Errorf("fail to guess GOARCH from asset: %s", asset)
	}
}

func getBinaryName(owner, repo, goos string) (string, error) {
	if isBinaryNamePredefined(owner, repo) {
		return getPredefinedBinaryName(owner, repo, goos)
	}
	if goos == "windows" {
		return fmt.Sprintf("%s.exe", repo), nil
	}
	return repo, nil
}

func isBinaryNamePredefined(owner, repo string) bool {
	key := fmt.Sprintf("%s/%s", owner, repo)
	_, ok := binaryMap[key]
	return ok
}

func getPredefinedBinaryName(owner, repo, goos string) (string, error) {
	key := fmt.Sprintf("%s/%s", owner, repo)
	if bin, ok := binaryMap[key]; ok {
		if goos == "windows" {
			return fmt.Sprintf("%s.exe", bin), nil
		}
		return bin, nil
	}
	return "", fmt.Errorf("binary name is not predefined: %s", key)
}

func isArchived(asset string) bool {
	return hasExt(asset, []string{".tar", ".gz", ".tgz", ".bz2", ".tbz", ".Z", ".zip", ".bz2", ".lzh", ".7z", ".gz", ".rar", ".cab", ".afz"})
}

func hasExt(name string, exts []string) bool {
	for _, ext := range exts {
		if filepath.Ext(name) == ext {
			return true
		}
	}
	return false
}
