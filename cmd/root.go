package cmd

import (
	"context"
	"fmt"
	"os"
	"runtime"

	"github.com/Songmu/prompter"
	dist "github.com/shibataka000/go-get-release/distribution"
	"github.com/shibataka000/go-get-release/github"
	"github.com/spf13/cobra"
)

// NewCommand returns cobra command
func NewCommand() *cobra.Command {
	var (
		token        string
		repoFullName string
		tag          string
		goos         string
		goarch       string
	)

	command := &cobra.Command{
		Use:   "gh-release-install",
		Short: "Install executable binary from GitHub release asset.",
		RunE: func(_ *cobra.Command, _ []string) error {
			ctx := context.Background()
			app := github.NewApplicationService(
				github.NewAssetRepository(ctx, token),
			)
			asset, err := app.FindAsset(ctx, repoFullName, tag, dist.OS(goos), dist.Arch(goarch))
			if err != nil {
				return err
			}
			fmt.Println(asset)
			if !prompter.YN("Are you sure to install executable binary from above GitHub release asset?", true) {
				return nil
			}
			return nil
		},
	}

	command.Flags().StringVar(&token, "token", os.Getenv("GITHUB_TOKEN"), "GitHub token. [$GITHUB_TOKEN]")
	command.Flags().StringVarP(&repoFullName, "repo", "R", "", "Select repository using the OWNER/REPO format")
	command.Flags().StringVar(&tag, "tag", "", "")
	command.Flags().StringVar(&goos, "os", runtime.GOOS, "")
	command.Flags().StringVar(&goarch, "arch", runtime.GOARCH, "")
	command.MarkFlagRequired("repo") //nolint:errcheck
	command.MarkFlagRequired("tag")  //nolint:errcheck

	return command
}
