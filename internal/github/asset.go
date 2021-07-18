package github

import (
	"fmt"
	"path"
	"sort"
	"strings"

	"github.com/shibataka000/go-get-release/internal/file"
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

// Goos return asset's goos which is guessed by asset name
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
	case strings.HasSuffix(name, ".exe"):
		return "windows", nil
	default:
		return "", fmt.Errorf("fail to guess GOOS from asset name: %s", name)
	}
}

// Goarch return asset's goarch which is guessed by asset name
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

// IsReleaseBinary return true if thish asset contain release binary
func (a *asset) IsReleaseBinary() bool {
	name := strings.ReplaceAll(a.Name(), a.release.Version(), "")
	if a.repo.Owner() == "mozilla" && a.repo.Name() == "sops" {
		gooses := listGoos()
		for i := range gooses {
			gooses[i] = fmt.Sprintf(".%s", gooses[i])
		}
		return file.IsArchived(name) || file.IsExecBinary(name) || file.HasExt(name, gooses)
	}
	return file.IsArchived(name) || file.IsExecBinary(name)
}
