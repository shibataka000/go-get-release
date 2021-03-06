package pkg

import (
	"os"
	"reflect"
	"testing"
)

func TestSearch(t *testing.T) {
	tests := []struct {
		input  *SearchInput
		output SearchOutput
		length int
	}{
		{
			input: &SearchInput{
				Name:        "terraform",
				GithubToken: os.Getenv("GITHUB_PERSONAL_ACCESS_TOKEN"),
			},
			output: SearchOutput{
				{
					Owner:       "hashicorp",
					Repo:        "terraform",
					Description: "Terraform enables you to safely and predictably create, change, and improve infrastructure. It is an open source tool that codifies APIs into declarative configuration files that can be shared amongst team members, treated as code, edited, reviewed, and versioned.",
				},
			},
			length: 30,
		},
	}

	for _, tt := range tests {
		t.Run(tt.input.Name, func(t *testing.T) {
			actual, err := Search(tt.input)
			if err != nil {
				t.Fatal(err)
			}
			if len(actual) != tt.length {
				t.Fatalf("Excepted length is %d but actual length is %d", tt.length, len(actual))
			}
			for i := range tt.output {
				if !reflect.DeepEqual(actual[i], tt.output[i]) {
					t.Fatalf("Expected %d th element is %v but actual is %v", i, tt.output[i], actual[i])
				}
			}
		})
	}
}
