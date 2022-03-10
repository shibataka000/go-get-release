package main

import (
	"fmt"

	"github.com/shibataka000/go-get-release/internal/github"
)

func tags(name, token string, n int) error {
	owner, repoName, _, err := parse(name)
	if err != nil {
		return err
	}

	client, err := github.NewClient(token)
	if err != nil {
		return err
	}

	var repo github.Repository
	if owner != "" {
		repo, err = client.Repository(owner, repoName)
	} else {
		repo, err = client.FindRepository(repoName)
	}
	if err != nil {
		return err
	}

	fmt.Printf("%s/%s's release tags:\n", repo.Owner(), repo.Name())

	releases, err := repo.ListRelease(n)
	if err != nil {
		return err
	}

	width := 0
	for _, release := range releases {
		if len(release.Tag()) > width {
			width = len(release.Tag())
		}
	}

	for _, release := range releases {
		fmt.Printf("* %-*s (%s)\n", width, release.Tag(), release.PublishedAt().Format("2006.01.02"))
	}

	return nil
}
