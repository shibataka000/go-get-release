package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

func TestInstall(t *testing.T) {
	tests := []struct {
		name          string
		verifyCommand []string
	}{
		{
			name:          "shibataka000/go-get-release",
			verifyCommand: []string{"go-get-release", "version"},
		},
		{
			name:          "terraform",
			verifyCommand: []string{"terraform", "version"},
		},
		{
			name:          "istio=1.6.0",
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
	}

	token := os.Getenv("GITHUB_PERSONAL_ACCESS_TOKEN")
	dir := filepath.Join(os.Getenv("GOPATH"), "bin")
	pathEnv := fmt.Sprintf("PATH=%s:%s", os.Getenv("PATH"), dir)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := exec.Command(tt.verifyCommand[0], tt.verifyCommand[1:]...)
			cmd.Env = append(os.Environ(), pathEnv)
			err := cmd.Run()
			if err == nil {
				t.Fatalf("Binary is already installed")
			}

			err = install(tt.name, token, os.Getenv("GOOS"), os.Getenv("GOARCH"), dir, false)
			if err != nil {
				t.Fatal(err)
			}

			cmd = exec.Command(tt.verifyCommand[0], tt.verifyCommand[1:]...)
			cmd.Env = append(os.Environ(), pathEnv)
			err = cmd.Run()
			if err != nil {
				t.Fatalf("Installation failed: %v", err)
			}
		})
	}
}
