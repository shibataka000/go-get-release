package pkg

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"io/fs"
	"net/http"
	"os"
	"path/filepath"

	"github.com/cheggaaa/pb/v3"
	"github.com/google/go-github/v48/github"
	"golang.org/x/oauth2"
	"gopkg.in/yaml.v3"

	_ "embed"
)

//go:embed index.yaml
var builtInIndexData []byte

// Repository for package domain.
type Repository struct {
	github *github.Client
}

// NewRepository return new repository instance.
func NewRepository(ctx context.Context, token string) *Repository {
	var httpClient *http.Client
	if token != "" {
		tokenSource := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})
		httpClient = oauth2.NewClient(ctx, tokenSource)
	}
	githubClient := github.NewClient(httpClient)
	return &Repository{
		github: githubClient,
	}
}

// SearchGitHubRepository search GitHub repository.
func (r *Repository) SearchGitHubRepository(ctx context.Context, query string) (GitHubRepository, error) {
	result, _, err := r.github.Search.Repositories(ctx, query, &github.SearchOptions{})
	if err != nil {
		return GitHubRepository{}, err
	}
	repos := result.Repositories
	if len(repos) == 0 {
		return GitHubRepository{}, fmt.Errorf("no repository was found")
	}
	repo := repos[0]
	return NewGitHubRepository(repo.GetOwner().GetLogin(), repo.GetName()), nil
}

// FindGitHubRepository find GitHub repository.
func (r *Repository) FindGitHubRepository(ctx context.Context, owner string, name string) (GitHubRepository, error) {
	repo, _, err := r.github.Repositories.Get(ctx, owner, name)
	if err != nil {
		return GitHubRepository{}, err
	}
	return NewGitHubRepository(repo.GetOwner().GetLogin(), repo.GetName()), nil
}

// LatestGitHubRelease return latest GitHub release.
func (r *Repository) LatestGitHubRelease(ctx context.Context, repo GitHubRepository) (GitHubRelease, error) {
	release, _, err := r.github.Repositories.GetLatestRelease(ctx, repo.Owner, repo.Name)
	if err != nil {
		return GitHubRelease{}, err
	}
	return NewGitHubRelease(release.GetID(), release.GetTagName()), nil
}

// FindGitHubReleaseByTag return GitHub release by tag.
func (r *Repository) FindGitHubReleaseByTag(ctx context.Context, repo GitHubRepository, tag string) (GitHubRelease, error) {
	release, _, err := r.github.Repositories.GetReleaseByTag(ctx, repo.Owner, repo.Name, tag)
	if err != nil {
		return GitHubRelease{}, err
	}
	return NewGitHubRelease(release.GetID(), release.GetTagName()), nil
}

// ListGitHubAssets list assets in GitHub release.
func (r *Repository) ListGitHubAssets(ctx context.Context, repo GitHubRepository, release GitHubRelease) ([]GitHubAsset, error) {
	result := []GitHubAsset{}
	for page := 1; page != 0; {
		assets, resp, err := r.github.Repositories.ListReleaseAssets(ctx, repo.Owner, repo.Name, release.ID, &github.ListOptions{
			Page:    page,
			PerPage: 100,
		})
		if err != nil {
			return result, err
		}
		for _, asset := range assets {
			downloadURL := NewURL(asset.GetBrowserDownloadURL())
			result = append(result, NewGitHubAsset(downloadURL))
		}
		page = resp.NextPage
	}
	return result, nil
}

// LoadBuiltInIndex load and return built-in index.
func (r *Repository) LoadBuiltInIndex() (Index, error) {
	repos := []RepositoryInIndex{}
	err := yaml.Unmarshal(builtInIndexData, &repos)
	if err != nil {
		return Index{}, err
	}
	return NewIndex(repos), nil
}

// Download file.
func (r *Repository) Download(url URL, progressBar io.Writer) (File, error) {
	resp, err := http.Get(url.String())
	if err != nil {
		return File{}, err
	}
	defer resp.Body.Close()

	bar := pb.Full.Start64(resp.ContentLength).SetWriter(progressBar)
	src := bar.NewProxyReader(resp.Body)

	dst := new(bytes.Buffer)

	_, err = io.Copy(dst, src)
	if err != nil {
		return File{}, err
	}

	return NewFile(url.FileName(), dst.Bytes()), nil
}

// WriteFile write file to specified directory.
func (r *Repository) WriteFile(file File, dir string, perm fs.FileMode) error {
	path := filepath.Join(dir, file.Name.String())
	return os.WriteFile(path, file.Body, perm)
}
