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

	width := 0
	for _, p := range ps {
		s := fmt.Sprintf("%s/%s", p.Owner, p.Repo)
		n := len(s)
		if n > width {
			width = n
		}
	}

	for _, p := range ps {
		s := fmt.Sprintf("%s/%s", p.Owner, p.Repo)
		fmt.Printf("* %-*s - %s\n", width, s, p.Description)
	}
	return nil
}
