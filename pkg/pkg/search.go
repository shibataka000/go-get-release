package pkg

import (
	"github.com/shibataka000/go-get-release/pkg/pkg/github"
)

type SearchInput struct {
	Name        string
	GithubToken string
}

type SearchOutput []SearchOutputItem

type SearchOutputItem struct {
	Owner string
	Repo  string
}

func Search(input *SearchInput) (SearchOutput, error) {
	client, err := github.NewClient(input.GithubToken)
	if err != nil {
		return nil, err
	}

	repos, err := client.SearchRepositories(input.Name)
	if err != nil {
		return nil, err
	}

	output := SearchOutput{}
	for _, repo := range repos {
		output = append(output, SearchOutputItem{
			Owner: repo.Owner(),
			Repo:  repo.Name(),
		})
	}
	return output, nil
}
