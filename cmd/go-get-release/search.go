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
		description := ""
		if len(p.Description) > 100 {
			description = fmt.Sprintf("%s...", p.Description[:100])
		} else {
			description = p.Description
		}

		s := fmt.Sprintf("%s/%s", p.Owner, p.Repo)
		fmt.Printf("* %-*s - %s\n", width, s, description)
	}
	return nil
}
