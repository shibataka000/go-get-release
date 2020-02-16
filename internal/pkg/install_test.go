package pkg

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

func TestInstall(t *testing.T) {
	tests := []struct {
		pkgName       string
		verifyCommand []string
	}{
		{
			pkgName:       "argo",
			verifyCommand: []string{"argo", "version"},
		},
		// {
		// 	pkgName:       "docker/compose",
		// 	verifyCommand: []string{"docker-compose", "version"},
		// },
		{
			pkgName:       "docker/machine",
			verifyCommand: []string{"docker-machine", "version"},
		},
		{
			pkgName:       "helm=v3.1.0",
			verifyCommand: []string{"helm", "version"},
		},
		{
			pkgName:       "istio",
			verifyCommand: []string{"istioctl", "version", "--remote=false"},
		},
		{
			pkgName:       "kubectl-bindrole",
			verifyCommand: []string{"kubectl-bindrole", "--version"},
		},
		{
			pkgName:       "kustomize",
			verifyCommand: []string{"kustomize", "version"},
		},
		{
			pkgName:       "stern",
			verifyCommand: []string{"stern", "--version"},
		},
		{
			pkgName:       "terraform",
			verifyCommand: []string{"terraform", "version"},
		},
	}

	option := Option{
		GithubToken: os.Getenv("GITHUB_PERSONAL_ACCESS_TOKEN"),
		OS:          os.Getenv("GOOS"),
		Arch:        os.Getenv("GOARCH"),
		InstallDir:  filepath.Join(os.Getenv("GOPATH"), "bin"),
		ShowPrompt:  false,
	}
	pathEnv := fmt.Sprintf("PATH=%s:%s", os.Getenv("PATH"), option.InstallDir)

	for _, tt := range tests {
		t.Run(tt.pkgName, func(t *testing.T) {
			cmd := exec.Command(tt.verifyCommand[0], tt.verifyCommand[1:]...)
			cmd.Env = append(os.Environ(), pathEnv)
			err := cmd.Run()
			if err == nil {
				t.Errorf("Binary is already installed")
				return
			}

			err = Install(tt.pkgName, &option)
			if err != nil {
				t.Error(err)
				return
			}
			// defer os.Remove(filepath.Join(option.InstallDir, tt.verifyCommand[0]))

			cmd = exec.Command(tt.verifyCommand[0], tt.verifyCommand[1:]...)
			cmd.Env = append(os.Environ(), pathEnv)
			err = cmd.Run()
			if err != nil {
				t.Errorf("Installation failed: %v", err)
				return
			}
		})
	}
}
