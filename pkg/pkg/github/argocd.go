package github

import "fmt"

func getArgoCDAsset(c *client, repo *repository, release *release, goos, goarch string) (Asset, error) {
	if goos == "windows" {
		return nil, fmt.Errorf("argoproj/argo-cd doesn't support windows")
	}

	a, err := getGeneralAsset(c, repo, release, goos, goarch)
	if err != nil {
		return nil, err
	}

	binaryName := "argocd"

	return &asset{
		name:        a.Name(),
		downloadURL: a.DownloadURL(),
		binaryName:  binaryName,
	}, nil
}
