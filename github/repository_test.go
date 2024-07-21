package github

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewRepositoryFromFullName(t *testing.T) {
	tests := []struct {
		name     string
		fullName string
		repo     Repository
	}{
		{
			name:     "hashicorp/terraform",
			fullName: "hashicorp/terraform",
			repo:     newRepository("hashicorp", "terraform"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require := require.New(t)
			repo, err := newRepositoryFromFullName(tt.fullName)
			require.NoError(err)
			require.Equal(tt.repo, repo)
		})
	}
}

func TestRepositoryUnmarshalCSV(t *testing.T) {
	tests := []struct {
		name  string
		value string
		repo  Repository
	}{
		{
			name:  "hashicorp/terraform",
			value: "hashicorp/terraform",
			repo:  newRepository("hashicorp", "terraform"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require := require.New(t)
			repo := Repository{}
			repo.UnmarshalCSV(tt.value)
			require.Equal(tt.repo, repo)
		})
	}
}

// UnmarshalCSV converts the CSV string as GitHub repository.
func (r *Repository) UnmarshalCSV(value string) error {
	repo, err := newRepositoryFromFullName(value)
	if err != nil {
		return err
	}
	r.owner, r.name = repo.owner, repo.name
	return nil
}
