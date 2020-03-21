package github

import "testing"

func TestGetGoosAndGoarchByAsset(t *testing.T) {
	tests := []struct {
		asset  string
		goos   string
		goarch string
	}{
		{
			asset:  "argo-linux-amd64",
			goos:   "linux",
			goarch: "amd64",
		},
		{
			asset:  "argo-windows-amd64",
			goos:   "windows",
			goarch: "amd64",
		},
		{
			asset:  "argo-darwin-amd64",
			goos:   "darwin",
			goarch: "amd64",
		},
		{
			asset:  "docker-compose-Linux-x86_64",
			goos:   "linux",
			goarch: "amd64",
		},
		{
			asset:  "docker-compose-Windows-x86_64.exe",
			goos:   "windows",
			goarch: "amd64",
		},
		{
			asset:  "docker-compose-Darwin-x86_64",
			goos:   "darwin",
			goarch: "amd64",
		},
	}

	for _, tt := range tests {
		t.Run(tt.asset, func(t *testing.T) {
			goos, err := getGoosByAsset(tt.asset)
			if err != nil {
				t.Error(err)
				return
			}
			if goos != tt.goos {
				t.Errorf("Expected is %v but actual is %v", tt.goos, goos)
				return
			}
			goarch, err := getGoarchByAsset(tt.asset)
			if err != nil {
				t.Error(err)
				return
			}
			if goarch != tt.goarch {
				t.Errorf("Expected is %v but actual is %v", tt.goarch, goarch)
				return
			}
		})
	}
}
