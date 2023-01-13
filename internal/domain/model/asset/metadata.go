package asset

import (
	"bytes"
	"fmt"
	"path"
	"path/filepath"
	"sort"
	"strings"
	"text/template"

	"github.com/shibataka000/go-get-release/internal/domain/model/platform"
	"golang.org/x/exp/slices"
)

// Metadata about GitHub asset.
type Metadata struct {
	downloadURL string
}

// NewMetadata return metadata instance about GitHub asset.
func NewMetadata(downloadURL string) Metadata {
	return Metadata{
		downloadURL: downloadURL,
	}
}

// NewMetadataFromTemplate return metadata instance about GitHub asset from download URL template.
func NewMetadataFromTemplate(downloadURLTemplate string, release Release) (Metadata, error) {
	downloadURL, err := renderTemplateWithRelease("DownloadURL", downloadURLTemplate, release)
	if err != nil {
		return Metadata{}, err
	}
	return NewMetadata(downloadURL), nil
}

// DownloadURL return URL to download asset.
func (m Metadata) DownloadURL() string {
	return m.downloadURL
}

// Name return asset Name.
func (m Metadata) Name() string {
	return path.Base(m.downloadURL)
}

// IsExecBinary return true if asset file is exec binary.
func (m Metadata) IsExecBinary() bool {
	exts := []string{"", ".exe"}
	return slices.Contains(exts, m.ext())
}

// IsArchived return true if asset file is archived.
func (m Metadata) IsArchived() bool {
	exts := []string{
		// Archived as tarball
		".tar", ".tgz", ".tar.gz", ".tbz2", ".tar.bz2", ".txz", ".tar.xz",
		// Archived by others
		".zip", ".lzh", ".rar", ".7z",
	}
	return slices.Contains(exts, m.ext())
}

// IsCompressed return true if asset file is compressed.
func (m Metadata) IsCompressed() bool {
	exts := []string{
		// Compressed by gzip
		".gz", ".tgz", ".tar.gz",
		// Compressed by bzip
		".bz2", ".tbz2", ".tar.bz2",
		// Compressed by xz
		".xz", ".txz", ".tar.xz",
		// Compressed by others
		".zip", ".lzh", ".rar", ".7z",
	}
	return slices.Contains(exts, m.ext())
}

// Platform return Platform instance guessed by asset file name.
func (m Metadata) Platform() (platform.Platform, error) {
	os, err := m.os()
	if err != nil {
		return platform.Platform{}, err
	}
	arch, err := m.arch()
	if err != nil {
		return platform.Platform{}, err
	}
	return platform.New(os, arch), nil
}

// ext return asset name extension.
func (m Metadata) ext() string {
	ext := filepath.Ext(m.Name())
	name := strings.TrimSuffix(m.Name(), ext)
	if filepath.Ext(name) == ".tar" {
		ext = ".tar" + ext
	}
	return ext
}

// os return GOOS guessed by asset file name.
func (m Metadata) os() (string, error) {
	// GOOS are listed by following command.
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
	lowner := strings.ToLower(m.Name())
	return findKeyWhichHasLongestMatchValue(platforms, lowner)
}

// arch return GOARCH guessed by asset file name.
// This return amd64 by default.
func (m Metadata) arch() (string, error) {
	// GOARCH are listed by following command.
	// `go tool dist list | sed -r "s/(\w+)\/(\w+)/\2/g" | sort | uniq`
	platforms := map[string][]string{
		"386":      {"386", "x86_32", "32bit", "win32"},
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
	lowner := strings.ToLower(m.Name())
	arch, err := findKeyWhichHasLongestMatchValue(platforms, lowner)
	if err != nil {
		return "amd64", nil
	}
	return arch, nil
}

// renderTemplateWithRelease render template text with GitHub release.
func renderTemplateWithRelease(name string, text string, release Release) (string, error) {
	semver, err := release.SemVer()
	if err != nil {
		return "", err
	}
	param := struct {
		Tag    string
		SemVer string
	}{
		Tag:    release.Tag(),
		SemVer: semver,
	}

	tmpl, err := template.New(name).Parse(text)
	if err != nil {
		return "", err
	}
	buf := new(bytes.Buffer)
	err = tmpl.Execute(buf, param)
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}

// findKeyWhichHasLongestMatchValue return key in map which has longest matched value.
func findKeyWhichHasLongestMatchValue(m map[string][]string, value string) (string, error) {
	values := []string{}
	for _, vs := range m {
		values = append(values, vs...)
	}
	sort.Slice(values, func(i, j int) bool { return len(values[i]) > len(values[j]) })

	longestMatchValue := ""
	found := false
	for _, v := range values {
		if strings.Contains(value, v) {
			longestMatchValue = v
			found = true
			break
		}
	}
	if !found {
		return "", fmt.Errorf("no value was matched")
	}

	for k, vs := range m {
		if slices.Contains(vs, longestMatchValue) {
			return k, nil
		}
	}
	return "", fmt.Errorf("no key was found")
}
