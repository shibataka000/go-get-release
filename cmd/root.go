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
		assets       []string
		execBinart   []string
		token        string
	)

	command := &cobra.Command{
		Use:   "gh-release-install",
		Short: "Install executable binary from GitHub release asset.",
		RunE: func(_ *cobra.Command, _ []string) error {
			ctx := context.Background()
			app := github.NewApplicationService(
				github.NewAssetRepository(ctx, token),
			)
			asset, err := app.FindAsset(ctx, repoFullName, tag, assets, execBinart)
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

	command.Flags().StringVarP(&repoFullName, "repo", "R", "", "Select GitHub repository using the OWNER/REPO format")
	command.Flags().StringVar(&tag, "tag", "", "")
	command.Flags().StringArrayVar(&assets, "asset", []string{}, "")
	command.Flags().StringArrayVar(&execBinart, "exec-binary", []string{}, "")
	command.Flags().StringVar(&token, "token", os.Getenv("GH_TOKEN"), "GitHub token. [$GH_TOKEN]")
	command.MarkFlagRequired("repo")        //nolint:errcheck
	command.MarkFlagRequired("tag")         //nolint:errcheck
	command.MarkFlagRequired("asset")       //nolint:errcheck
	command.MarkFlagRequired("exec-binary") //nolint:errcheck

	return command
}
