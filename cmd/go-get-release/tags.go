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

	fmt.Printf("Show release tags in '%s/%s' repository.\n", repo.Owner(), repo.Name())

	releases, err := repo.ListRelease(n)
	if err != nil {
		return err
	}

	for _, release := range releases {
		fmt.Printf("* %s\n", release.Tag())
	}

	return nil
}
