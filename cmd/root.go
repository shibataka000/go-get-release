package cmd

import (
	"os"
	"path/filepath"

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
			return install(args[0], token, goos, goarch, installDir, true)
		},
	}

	command.Flags().StringVar(&token, "token", os.Getenv("GITHUB_TOKEN"), "github token [$GITHUB_TOKEN]")
	command.Flags().StringVar(&goos, "goos", os.Getenv("GOOS"), "goos [$GOOS]")
	command.Flags().StringVar(&goarch, "goarch", os.Getenv("GOARCH"), "goarch [$GOARCH]")
	command.Flags().StringVar(&installDir, "install-dir", filepath.Join(os.Getenv("GOPATH"), "bin"), "directory where release binary will be installed to")

	return command
}
