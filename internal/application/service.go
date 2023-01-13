package application

import (
	"context"
	"fmt"

	"github.com/Songmu/prompter"
	assetfactory "github.com/shibataka000/go-get-release/internal/domain/factory/asset"
	execbinaryfactory "github.com/shibataka000/go-get-release/internal/domain/factory/execbinary"
	"github.com/shibataka000/go-get-release/internal/domain/model/execbinary"
	"github.com/shibataka000/go-get-release/internal/domain/model/release"
	"github.com/shibataka000/go-get-release/internal/domain/model/repository"
	assetservice "github.com/shibataka000/go-get-release/internal/domain/service/asset"
	"github.com/shibataka000/go-get-release/internal/infrastructure/embed"
	"github.com/shibataka000/go-get-release/internal/infrastructure/filesystem"
	"github.com/shibataka000/go-get-release/internal/infrastructure/github"
	"github.com/shibataka000/go-get-release/internal/infrastructure/http"
)

// Service is application service to install executable binary from GitHub release asset.
type Service struct {
	github     GitHubRepository
	http       HTTPRepository
	filesystem FileSystemRepository
	asset      AssetFactory
	execbinary ExecBinaryFactory
}

// NewService return application service to install executable binary from GitHub release asset.
func NewService(ctx context.Context, token string) (*Service, error) {
	githubRepository := github.New(ctx, token)
	embedRepository, err := embed.New()
	if err != nil {
		return &Service{}, err
	}

	return &Service{
		github:     githubRepository,
		http:       http.New(),
		filesystem: filesystem.New(),
		asset:      assetfactory.New(githubRepository, embedRepository, assetservice.New()),
		execbinary: execbinaryfactory.New(embedRepository),
	}, nil
}

// Install executable binary from GitHub release asset.
func (s *Service) Install(ctx context.Context, command Command) error {
	meta, err := s.findMetadata(ctx, command)
	if err != nil {
		return err
	}
	if command.isInteractive() {
		fmt.Printf("%s\n\n", meta.Format())
		if !prompter.YN("Are you sure to install release binary from above repository?", true) {
			return nil
		}
		fmt.Println()
	}
	return s.install(meta, command)
}

// findMetadata find metadata about GitHub release asset.
func (s *Service) findMetadata(ctx context.Context, command Command) (Metadata, error) {
	var err error

	var repo repository.Repository
	if command.hasOwner() {
		repo, err = s.github.FindRepository(ctx, command.Owner(), command.Repo())
	} else {
		repo, err = s.github.SearchRepository(ctx, command.searchRepositoryQuery())
	}
	if err != nil {
		return Metadata{}, err
	}

	var release release.Release
	if command.hasTag() {
		release, err = s.github.FindReleaseByTag(ctx, repo, command.Tag())
	} else {
		release, err = s.github.LatestRelease(ctx, repo)
	}
	if err != nil {
		return Metadata{}, err
	}

	asset, err := s.asset.FindMetadata(ctx, repo, release, command.platform())
	if err != nil {
		return Metadata{}, err
	}

	execBinary, err := s.execbinary.FindMetadata(repo, command.platform())
	if err != nil {
		return Metadata{}, err
	}

	return NewMetadata(repo, release, asset, execBinary), nil
}

// install executable binary from GitHub release asset.
func (s *Service) install(meta Metadata, command Command) error {
	assetContent, err := s.http.DownloadAssetContent(meta.Asset(), command.isInteractive())
	if err != nil {
		return err
	}
	execBinaryContentInAsset, err := assetContent.FindExecBinary(meta.ExecBinary().Name())
	if err != nil {
		return err
	}
	execBinaryContent := execbinary.NewContentFromAsset(execBinaryContentInAsset)
	return s.filesystem.WriteExecBinaryContent(meta.ExecBinary(), execBinaryContent, command.InstallDir())
}
