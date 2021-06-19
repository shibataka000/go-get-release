package pkg

import (
	"os"
	"reflect"
	"testing"
)

func TestFind(t *testing.T) {
	tests := []struct {
		input  *FindInput
		output *FindOutput
	}{
		{
			input: &FindInput{
				Name:        "shibataka000/go-get-release-test",
				GithubToken: os.Getenv("GITHUB_PERSONAL_ACCESS_TOKEN"),
				Goos:        "linux",
				Goarch:      "amd64",
			},
			output: &FindOutput{
				Owner:       "shibataka000",
				Repo:        "go-get-release-test",
				Tag:         "v0.0.2",
				Asset:       "go-get-release_v0.0.2_linux_amd64",
				DownloadURL: "https://github.com/shibataka000/go-get-release-test/releases/download/v0.0.2/go-get-release_v0.0.2_linux_amd64",
				BinaryName:  "go-get-release-test",
				IsArchived:  false,
			},
		},
		{
			input: &FindInput{
				Name:        "shibataka000/go-get-release-test=v0.0.1",
				GithubToken: os.Getenv("GITHUB_PERSONAL_ACCESS_TOKEN"),
				Goos:        "linux",
				Goarch:      "amd64",
			},
			output: &FindOutput{
				Owner:       "shibataka000",
				Repo:        "go-get-release-test",
				Tag:         "v0.0.1",
				Asset:       "go-get-release_v0.0.1_linux_amd64",
				DownloadURL: "https://github.com/shibataka000/go-get-release-test/releases/download/v0.0.1/go-get-release_v0.0.1_linux_amd64",
				BinaryName:  "go-get-release-test",
				IsArchived:  false,
			},
		},
		{
			input: &FindInput{
				Name:        "terraform=v0.12.20",
				GithubToken: os.Getenv("GITHUB_PERSONAL_ACCESS_TOKEN"),
				Goos:        "linux",
				Goarch:      "amd64",
			},
			output: &FindOutput{
				Owner:       "hashicorp",
				Repo:        "terraform",
				Tag:         "v0.12.20",
				Asset:       "terraform_0.12.20_linux_amd64.zip",
				DownloadURL: "https://releases.hashicorp.com/terraform/0.12.20/terraform_0.12.20_linux_amd64.zip",
				BinaryName:  "terraform",
				IsArchived:  true,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.input.Name, func(t *testing.T) {
			actual, err := Find(tt.input)
			if err != nil {
				t.Fatal(err)
			}
			if !reflect.DeepEqual(actual, tt.output) {
				t.Fatalf("Expected is %v but actual is %v", tt.output, actual)
			}
		})
	}
}

func TestParse(t *testing.T) {
	tests := []struct {
		in  string
		out []string
	}{
		{
			in:  "go-get-release",
			out: []string{"", "go-get-release", ""},
		},
		{
			in:  "shibataka000/go-get-release",
			out: []string{"shibataka000", "go-get-release", ""},
		},
		{
			in:  "go-get-release=1.0.0",
			out: []string{"", "go-get-release", "1.0.0"},
		},
		{
			in:  "shibataka000/go-get-release=1.0.0",
			out: []string{"shibataka000", "go-get-release", "1.0.0"},
		},
		{
			in:  "shibataka000/go-get-release=alpha/1.0.0",
			out: []string{"shibataka000", "go-get-release", "alpha/1.0.0"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.in, func(t *testing.T) {
			owner, repo, version, err := parse(tt.in)
			if err != nil {
				t.Fatal(err)
			}
			actual := []string{owner, repo, version}
			if !reflect.DeepEqual(actual, tt.out) {
				t.Fatalf("Expected is %v but actual is %v", tt.out, actual)
			}
		})
	}
}

func TestIsArchived(t *testing.T) {
	tests := []struct {
		in  string
		out bool
	}{
		{
			in:  "go-get-release-test_v0.0.1_linux_amd64",
			out: false,
		},
		{
			in:  "go-get-release-test_v0.0.1_linux_amd64.zip",
			out: true,
		},
		{
			in:  "go-get-release-test_v0.0.1_linux_amd64.tar.gz",
			out: true,
		},
		{
			in:  "go-get-release-test_v0.0.1_linux_amd64.exe",
			out: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.in, func(t *testing.T) {
			actual := isArchived(tt.in)
			if actual != tt.out {
				t.Fatalf("Expected is %v but actual is %v", tt.out, actual)
			}
		})
	}
}
