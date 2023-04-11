package cmd

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/Songmu/prompter"
	"github.com/shibataka000/go-get-release/pkg"
	"github.com/spf13/cobra"
)

// NewCommand return cobra command
func NewCommand() *cobra.Command {
	var (
		token      string
		goos       string
		goarch     string
		installDir string
	)

	command := &cobra.Command{
		Use:   "go-get-release [<owner>/]<repo>[=<tag>]",
		Short: "Install executable binary from GitHub release asset.",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := context.Background()
			client, err := pkg.NewClient(ctx, token)
			if err != nil {
				return err
			}
			query, err := pkg.ParseQuery(args[0])
			if err != nil {
				return err
			}
			platform := pkg.NewPlatform(goos, goarch)
			pkg, err := client.Search(ctx, query, platform)
			if err != nil {
				return err
			}
			fmt.Printf("%s\n\n", pkg.StringToPrompt())
			if !prompter.YN("Are you sure to install executable binary from above GitHub release asset?", true) {
				return nil
			}
			fmt.Println()
			return client.Install(pkg, installDir, os.Stderr)
		},
	}

	command.Flags().StringVar(&token, "token", os.Getenv("GITHUB_TOKEN"), "github token [$GITHUB_TOKEN]")
	command.Flags().StringVar(&goos, "goos", os.Getenv("GOOS"), "goos [$GOOS]")
	command.Flags().StringVar(&goarch, "goarch", os.Getenv("GOARCH"), "goarch [$GOARCH]")
	command.Flags().StringVar(&installDir, "install-dir", filepath.Join(os.Getenv("GOPATH"), "bin"), "directory where executable binary will be installed to")

	return command
}
