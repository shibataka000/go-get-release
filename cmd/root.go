package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/Songmu/prompter"
	"github.com/shibataka000/go-get-release/github"
	"github.com/shibataka000/go-get-release/pkg"
	"github.com/spf13/cobra"
)

// NewCommand returns cobra command
func NewCommand() *cobra.Command {
	var (
		repoFullName string
		tag          string
		token        string
		dir          string
	)

	command := &cobra.Command{
		Use:   "go-get-release <owner>/<repo>",
		Short: "Install executable binary from GitHub release asset.",
		RunE: func(_ *cobra.Command, args []string) error {
			ctx := context.Background()
			repository := pkg.NewInfrastructureRepository(ctx, token)
			factory := pkg.NewFactory()
			app := pkg.NewApplicationService(repository, factory)
			platform := pkg.NewPlatform(goos, goarch)
			query, err := pkg.ParseQuery(args[0])
			if err != nil {
				return err
			}
			pkg, err := app.Search(ctx, query, platform)
			if err != nil {
				return err
			}
			fmt.Printf("%s\n\n", pkg.StringToPrompt())
			if !prompter.YN("Are you sure to install executable binary from above GitHub release asset?", true) {
				return nil
			}
			fmt.Println()
			return app.Install(pkg, dir, os.Stderr)
		},
	}

	command.Flags().StringVar(&token, "token", os.Getenv("GITHUB_TOKEN"), "GitHub token. [$GITHUB_TOKEN]")
	command.Flags().StringVar(&dir, "dir", "/usr/local/bin", "The directory to download files into.")

	return command
}

func run(repoFullName string, tag string, token string, dir string) error {
	ctx := context.Background()
	app := newGitHubApplicationService(ctx, token)
	assetMeta, err := app.FindAssetMeta(ctx, repoFullName, tag)
	if err != nil {
		return err
	}
	binMeta, err := app.FindAssetMeta()
}

func newGitHubApplicationService(ctx context.Context, token string) github.ApplicationService {
	return *github.NewApplicationService(
		github.NewAssetRepository(ctx, token),
		github.NewExecutableBinaryRepository(),
	)
}
