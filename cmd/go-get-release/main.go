package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/shibataka000/go-get-release/pkg/cmd"
	"github.com/urfave/cli/v2"
)

var version string

func main() {
	log.SetFlags(0)

	app := &cli.App{
		Name:  "go-get-release",
		Usage: "install golang release binary",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "github-personal-access-token",
				Value:   "",
				EnvVars: []string{"GITHUB_PERSONAL_ACCESS_TOKEN"},
			},
			&cli.StringFlag{
				Name:    "goos",
				Value:   "linux",
				EnvVars: []string{"GOOS"},
			},
			&cli.StringFlag{
				Name:    "goarch",
				Value:   "amd64",
				EnvVars: []string{"GOARCH"},
			},
			&cli.StringFlag{
				Name:  "install-dir",
				Value: filepath.Join(os.Getenv("GOPATH"), "bin"),
			},
		},
		Commands: []*cli.Command{
			{
				Name:  "install",
				Usage: "Install golang release binary",
				Action: func(c *cli.Context) error {
					option := cmd.Option{
						GithubToken: c.String("github-personal-access-token"),
						Goos:        c.String("goos"),
						Goarch:      c.String("goarch"),
						InstallDir:  c.String("install-dir"),
						ShowPrompt:  true,
					}
					if c.Args().Len() == 0 {
						return fmt.Errorf("No repository is specified")
					}
					return cmd.Install(c.Args().Get(0), &option)
				},
			},
			{
				Name:  "search",
				Usage: "Search GitHub repository",
				Action: func(c *cli.Context) error {
					option := cmd.Option{
						GithubToken: c.String("github-personal-access-token"),
						Goos:        c.String("goos"),
						Goarch:      c.String("goarch"),
						InstallDir:  c.String("install-dir"),
						ShowPrompt:  true,
					}
					if c.Args().Len() == 0 {
						return fmt.Errorf("No repository is specified")
					}
					return cmd.Search(c.Args().Get(0), &option)
				},
			},
			{
				Name:  "tags",
				Usage: "Show tags of GitHub repository",
				Action: func(c *cli.Context) error {
					return nil
				},
			},
			{
				Name:  "version",
				Usage: "Show client version",
				Action: func(c *cli.Context) error {
					fmt.Println(version)
					return nil
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
