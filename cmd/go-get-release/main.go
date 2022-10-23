package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/urfave/cli/v2"
)

func main() {
	log.SetFlags(0)

	app := &cli.App{
		Name:  "go-get-release",
		Usage: "install golang release binary",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "github-token",
				Value:   "",
				EnvVars: []string{"GITHUB_TOKEN"},
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
		Action: func(c *cli.Context) error {
			if c.Args().Len() == 0 {
				return cli.ShowAppHelp(c)
			}

			name := c.Args().Get(0)
			token := c.String("github-token")
			goos := c.String("goos")
			goarch := c.String("goarch")
			installDir := c.String("install-dir")
			showPrompt := true

			return install(name, token, goos, goarch, installDir, showPrompt)
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
