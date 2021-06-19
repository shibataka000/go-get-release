package main

import (
	"fmt"

	"github.com/shibataka000/go-get-release/internal/pkg"
)

func search(name, token string) error {
	ps, err := pkg.Search(&pkg.SearchInput{
		Name:        name,
		GithubToken: token,
	})
	if err != nil {
		return err
	}

	for _, p := range ps {
		fmt.Printf("%s/%s\n", p.Owner, p.Repo)
	}
	return nil
}
