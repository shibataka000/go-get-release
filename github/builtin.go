package github

import (
	_ "embed"
	"slices"

	"github.com/shibataka000/go-get-release/file"
	"github.com/shibataka000/go-get-release/platform"
	"github.com/shibataka000/go-get-release/url"
)

//go:embed builtin.yaml
var builtin []byte

type BuiltIn struct {
	Repository       Repository
	Assets           AssetTemplateList
	ExecutableBinary ExecutableBinaryTempalte
}

type BuiltInList []BuiltIn

func (s BuiltInList) find(repo Repository) (BuiltIn, error) {
	index := slices.IndexFunc(s, func(r BuiltIn) bool {
		return r.Repository.owner == repo.owner && r.Repository.name == repo.name
	})
	if index == -1 {
		return BuiltIn{}, &AssetNotFoundError{}
	}
	return s[index], nil
}

type AssetTemplate struct {
	DownloadURL url.Template
}

func (a AssetTemplate) render(release Release) (url.URL, error) {
	return "", nil
}

type AssetTemplateList []AssetTemplate

func (a AssetTemplateList) render(release Release) ([]url.URL, error) {
	return nil, nil
}

type ExecutableBinaryTempalte struct {
	BaseName file.Name
}

func (b ExecutableBinaryTempalte) render(os platform.OS) ExecutableBinary {
	return newExecutableBinary(b.BaseName, os)
}

func loadBuiltInData() (BuiltInList, error) {
	return nil, nil
}
