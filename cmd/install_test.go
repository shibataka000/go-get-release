package cmd

import (
	"os"
	"os/exec"
	"testing"
)

func TestInstall(t *testing.T) {
	tests := []struct {
		name          string
		verifyCommand []string
	}{
		{
			name:          "terraform",
			verifyCommand: []string{"terraform", "version"},
		},
		{
			name:          "istio",
			verifyCommand: []string{"istioctl", "version", "--remote=false"},
		},
		{
			name:          "protocolbuffers/protobuf",
			verifyCommand: []string{"protoc", "--version"},
		},
		{
			name:          "vmware-tanzu/velero",
			verifyCommand: []string{"velero", "--help"},
		},
		{
			name:          "argoproj/argo-workflows",
			verifyCommand: []string{"argo", "version"},
		},
		{
			name:          "buildpacks/pack",
			verifyCommand: []string{"pack", "--version"},
		},
		{
			name:          "koalaman/shellcheck",
			verifyCommand: []string{"shellcheck", "--version"},
		},
	}

	token := os.Getenv("GITHUB_TOKEN")
	installDir, err := os.MkdirTemp("", "go-get-release-")
	if err != nil {
		t.Fatal(err)
	}
	t.Setenv("PATH", installDir)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := exec.Command(tt.verifyCommand[0], tt.verifyCommand[1:]...)
			err := cmd.Run()
			if err == nil {
				t.Fatalf("Binary is already installed")
			}

			err = install(tt.name, token, os.Getenv("GOOS"), os.Getenv("GOARCH"), installDir, false)
			if err != nil {
				t.Fatal(err)
			}

			cmd = exec.Command(tt.verifyCommand[0], tt.verifyCommand[1:]...)
			err = cmd.Run()
			if err != nil {
				t.Fatalf("Installation failed: %v", err)
			}
		})
	}
}
