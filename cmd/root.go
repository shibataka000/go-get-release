package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/Songmu/prompter"
	"github.com/shibataka000/go-get-release/github"
	"github.com/spf13/cobra"
)

// NewCommand returns cobra command
func NewCommand() *cobra.Command {
	var (
		repoFullName       string
		tag                string
		assetPatterns      []string
		execBinaryPatterns []string
		dir                string
		token              string
	)

	command := &cobra.Command{
		Use:   "gh-release-install",
		Short: "Install executable binary from GitHub release asset.",
		RunE: func(_ *cobra.Command, _ []string) error {
			ctx := context.Background()
			app := github.NewApplicationService(
				github.NewAssetRepository(ctx, token),
				github.NewExecBinaryRepository(),
			)
			asset, execBinary, err := app.Find(ctx, repoFullName, tag, assetPatterns, execBinaryPatterns)
			if err != nil {
				return err
			}
			fmt.Println(asset)
			if !prompter.YN("Are you sure to install executable binary from above GitHub release asset?", true) {
				return nil
			}
			return app.Install(ctx, repoFullName, asset, execBinary, dir, os.Stdout)
		},
	}

	command.Flags().StringVarP(&repoFullName, "repo", "R", "", "GitHub repository name using the OWNER/REPO format")
	command.Flags().StringVar(&tag, "tag", "", "")
	command.Flags().StringArrayVar(&assetPatterns, "asset", []string{}, "")
	command.Flags().StringArrayVar(&execBinaryPatterns, "exec-binary", []string{}, "")
	command.Flags().StringVarP(&dir, "dir", "D", ".", "Directory to install executable binary into")
	command.Flags().StringVar(&token, "token", "", "Authentication token for GitHub API requests.")

	requiredFlags := []string{"repo", "tag", "asset", "exec-binary"}

	for _, flag := range requiredFlags {
		if err := command.MarkFlagRequired(flag); err != nil {
			panic(err)
		}
	}

	return command
}
