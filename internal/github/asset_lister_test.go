package github

import (
	"os"
	"reflect"
	"testing"
)

func TestAssets(t *testing.T) {
	tests := []struct {
		description  string
		owner        string
		repo         string
		tag          string
		downloadURLs []string
	}{
		{
			description: "shibataka000/go-get-release-test(v0.0.2)",
			owner:       "shibataka000",
			repo:        "go-get-release-test",
			tag:         "v0.0.2",
			downloadURLs: []string{
				"https://github.com/shibataka000/go-get-release-test/releases/download/v0.0.2/go-get-release_v0.0.2_darwin_amd64",
				"https://github.com/shibataka000/go-get-release-test/releases/download/v0.0.2/go-get-release_v0.0.2_linux_amd64",
				"https://github.com/shibataka000/go-get-release-test/releases/download/v0.0.2/go-get-release_v0.0.2_windows_amd64.exe",
			},
		},
		{
			description: "hashicorp/terraform(v0.12.20)",
			owner:       "hashicorp",
			repo:        "terraform",
			tag:         "v0.12.20",
			downloadURLs: []string{
				"https://releases.hashicorp.com/terraform/0.12.20/terraform_0.12.20_SHA256SUMS",
				"https://releases.hashicorp.com/terraform/0.12.20/terraform_0.12.20_SHA256SUMS.348FFC4C.sig",
				"https://releases.hashicorp.com/terraform/0.12.20/terraform_0.12.20_SHA256SUMS.72D7468F.sig",
				"https://releases.hashicorp.com/terraform/0.12.20/terraform_0.12.20_SHA256SUMS.sig",
				"https://releases.hashicorp.com/terraform/0.12.20/terraform_0.12.20_darwin_amd64.zip",
				"https://releases.hashicorp.com/terraform/0.12.20/terraform_0.12.20_freebsd_386.zip",
				"https://releases.hashicorp.com/terraform/0.12.20/terraform_0.12.20_freebsd_amd64.zip",
				"https://releases.hashicorp.com/terraform/0.12.20/terraform_0.12.20_freebsd_arm.zip",
				"https://releases.hashicorp.com/terraform/0.12.20/terraform_0.12.20_linux_386.zip",
				"https://releases.hashicorp.com/terraform/0.12.20/terraform_0.12.20_linux_amd64.zip",
				"https://releases.hashicorp.com/terraform/0.12.20/terraform_0.12.20_linux_arm.zip",
				"https://releases.hashicorp.com/terraform/0.12.20/terraform_0.12.20_openbsd_386.zip",
				"https://releases.hashicorp.com/terraform/0.12.20/terraform_0.12.20_openbsd_amd64.zip",
				"https://releases.hashicorp.com/terraform/0.12.20/terraform_0.12.20_solaris_amd64.zip",
				"https://releases.hashicorp.com/terraform/0.12.20/terraform_0.12.20_windows_386.zip",
				"https://releases.hashicorp.com/terraform/0.12.20/terraform_0.12.20_windows_amd64.zip",
			},
		},
		{
			description: "hashicorp/terraform(v0.12.20)",
			owner:       "helm",
			repo:        "helm",
			tag:         "v3.1.0",
			downloadURLs: []string{
				"https://get.helm.sh/helm-v3.1.0-darwin-amd64.tar.gz",
				"https://get.helm.sh/helm-v3.1.0-linux-386.tar.gz",
				"https://get.helm.sh/helm-v3.1.0-linux-amd64.tar.gz",
				"https://get.helm.sh/helm-v3.1.0-linux-arm.tar.gz",
				"https://get.helm.sh/helm-v3.1.0-linux-arm64.tar.gz",
				"https://get.helm.sh/helm-v3.1.0-linux-ppc64le.tar.gz",
				"https://get.helm.sh/helm-v3.1.0-linux-s390x.tar.gz",
				"https://get.helm.sh/helm-v3.1.0-windows-amd64.zip",
			},
		},
		{
			description: "gravitational/teleport(v8.3.4)",
			owner:       "gravitational",
			repo:        "teleport",
			tag:         "v8.3.4",
			downloadURLs: []string{
				"https://get.gravitational.com/teleport-v8.3.4-linux-amd64-centos7-bin.tar.gz",
				"https://get.gravitational.com/teleport-v8.3.4-linux-arm-bin.tar.gz",
				"https://get.gravitational.com/teleport-v8.3.4-linux-arm64-bin.tar.gz",
				"https://get.gravitational.com/teleport-v8.3.4-windows-amd64-bin.zip",
				"https://get.gravitational.com/teleport-8.3.4-1.arm.rpm",
				"https://get.gravitational.com/teleport-8.3.4-1.arm64.rpm",
				"https://get.gravitational.com/teleport-8.3.4-1.i386.rpm",
				"https://get.gravitational.com/teleport-8.3.4-1.x86_64.rpm",
				"https://get.gravitational.com/teleport-8.3.4.pkg",
				"https://get.gravitational.com/teleport-v8.3.4-darwin-amd64-bin.tar.gz",
				"https://get.gravitational.com/teleport-v8.3.4-linux-386-bin.tar.gz",
				"https://get.gravitational.com/teleport-v8.3.4-linux-amd64-bin.tar.gz",
				"https://get.gravitational.com/teleport-v8.3.4-linux-amd64-centos6-bin.tar.gz",
				"https://get.gravitational.com/teleport_8.3.4_arm.deb",
				"https://get.gravitational.com/teleport_8.3.4_i386.deb",
				"https://get.gravitational.com/teleport_8.3.4_amd64.deb",
				"https://get.gravitational.com/teleport_8.3.4_arm64.deb",
				"https://get.gravitational.com/tsh-8.3.4.pkg",
			},
		},
	}

	token := os.Getenv("GITHUB_PERSONAL_ACCESS_TOKEN")
	c, err := NewClient(token)
	if err != nil {
		t.Fatal(err)
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			repo, err := c.Repository(tt.owner, tt.repo)
			if err != nil {
				t.Fatal(err)
			}
			release, err := repo.Release(tt.tag)
			if err != nil {
				t.Fatal(err)
			}
			assets, err := release.Assets()
			if err != nil {
				t.Fatal(err)
			}

			downloadURLs := []string{}
			for _, a := range assets {
				downloadURLs = append(downloadURLs, a.DownloadURL())
			}

			if !reflect.DeepEqual(tt.downloadURLs, downloadURLs) {
				t.Fatalf("Expected is %v but actual is %v", tt.downloadURLs, downloadURLs)
			}
		})
	}
}
