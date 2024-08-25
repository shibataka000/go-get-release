package cmd

import (
	"context"
	"fmt"
	"os"
	"regexp"

	"github.com/Songmu/prompter"
	"github.com/shibataka000/go-get-release/github"
	"github.com/spf13/cobra"
)

// NewCommand returns cobra command
func NewCommand() *cobra.Command {
	var (
		repoFullName string
		tag          string
		rawPatterns  []string
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
			patterns, err := compilePatterns(rawPatterns)
			if err != nil {
				return err
			}
			asset, err := app.FindAsset(ctx, repoFullName, tag, patterns)
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
	command.Flags().StringArrayVarP(&rawPatterns, "pattern", "p", []string{".*"}, "Select GitHub release asset using a regexp pattern")
	command.Flags().StringVar(&token, "token", os.Getenv("GH_TOKEN"), "GitHub token. [$GH_TOKEN]")
	command.MarkFlagRequired("repo") //nolint:errcheck
	command.MarkFlagRequired("tag")  //nolint:errcheck

	return command
}

func compilePatterns(exprs []string) ([]*regexp.Regexp, error) {
	res := []*regexp.Regexp{}
	for _, expr := range exprs {
		re, err := regexp.Compile(expr)
		if err != nil {
			return nil, err
		}
		res = append(res, re)
	}
	return res, nil
}
