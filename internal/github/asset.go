package github

import (
	"fmt"
	"path"
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
	ContainReleaseBinary() bool
	IsArchived() bool
	IsExecBinary() bool
}

type asset struct {
	client      *client
	repo        *repository
	release     *release
	downloadURL string
}

// Name return asset name
func (a *asset) Name() string {
	_, f := path.Split(a.downloadURL)
	return f
}

// DownloadURL return asset's download URL
func (a *asset) DownloadURL() string {
	return a.downloadURL
}

// BinaryName return asset's binary name
func (a *asset) BinaryName() (string, error) {
	binaryName := a.repo.Name()
	key := fmt.Sprintf("%s/%s", a.repo.Owner(), a.repo.Name())
	if v, ok := binaryNameMap[key]; ok {
		binaryName = v
	}

	goos, err := a.Goos()
	if err != nil {
		return "", err
	}
	ext := ""
	if goos == "windows" {
		ext = ".exe"
	}

	return fmt.Sprintf("%s%s", binaryName, ext), nil
}

// Goos return asset's goos which is guessed from asset name
func (a *asset) Goos() (string, error) {
	goos, err := findPlatform(a.Name(), goosMap)
	if err != nil {
		return "", fmt.Errorf("fail to guess GOOS from asset name: %s", a.Name())
	}
	return goos, nil
}

// Goarch return asset's goarch which is guessed from asset name
func (a *asset) Goarch() (string, error) {
	goarch, err := findPlatform(a.Name(), goarchMap)
	if err != nil {
		return "amd64", nil
	}
	return goarch, nil
}

// ContainReleaseBinary return true if this asset contain release binary
func (a *asset) ContainReleaseBinary() bool {
	return a.IsArchived() || a.IsExecBinary()
}

// IsArchived check asset name and return true if it is archive file
func (a *asset) IsArchived() bool {
	return hasExt(a.Name(), []string{".tar", ".gz", ".tgz", ".bz2", ".tbz", ".Z", ".zip", ".bz2", ".lzh", ".7z", ".gz", ".rar", ".cab", ".afz", ".xz"})
}

// IsExecBinary check asset name and return true if it is executable binary
func (a *asset) IsExecBinary() bool {
	name := strings.ReplaceAll(a.Name(), a.release.Version(), "")
	exts := []string{"", ".exe"}
	if a.repo.Owner() == "mozilla" && a.repo.Name() == "sops" {
		for goos := range goosMap {
			exts = append(exts, fmt.Sprintf(".%s", goos))
		}
	}
	return hasExt(name, exts)
}

// hasExt return true if 'name' have specific extension which is in 'exts'
func hasExt(name string, exts []string) bool {
	for _, ext := range exts {
		if filepath.Ext(name) == ext {
			return true
		}
	}
	return false
}

// findPlatform find golang platform (GOOS/GOARCH) from platform map based on name.
func findPlatform(name string, platform map[string][]string) (string, error) {
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
