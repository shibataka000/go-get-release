package cmd

import (
	"context"
	"os"
	"path/filepath"

	"github.com/shibataka000/go-get-release/internal/application"
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
		Short: "Install release binary from GitHub.",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := context.Background()
			appService, err := application.NewService(ctx, token)
			if err != nil {
				return err
			}
			command, err := application.NewCommandFromQuery(args[0], goos, goarch, installDir, true)
			if err != nil {
				return err
			}
			return appService.Install(ctx, command)
		},
	}

	command.Flags().StringVar(&token, "token", os.Getenv("GITHUB_TOKEN"), "github token [$GITHUB_TOKEN]")
	command.Flags().StringVar(&goos, "goos", os.Getenv("GOOS"), "goos [$GOOS]")
	command.Flags().StringVar(&goarch, "goarch", os.Getenv("GOARCH"), "goarch [$GOARCH]")
	command.Flags().StringVar(&installDir, "install-dir", filepath.Join(os.Getenv("GOPATH"), "bin"), "directory where release binary will be installed to")

	return command
}
