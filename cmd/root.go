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
		readlimit    uint32
	)

	command := &cobra.Command{
		Use:   "gh-release-install",
		Short: "Install executable binary from GitHub release asset.",
		RunE: func(_ *cobra.Command, _ []string) error {
			ctx := context.Background()
			app := github.NewApplicationService(
				github.NewAssetRepository(ctx, token, readlimit),
			)
			asset, err := app.FindAsset(ctx, repoFullName, tag, platform.OS(pos), platform.Arch(arch))
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
	command.Flags().StringVar(&pos, "os", runtime.GOOS, "")
	command.Flags().StringVar(&arch, "arch", runtime.GOARCH, "")
	command.Flags().Uint32Var(&readlimit, "readlimit", 3072, "The maximum number of bytes read from the input used when detecting MIME.")

	return command
}
