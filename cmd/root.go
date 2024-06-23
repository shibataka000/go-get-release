package cmd

import (
	"context"
	"fmt"
	"os"
	"runtime"

	"github.com/Songmu/prompter"
	"github.com/shibataka000/go-get-release/github"
	"github.com/shibataka000/go-get-release/platform"
	"github.com/spf13/cobra"
)

// NewCommand returns cobra command
func NewCommand() *cobra.Command {
	var (
		token        string
		repoFullName string
		tag          string
		pos          string
		arch         string
	)

	command := &cobra.Command{
		Use:   "gh-release-install",
		Short: "Install executable binary from GitHub release asset.",
		RunE: func(_ *cobra.Command, args []string) error {
			return run(token, repoFullName, tag, platform.OS(pos), platform.Arch(arch))
		},
	}

	command.Flags().StringVar(&token, "token", os.Getenv("GITHUB_TOKEN"), "GitHub token. [$GITHUB_TOKEN]")
	command.Flags().StringVarP(&repoFullName, "repo", "R", "", "Select repository using the OWNER/REPO format")
	command.Flags().StringVar(&tag, "tag", "", "")
	command.Flags().StringVar(&pos, "os", runtime.GOOS, "")
	command.Flags().StringVar(&arch, "arch", runtime.GOARCH, "")

	return command
}

func run(token string, repoFullName string, tag string, os platform.OS, arch platform.Arch) error {
	ctx := context.Background()
	app := github.NewApplicationService(
		github.NewAssetRepository(ctx, token, 0),
	)
	asset, err := app.FindAsset(ctx, repoFullName, tag, os, arch)
	if err != nil {
		return err
	}
	fmt.Println(asset)
	if !prompter.YN("Are you sure to install executable binary from above GitHub release asset?", true) {
		return nil
	}
	return nil
}
