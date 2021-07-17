package github

import (
	"fmt"
	"path/filepath"
	"sort"
	"strings"
)

// Asset in GitHub repository
type Asset interface {
	Name() string
	DownloadURL() string
	Goos() (string, error)
	Goarch() (string, error)
	BinaryName() (string, error)
	IsReleaseBinary() bool
}

type asset struct {
	client      *client
	repo        *repository
	release     *release
	name        string
	downloadURL string
}

func (a *asset) Name() string {
	return a.name
}

func (a *asset) DownloadURL() string {
	return a.downloadURL
}

func (a *asset) BinaryName() (string, error) {
	name := a.repo.Name()
	key := fmt.Sprintf("%s/%s", a.repo.Owner(), a.repo.Name())
	if v, ok := binaryNameMap[key]; ok {
		name = v
	}

	goos, err := a.Goos()
	if err != nil {
		return "", err
	}
	ext := ""
	if goos == "windows" {
		ext = ".exe"
	}

	return fmt.Sprintf("%s%s", name, ext), nil
}

func (a *asset) Goos() (string, error) {
	name := strings.ToLower(a.Name())

	gooses := listGoos()
	sort.Slice(gooses, func(i, j int) bool { return len(gooses[i]) > len(gooses[j]) })
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
	name := strings.ToLower(a.Name())

	goarches := listGoarch()
	sort.Slice(goarches, func(i, j int) bool { return len(goarches[i]) > len(goarches[j]) })
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

func (a *asset) IsReleaseBinary() bool {
	binaryExts := []string{"", ".exe"}
	archivedExts := []string{".tar", ".gz", ".tgz", ".bz2", ".tbz", ".Z", ".zip", ".bz2", ".lzh", ".7z", ".gz", ".rar", ".cab", ".afz"}
	return hasExt(a.Name(), binaryExts) || hasExt(a.Name(), archivedExts)
}

func hasExt(name string, exts []string) bool {
	for _, ext := range exts {
		if filepath.Ext(name) == ext {
			return true
		}
	}
	return false
}
