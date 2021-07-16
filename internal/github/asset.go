package github

import (
	"fmt"
	"sort"
	"strings"
)

// Asset in GitHub repository
type Asset interface {
	Name() string
	DownloadURL() string
	Goos() (string, error)
	Goarch() (string, error)
}

type asset struct {
	name        string
	downloadURL string
}

func (a *asset) Name() string {
	return a.name
}

func (a *asset) DownloadURL() string {
	return a.downloadURL
}

func (a *asset) Goos() (string, error) {
	gooses := listGoos()
	sort.Slice(gooses, func(i, j int) bool { return len(gooses[i]) > len(gooses[j]) })

	name := strings.ToLower(a.Name())
	for _, goos := range gooses {
		if strings.Contains(name, goos) {
			return goos, nil
		}
	}

	switch {
	case strings.Contains(name, "macos"):
		return "darwin", nil
	case strings.Contains(name, "osx"):
		return "darwin", nil
	case strings.Contains(name, "win"):
		return "windows", nil
	default:
		return "", fmt.Errorf("fail to guess GOOS from asset name: %s", name)
	}
}

func (a *asset) Goarch() (string, error) {
	goarches := listGoarch()
	sort.Slice(goarches, func(i, j int) bool { return len(goarches[i]) > len(goarches[j]) })

	name := strings.ToLower(a.Name())
	for _, goarch := range goarches {
		if strings.Contains(name, goarch) {
			return goarch, nil
		}
	}

	switch {
	case strings.Contains(name, "x86_64"):
		return "amd64", nil
	default:
		return "", fmt.Errorf("fail to guess GOARCH from asset name: %s", name)
	}
}
