package github

import (
	"os"
	"testing"
)

func TestGetLatestRelease(t *testing.T) {
	tests := []struct {
		description string
		owner       string
		repo        string
		tag         string
	}{
		{
			description: "shibataka000/go-get-release",
			owner:       "shibataka000",
			repo:        "go-get-release-test",
			tag:         "v0.0.2",
		},
	}

	token := os.Getenv("GITHUB_PERSONAL_ACCESS_TOKEN")
	c, err := NewClient(token)
	if err != nil {
		t.Error(err)
		return
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			repo, err := c.GetRepository(tt.owner, tt.repo)
			if err != nil {
				t.Error(err)
				return
			}
			release, err := repo.GetLatestRelease()
			if err != nil {
				t.Error(err)
				return
			}
			if release.Tag() != tt.tag {
				t.Errorf("Expected is %s but actual is %s", tt.tag, release.Tag())
				return
			}
		})
	}
}
func TestGetRelease(t *testing.T) {
	tests := []struct {
		description string
		owner       string
		repo        string
		tag         string
	}{
		{
			description: "shibataka000/go-get-release",
			owner:       "shibataka000",
			repo:        "go-get-release-test",
			tag:         "v0.0.1",
		},
		{
			description: "shibataka000/go-get-release",
			owner:       "shibataka000",
			repo:        "go-get-release-test",
			tag:         "v0.0.2",
		},
	}

	token := os.Getenv("GITHUB_PERSONAL_ACCESS_TOKEN")
	c, err := NewClient(token)
	if err != nil {
		t.Error(err)
		return
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			repo, err := c.GetRepository(tt.owner, tt.repo)
			if err != nil {
				t.Error(err)
				return
			}
			release, err := repo.GetRelease(tt.tag)
			if err != nil {
				t.Error(err)
				return
			}
			if release.Tag() != tt.tag {
				t.Errorf("Expected is %s but actual is %s", tt.tag, release.Tag())
				return
			}
		})
	}
}
