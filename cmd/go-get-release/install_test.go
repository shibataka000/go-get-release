package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"testing"
)

func TestInstall(t *testing.T) {
	tests := []struct {
		name          string
		verifyCommand []string
	}{
		{
			name:          "shibataka000/go-get-release=v0.0.1",
			verifyCommand: []string{"go-get-release", "version"},
		},
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

func TestParse(t *testing.T) {
	tests := []struct {
		in  string
		out []string
	}{
		{
			in:  "go-get-release",
			out: []string{"", "go-get-release", ""},
		},
		{
			in:  "shibataka000/go-get-release",
			out: []string{"shibataka000", "go-get-release", ""},
		},
		{
			in:  "go-get-release=1.0.0",
			out: []string{"", "go-get-release", "1.0.0"},
		},
		{
			in:  "shibataka000/go-get-release=1.0.0",
			out: []string{"shibataka000", "go-get-release", "1.0.0"},
		},
		{
			in:  "shibataka000/go-get-release=alpha/1.0.0",
			out: []string{"shibataka000", "go-get-release", "alpha/1.0.0"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.in, func(t *testing.T) {
			owner, repo, version, err := parse(tt.in)
			if err != nil {
				t.Fatal(err)
			}
			actual := []string{owner, repo, version}
			if !reflect.DeepEqual(actual, tt.out) {
				t.Fatalf("Expected is %v but actual is %v", tt.out, actual)
			}
		})
	}
}

func TestHasExt(t *testing.T) {
	tests := []struct {
		description string
		name        string
		exts        []string
		hasExt      bool
	}{
		{
			description: "a.exe",
			name:        "a.exe",
			exts:        []string{".exe"},
			hasExt:      true,
		},
		{
			description: "a.exe_does_not_zip",
			name:        "a.exe",
			exts:        []string{".zip"},
			hasExt:      false,
		},
		{
			description: "a",
			name:        "a",
			exts:        []string{""},
			hasExt:      true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			actual := hasExt(tt.name, tt.exts)
			if actual != tt.hasExt {
				t.Fatalf("Expected is %v but actual is %v", tt.hasExt, actual)
			}
		})
	}
}
