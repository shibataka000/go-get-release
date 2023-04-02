package pkg

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v3"
)

func LoadIndexForTest(t *testing.T) (Index, error) {
	t.Helper()
	indexBytes, err := os.ReadFile("./testdata/index.yaml")
	if err != nil {
		return Index{}, err
	}
	repos := []RepositoryInIndex{}
	err = yaml.Unmarshal(indexBytes, &repos)
	if err != nil {
		return Index{}, err
	}
	return NewIndex(repos), nil
}

func TestIndexFindRepository(t *testing.T) {
	tests := []struct {
		name       string
		githubRepo GitHubRepository
		indexRepo  RepositoryInIndex
	}{
		{
			name:       "hashicorp/terraform",
			githubRepo: NewGitHubRepository("hashicorp", "terraform"),
			indexRepo: NewRepositoryInIndex("hashicorp", "terraform", []AssetInIndex{
				NewAssetInIndex("https://releases.hashicorp.com/terraform/{{.SemVer}}/terraform_{{.SemVer}}_linux_amd64.zip", "linux", "amd64"),
				NewAssetInIndex("https://releases.hashicorp.com/terraform/{{.SemVer}}/terraform_{{.SemVer}}_darwin_amd64.zip", "darwin", "amd64"),
				NewAssetInIndex("https://releases.hashicorp.com/terraform/{{.SemVer}}/terraform_{{.SemVer}}_windows_amd64.zip", "windows", "amd64"),
			}, NewExecBinaryInIndex("terraform")),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := require.New(t)
			index, err := LoadIndexForTest(t)
			assert.NoError(err)
			repo, err := index.FindRepository(tt.githubRepo)
			assert.NoError(err)
			assert.Equal(tt.indexRepo, repo)
		})
	}
}

func TestIndexFindAsset(t *testing.T) {
	tests := []struct {
		name       string
		repository GitHubRepository
		platform   Platform
		asset      AssetInIndex
	}{
		{
			name:       "hashicorp/terraform",
			repository: NewGitHubRepository("hashicorp", "terraform"),
			platform:   NewPlatform("linux", "amd64"),
			asset:      NewAssetInIndex("https://releases.hashicorp.com/terraform/{{.SemVer}}/terraform_{{.SemVer}}_linux_amd64.zip", "linux", "amd64"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := require.New(t)
			index, err := LoadIndexForTest(t)
			assert.NoError(err)
			asset, err := index.FindAsset(tt.repository, tt.platform)
			assert.NoError(err)
			assert.Equal(tt.asset, asset)
		})
	}
}

func TestIndexHasAsset(t *testing.T) {
	tests := []struct {
		name       string
		repository GitHubRepository
		platform   Platform
		hasAsset   bool
	}{
		{
			name:       "hashicorp/terraform",
			repository: NewGitHubRepository("hashicorp", "terraform"),
			platform:   NewPlatform("linux", "amd64"),
			hasAsset:   true,
		},
		{
			name:       "shibataka000/go-get-release",
			repository: NewGitHubRepository("shibataka000", "go-get-release"),
			platform:   NewPlatform("linux", "amd64"),
			hasAsset:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := require.New(t)
			index, err := LoadIndexForTest(t)
			assert.NoError(err)
			hasAsset := index.HasAsset(tt.repository, tt.platform)
			assert.Equal(tt.hasAsset, hasAsset)
		})
	}
}

func TestIndexFindExecBinary(t *testing.T) {
	tests := []struct {
		name       string
		repository GitHubRepository
		execBinary ExecBinaryInIndex
	}{
		{
			name:       "hashicorp/terraform",
			repository: NewGitHubRepository("hashicorp", "terraform"),
			execBinary: NewExecBinaryInIndex("terraform"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := require.New(t)
			index, err := LoadIndexForTest(t)
			assert.NoError(err)
			execBinary, err := index.FindExecBinary(tt.repository)
			assert.NoError(err)
			assert.Equal(tt.execBinary, execBinary)
		})
	}
}

func TestIndexHasExecBinary(t *testing.T) {
	tests := []struct {
		name          string
		repository    GitHubRepository
		hasExecBinary bool
	}{
		{
			name:          "hashicorp/terraform",
			repository:    NewGitHubRepository("hashicorp", "terraform"),
			hasExecBinary: true,
		},
		{
			name:          "shibataka000/go-get-release",
			repository:    NewGitHubRepository("shibataka000", "go-get-release"),
			hasExecBinary: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := require.New(t)
			index, err := LoadIndexForTest(t)
			assert.NoError(err)
			hasExecBinary := index.HasExecBinary(tt.repository)
			assert.Equal(tt.hasExecBinary, hasExecBinary)
		})
	}
}

func TestRepositoryInIndexEquals(t *testing.T) {
	tests := []struct {
		name       string
		indexRepo  RepositoryInIndex
		githubRepo GitHubRepository
		equals     bool
	}{
		{
			name:       "hashicorp/terraform",
			indexRepo:  NewRepositoryInIndex("hashicorp", "terraform", []AssetInIndex{}, NewExecBinaryInIndex("")),
			githubRepo: NewGitHubRepository("hashicorp", "terraform"),
			equals:     true,
		},
		{
			name:       "hashicorp/terraform",
			indexRepo:  NewRepositoryInIndex("hashicorp", "terraform", []AssetInIndex{}, NewExecBinaryInIndex("")),
			githubRepo: NewGitHubRepository("hashicorp", "vault"),
			equals:     false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := require.New(t)
			equals := tt.indexRepo.Equals(tt.githubRepo)
			assert.Equal(tt.equals, equals)
		})
	}
}

func TestRepositoryInIndexFindAsset(t *testing.T) {
	tests := []struct {
		name       string
		repository RepositoryInIndex
		platform   Platform
		asset      AssetInIndex
	}{
		{
			name: "hashicorp/terraform",
			repository: NewRepositoryInIndex("hashicorp", "terraform", []AssetInIndex{
				NewAssetInIndex("https://releases.hashicorp.com/terraform/{{.SemVer}}/terraform_{{.SemVer}}_linux_amd64.zip", "linux", "amd64"),
				NewAssetInIndex("https://releases.hashicorp.com/terraform/{{.SemVer}}/terraform_{{.SemVer}}_darwin_amd64.zip", "darwin", "amd64"),
				NewAssetInIndex("https://releases.hashicorp.com/terraform/{{.SemVer}}/terraform_{{.SemVer}}_windows_amd64.zip", "windows", "amd64"),
			}, NewExecBinaryInIndex("terraform")),
			platform: NewPlatform("linux", "amd64"),
			asset:    NewAssetInIndex("https://releases.hashicorp.com/terraform/{{.SemVer}}/terraform_{{.SemVer}}_linux_amd64.zip", "linux", "amd64"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := require.New(t)
			asset, err := tt.repository.FindAsset(tt.platform)
			assert.NoError(err)
			assert.Equal(tt.asset, asset)
		})
	}
}

func TestExecBinaryInIndexIsEmpty(t *testing.T) {
	tests := []struct {
		name       string
		execBinary ExecBinaryInIndex
		isEmpty    bool
	}{
		{
			name:       "terraform",
			execBinary: NewExecBinaryInIndex("terraform"),
			isEmpty:    false,
		},
		{
			name:       "empty",
			execBinary: NewExecBinaryInIndex(""),
			isEmpty:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := require.New(t)
			assert.Equal(tt.isEmpty, tt.execBinary.IsEmpty())
		})
	}
}
