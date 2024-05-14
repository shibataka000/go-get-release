package api

import (
	"context"
	"strings"

	"github.com/google/go-github/v48/github"
	"github.com/shibataka000/go-get-release/runtime"
	"github.com/shibataka000/go-get-release/url"
)

// AssetMeta represents a GitHub release asset in a repository.
type AssetMeta struct {
	DownloadURL url.URL
}

// AssetMetaList is a list of AssetMeta.
type AssetMetaList []AssetMeta

// newAssetMeat return new AssetMeta object.
func newAssetMeta(downloadURL url.URL) AssetMeta {
	return AssetMeta{
		DownloadURL: downloadURL,
	}
}

// GOOS return GOOS guessed by asset file name.
// Default is "unknown".
func (a AssetMeta) GOOS() runtime.GOOS {
	// These are listed by following command.
	// `go tool dist list | sed -r "s/(\w+)\/(\w+)/\1/g" | sort | uniq`
	platforms := map[string][]string{
		"aix":       {"aix"},
		"android":   {"android"},
		"darwin":    {"darwin", "macos", "osx"},
		"dragonfly": {"dragonfly"},
		"freebsd":   {"freebsd"},
		"illumos":   {"illumos"},
		"ios":       {"ios"},
		"js":        {"js"},
		"linux":     {"linux"},
		"netbsd":    {"netbsd"},
		"openbsd":   {"openbsd"},
		"plan9":     {"plan9"},
		"solaris":   {"solaris"},
		"windows":   {"windows", "win", ".exe"},
	}
	lowner := strings.ToLower(a.DownloadURL.Base())
	os, err := findKeyWhichHasLongestMatchValue(platforms, lowner)
	if err != nil {
		return "unknown"
	}
	return runtime.GOOS(os)
}

// GOARCH return GOARCH guessed by asset file name.
// Default is "amd64".
func (a AssetMeta) GOARCH() runtime.GOARCH {
	// These are listed by following command.
	// `go tool dist list | sed -r "s/(\w+)\/(\w+)/\2/g" | sort | uniq`
	platforms := map[string][]string{
		"386":      {"386", "x86_32", "32bit", "win32"},
		"686":      {"686"},
		"amd64":    {"amd64", "x86_64", "64bit", "win64"},
		"arm":      {"arm"},
		"arm64":    {"arm64", "aarch64", "aarch_64"},
		"mips":     {"mips"},
		"mips64":   {"mips64"},
		"mips64le": {"mips64le"},
		"mipsle":   {"mipsle"},
		"ppc64":    {"ppc64"},
		"ppc64le":  {"ppc64le", "ppcle_64"},
		"riscv64":  {"riscv64"},
		"s390x":    {"s390x", "s390"},
		"wasm":     {"wasm"},
	}
	lowner := strings.ToLower(a.DownloadURL.Base())
	arch, err := findKeyWhichHasLongestMatchValue(platforms, lowner)
	if err != nil {
		return "amd64"
	}
	return runtime.GOARCH(arch)
}

// findKeyWhichHasLongestMatchValue return key in map which has longest matched value.
func findKeyWhichHasLongestMatchValue(m map[string][]string, value string) (string, error) {
	matchKey, matchValue := "", ""
	for key, values := range m {
		for _, v := range values {
			if strings.Contains(value, v) && len(matchValue) < len(v) {
				matchKey = key
				matchValue = v
			}
		}
	}
	if matchKey == "" {
		return "", &NotMatchedError{}
	}
	return matchKey, nil
}

// NotMatchedError is error raised when none of elements in map was matched to value in findKeyWhichHasLongestMatchValue function.
type NotMatchedError struct{}

// Error return error string.
func (e *NotMatchedError) Error() string {
	return "Not matched."
}

// find AssetMeta by GOOS/GOARCH.
func (a AssetMetaList) find(goos runtime.GOOS, goarch runtime.GOARCH) (AssetMeta, error) {
	for _, asset := range a {
		if asset.GOOS() == goos && asset.GOARCH() == goarch {
			return asset, nil
		}
	}
	return AssetMeta{}, &AssetNotFoundError{}
}

// AssetNotFoundError is error raised when try to find AssetMeta by GOOS/GOARCH but no AssetMeta was found.
type AssetNotFoundError struct{}

// Error return error string.
func (e *AssetNotFoundError) Error() string {
	return "No asset was found."
}

// AssetRepository is repository for Asset.
type AssetRepository struct {
	client *github.Client
}

// NewAssetRepository return new AssetRepository object.
func NewAssetRepository(ctx context.Context, token string) *AssetRepository {
	return &AssetRepository{
		client: newGitHubClient(ctx, token),
	}
}

// list return lit of AssetMeta in a GitHub release.
func (r *AssetRepository) list(ctx context.Context, repo Repository, release Release) (AssetMetaList, error) {
	githubRelease, _, err := r.client.Repositories.GetReleaseByTag(ctx, repo.Owner, repo.Name, release.Tag)
	if err != nil {
		return nil, err
	}

	result := AssetMetaList{}
	for page := 1; page != 0; {
		assets, resp, err := r.client.Repositories.ListReleaseAssets(ctx, repo.Owner, repo.Name, *githubRelease.ID, &github.ListOptions{
			Page:    page,
			PerPage: 100,
		})
		if err != nil {
			return result, err
		}
		for _, asset := range assets {
			downloadURL := url.URL(asset.GetBrowserDownloadURL())
			result = append(result, newAssetMeta(downloadURL))
		}
		page = resp.NextPage
	}
	return result, nil
}
