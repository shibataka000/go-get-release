package cmd

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
			pkgName:       "shibataka000/go-get-release",
			verifyCommand: []string{"go-get-release", "version"},
		},
		{
			pkgName:       "terraform",
			verifyCommand: []string{"terraform", "version"},
		},
		{
			pkgName:       "istio=1.6.0",
			verifyCommand: []string{"istioctl", "version", "--remote=false"},
		},
		{
			pkgName:       "protocolbuffers/protobuf",
			verifyCommand: []string{"protoc", "--version"},
		},
		{
			pkgName:       "vmware-tanzu/velero",
			verifyCommand: []string{"velero", "--help"},
		},
	}

	option := Option{
		GithubToken: os.Getenv("GITHUB_PERSONAL_ACCESS_TOKEN"),
		Goos:        os.Getenv("GOOS"),
		Goarch:      os.Getenv("GOARCH"),
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
