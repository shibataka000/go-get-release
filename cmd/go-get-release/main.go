package main

import (
	"fmt"
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
				return fmt.Errorf("no repository is specified")
			}
			return install(c.Args().Get(0), c.String("github-token"), c.String("goos"), c.String("goarch"), c.String("install-dir"), true)
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
