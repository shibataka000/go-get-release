package cmd

import (
	"fmt"

	"github.com/shibataka000/go-get-release/pkg/pkg"
)

func Search(pkgName string, option *Option) error {
	pkgInfos, err := pkg.Search(&pkg.SearchInput{
		Name:        pkgName,
		GithubToken: option.GithubToken,
	})
	if err != nil {
		return err
	}

	for _, pkgInfo := range pkgInfos {
		fmt.Printf("%s/%s\n", pkgInfo.Owner, pkgInfo.Repo)
	}
	return nil
}
