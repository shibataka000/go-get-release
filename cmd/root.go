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
		repoFullName string
		tag          string
		patterns     map[string]string
		dir          string
		token        string
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
			asset, execBinary, err := app.Find(ctx, repoFullName, tag, patterns)
			if err != nil {
				return err
			}
			message := fmt.Sprintf("Do you want to install %s from %s ?", execBinary.Name, asset.DownloadURL.String())
			if !prompter.YN(message, true) {
				return nil
			}
			return app.Install(ctx, repoFullName, asset, execBinary, dir, os.Stdout)
		},
		SilenceErrors: true,
		SilenceUsage:  true,
	}

	command.Flags().StringVarP(&repoFullName, "repo", "R", "", "GitHub repository name. This should be OWNER/REPO format.")
	command.Flags().StringVar(&tag, "tag", "", "GitHub release tag.")
	command.Flags().StringToStringVar(&patterns, "pattern", github.DefaultPatterns, "")
	command.Flags().StringVarP(&dir, "dir", "D", ".", "Directory where executable binary will be installed into.")
	command.Flags().StringVar(&token, "token", "", "Authentication token for GitHub API requests.")

	requiredFlags := []string{"repo", "tag"}

	for _, flag := range requiredFlags {
		if err := command.MarkFlagRequired(flag); err != nil {
			panic(err)
		}
	}

	return command
}
