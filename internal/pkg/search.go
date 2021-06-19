package pkg

import (
	"github.com/shibataka000/go-get-release/internal/github"
)

// SearcgInput is input to Search function
type SearchInput struct {
	Name        string
	GithubToken string
}

// FindOutput is output by Search function
type SearchOutput []SearchOutputItem

// FindOutputItem is element of SearchOutput struct
type SearchOutputItem struct {
	Owner       string
	Repo        string
	Description string
}

// Search GitHub repositories
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
			Owner:       repo.Owner(),
			Repo:        repo.Name(),
			Description: repo.Description(),
		})
	}
	return output, nil
}
