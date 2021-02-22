package github

import (
	"os"
	"testing"
)

func TestFindRepository(t *testing.T) {
	tests := []struct {
		keyword string
		owner   string
		repo    string
	}{
		{
			keyword: "terraform",
			owner:   "hashicorp",
			repo:    "terraform",
		},
	}

	token := os.Getenv("GITHUB_PERSONAL_ACCESS_TOKEN")
	c, err := NewClient(token)
	if err != nil {
		t.Error(err)
		return
	}

	for _, tt := range tests {
		t.Run(tt.keyword, func(t *testing.T) {
			repo, err := c.FindRepository(tt.keyword)
			if err != nil {
				t.Error(err)
				return
			}
			if repo.Owner() != tt.owner || repo.Name() != tt.repo {
				t.Errorf("Expected is %s/%s but actual is %s/%s", tt.owner, tt.repo, repo.Owner(), repo.Name())
				return
			}
		})
	}
}

func TestSearchRepository(t *testing.T) {
	tests := []struct {
		keyword string
		output  []repository
		length  int
	}{
		{
			keyword: "terraform",
			output: []repository{
				{
					owner: "hashicorp",
					name:  "terraform",
				},
			},
			length: 30,
		},
	}

	token := os.Getenv("GITHUB_PERSONAL_ACCESS_TOKEN")
	c, err := NewClient(token)
	if err != nil {
		t.Error(err)
		return
	}

	for _, tt := range tests {
		t.Run(tt.keyword, func(t *testing.T) {
			actual, err := c.SearchRepositories(tt.keyword)
			if err != nil {
				t.Error(err)
				return
			}
			if len(actual) != tt.length {
				t.Errorf("Excepted length is %d but actual length is %d", tt.length, len(actual))
				return
			}
			for i := range tt.output {
				if actual[i].Owner() != tt.output[i].Owner() || actual[i].Name() != tt.output[i].Name() {
					t.Errorf("Expected is %s/%s but actual is %s/%s", tt.output[i].Owner(), tt.output[i].Name(), actual[i].Owner(), actual[i].Name())
					return
				}
			}
		})
	}
}
