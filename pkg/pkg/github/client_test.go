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
