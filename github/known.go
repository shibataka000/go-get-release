package github

import "github.com/shibataka000/go-get-release/slices"

type KnownItem struct {
	Repository Repository
	Assets     []AssetMetaTemplate
	ExecBinary ExecBinaryMetaTemplate
}

type KnownItemdSet []KnownItem

func (s KnownItemdSet) Find(repo Repository) (KnownItem, error) {
	return slices.Find(s, func(e KnownItem) bool {
		return e.Repository.Equal(repo)
	})
}
