package main

import (
	"fmt"

	"github.com/shibataka000/go-get-release/internal/github"
)

func search(name, token string) error {
	client, err := github.NewClient(token)
	if err != nil {
		return err
	}

	repos, err := client.SearchRepositories(name)
	if err != nil {
		return err
	}

	width := 0
	for _, repo := range repos {
		s := fmt.Sprintf("%s/%s", repo.Owner(), repo.Name())
		if len(s) > width {
			width = len(s)
		}
	}

	for _, repo := range repos {
		description := repo.Description()
		if len(description) > 100 {
			description = fmt.Sprintf("%s...", description[:100])
		}

		s := fmt.Sprintf("%s/%s", repo.Owner(), repo.Name())
		fmt.Printf("* %-*s - %s\n", width, s, description)
	}
	return nil
}
