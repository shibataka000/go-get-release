$ GITHUB_TOKEN=$(gh auth token) go test -run ^TestListAssets$ github.com/shibataka000/go-get-release/github
--- FAIL: TestListAssets (23.22s)
    --- FAIL: TestListAssets/kubernetes/kubernetes/v1.30.2 (0.53s)
        asset_test.go:343:
                Error Trace:    /home/takao/src/shibataka000/go-get-release/github/asset_test.go:343
                Error:          Not equal:
                                expected: github.AssetList{github.Asset{DownloadURL:(*url.URL)(0xc0000f01b0)}, github.Asset{DownloadURL:(*url.URL)(0xc0000f0240)}, github.Asset{DownloadURL:(*url.URL)(0xc0000f02d0)}, github.Asset{DownloadURL:(*url.URL)(0xc0000f0360)}, github.Asset{DownloadURL:(*url.URL)(0xc0000f03f0)}}
                                actual  : github.AssetList{}

                                Diff:
                                --- Expected
                                +++ Actual
                                @@ -1,77 +1,2 @@
                                -(github.AssetList) (len=5) {
                                - (github.Asset) {
                                -  DownloadURL: (*url.URL)({
                                -   Scheme: (string) (len=5) "https",
                                -   Opaque: (string) "",
                                -   User: (*url.Userinfo)(<nil>),
                                -   Host: (string) (len=9) "dl.k8s.io",
                                -   Path: (string) (len=41) "/release/v1.30.2/bin/darwin/amd64/kubectl",
                                -   RawPath: (string) "",
                                -   OmitHost: (bool) false,
                                -   ForceQuery: (bool) false,
                                -   RawQuery: (string) "",
                                -   Fragment: (string) "",
                                -   RawFragment: (string) ""
                                -  })
                                - },
                                - (github.Asset) {
                                -  DownloadURL: (*url.URL)({
                                -   Scheme: (string) (len=5) "https",
                                -   Opaque: (string) "",
                                -   User: (*url.Userinfo)(<nil>),
                                -   Host: (string) (len=9) "dl.k8s.io",
                                -   Path: (string) (len=41) "/release/v1.30.2/bin/darwin/arm64/kubectl",
                                -   RawPath: (string) "",
                                -   OmitHost: (bool) false,
                                -   ForceQuery: (bool) false,
                                -   RawQuery: (string) "",
                                -   Fragment: (string) "",
                                -   RawFragment: (string) ""
                                -  })
                                - },
                                - (github.Asset) {
                                -  DownloadURL: (*url.URL)({
                                -   Scheme: (string) (len=5) "https",
                                -   Opaque: (string) "",
                                -   User: (*url.Userinfo)(<nil>),
                                -   Host: (string) (len=9) "dl.k8s.io",
                                -   Path: (string) (len=40) "/release/v1.30.2/bin/linux/amd64/kubectl",
                                -   RawPath: (string) "",
                                -   OmitHost: (bool) false,
                                -   ForceQuery: (bool) false,
                                -   RawQuery: (string) "",
                                -   Fragment: (string) "",
                                -   RawFragment: (string) ""
                                -  })
                                - },
                                - (github.Asset) {
                                -  DownloadURL: (*url.URL)({
                                -   Scheme: (string) (len=5) "https",
                                -   Opaque: (string) "",
                                -   User: (*url.Userinfo)(<nil>),
                                -   Host: (string) (len=9) "dl.k8s.io",
                                -   Path: (string) (len=40) "/release/v1.30.2/bin/linux/arm64/kubectl",
                                -   RawPath: (string) "",
                                -   OmitHost: (bool) false,
                                -   ForceQuery: (bool) false,
                                -   RawQuery: (string) "",
                                -   Fragment: (string) "",
                                -   RawFragment: (string) ""
                                -  })
                                - },
                                - (github.Asset) {
                                -  DownloadURL: (*url.URL)({
                                -   Scheme: (string) (len=5) "https",
                                -   Opaque: (string) "",
                                -   User: (*url.Userinfo)(<nil>),
                                -   Host: (string) (len=9) "dl.k8s.io",
                                -   Path: (string) (len=46) "/release/v1.30.2/bin/windows/amd64/kubectl.exe",
                                -   RawPath: (string) "",
                                -   OmitHost: (bool) false,
                                -   ForceQuery: (bool) false,
                                -   RawQuery: (string) "",
                                -   Fragment: (string) "",
                                -   RawFragment: (string) ""
                                -  })
                                - }
                                +(github.AssetList) {
                                 }
                Test:           TestListAssets/kubernetes/kubernetes/v1.30.2
    --- FAIL: TestListAssets/aquasecurity/trivy/v0.53.0 (1.04s)
        asset_test.go:343:
                Error Trace:    /home/takao/src/shibataka000/go-get-release/github/asset_test.go:343
                Error:          Not equal:
                                expected: github.AssetList{github.Asset{DownloadURL:(*url.URL)(0xc0002661b0)}, github.Asset{DownloadURL:(*url.URL)(0xc0002663f0)}, github.Asset{DownloadURL:(*url.URL)(0xc000266480)}, github.Asset{DownloadURL:(*url.URL)(0xc000266510)}, github.Asset{DownloadURL:(*url.URL)(0xc0002665a0)}, github.Asset{DownloadURL:(*url.URL)(0xc000266630)}, github.Asset{DownloadURL:(*url.URL)(0xc0002666c0)}, github.Asset{DownloadURL:(*url.URL)(0xc000266750)}, github.Asset{DownloadURL:(*url.URL)(0xc0002667e0)}, github.Asset{DownloadURL:(*url.URL)(0xc000266870)}, github.Asset{DownloadURL:(*url.URL)(0xc000266900)}, github.Asset{DownloadURL:(*url.URL)(0xc000266990)}, github.Asset{DownloadURL:(*url.URL)(0xc000266a20)}, github.Asset{DownloadURL:(*url.URL)(0xc000266ab0)}, github.Asset{DownloadURL:(*url.URL)(0xc000266b40)}, github.Asset{DownloadURL:(*url.URL)(0xc000266bd0)}, github.Asset{DownloadURL:(*url.URL)(0xc000266c60)}, github.Asset{DownloadURL:(*url.URL)(0xc000266cf0)}, github.Asset{DownloadURL:(*url.URL)(0xc000266d80)}, github.Asset{DownloadURL:(*url.URL)(0xc000266e10)}, github.Asset{DownloadURL:(*url.URL)(0xc000266ea0)}, github.Asset{DownloadURL:(*url.URL)(0xc000266f30)}, github.Asset{DownloadURL:(*url.URL)(0xc000266fc0)}, github.Asset{DownloadURL:(*url.URL)(0xc000267050)}, github.Asset{DownloadURL:(*url.URL)(0xc0002670e0)}, github.Asset{DownloadURL:(*url.URL)(0xc000267170)}, github.Asset{DownloadURL:(*url.URL)(0xc000267200)}, github.Asset{DownloadURL:(*url.URL)(0xc000267290)}, github.Asset{DownloadURL:(*url.URL)(0xc000267320)}, github.Asset{DownloadURL:(*url.URL)(0xc0002673b0)}, github.Asset{DownloadURL:(*url.URL)(0xc000267440)}, github.Asset{DownloadURL:(*url.URL)(0xc0002674d0)}, github.Asset{DownloadURL:(*url.URL)(0xc000267560)}, github.Asset{DownloadURL:(*url.URL)(0xc0002675f0)}, github.Asset{DownloadURL:(*url.URL)(0xc000267680)}, github.Asset{DownloadURL:(*url.URL)(0xc000267710)}, github.Asset{DownloadURL:(*url.URL)(0xc0002677a0)}, github.Asset{DownloadURL:(*url.URL)(0xc000267830)}, github.Asset{DownloadURL:(*url.URL)(0xc0002678c0)}, github.Asset{DownloadURL:(*url.URL)(0xc000267950)}, github.Asset{DownloadURL:(*url.URL)(0xc0002679e0)}, github.Asset{DownloadURL:(*url.URL)(0xc000267a70)}, github.Asset{DownloadURL:(*url.URL)(0xc000267b00)}, github.Asset{DownloadURL:(*url.URL)(0xc000267b90)}, github.Asset{DownloadURL:(*url.URL)(0xc000267c20)}, github.Asset{DownloadURL:(*url.URL)(0xc000267cb0)}, github.Asset{DownloadURL:(*url.URL)(0xc000267d40)}, github.Asset{DownloadURL:(*url.URL)(0xc000267dd0)}, github.Asset{DownloadURL:(*url.URL)(0xc000267e60)}, github.Asset{DownloadURL:(*url.URL)(0xc000267ef0)}, github.Asset{DownloadURL:(*url.URL)(0xc000270000)}, github.Asset{DownloadURL:(*url.URL)(0xc000270090)}, github.Asset{DownloadURL:(*url.URL)(0xc000270120)}, github.Asset{DownloadURL:(*url.URL)(0xc0002701b0)}, github.Asset{DownloadURL:(*url.URL)(0xc000270240)}, github.Asset{DownloadURL:(*url.URL)(0xc0002702d0)}, github.Asset{DownloadURL:(*url.URL)(0xc000270360)}, github.Asset{DownloadURL:(*url.URL)(0xc0002703f0)}, github.Asset{DownloadURL:(*url.URL)(0xc000270480)}, github.Asset{DownloadURL:(*url.URL)(0xc000270510)}, github.Asset{DownloadURL:(*url.URL)(0xc0002705a0)}, github.Asset{DownloadURL:(*url.URL)(0xc000266240)}, github.Asset{DownloadURL:(*url.URL)(0xc0002662d0)}, github.Asset{DownloadURL:(*url.URL)(0xc000266360)}, github.Asset{DownloadURL:(*url.URL)(0xc000270630)}, github.Asset{DownloadURL:(*url.URL)(0xc0002706c0)}, github.Asset{DownloadURL:(*url.URL)(0xc000270750)}, github.Asset{DownloadURL:(*url.URL)(0xc0002707e0)}, github.Asset{DownloadURL:(*url.URL)(0xc000270870)}, github.Asset{DownloadURL:(*url.URL)(0xc000270900)}, github.Asset{DownloadURL:(*url.URL)(0xc000270990)}, github.Asset{DownloadURL:(*url.URL)(0xc000270a20)}, github.Asset{DownloadURL:(*url.URL)(0xc000270ab0)}, github.Asset{DownloadURL:(*url.URL)(0xc000270b40)}}
                                actual  : github.AssetList{github.Asset{DownloadURL:(*url.URL)(0xc00022c1b0)}, github.Asset{DownloadURL:(*url.URL)(0xc00022c3f0)}, github.Asset{DownloadURL:(*url.URL)(0xc00022c480)}, github.Asset{DownloadURL:(*url.URL)(0xc00022c510)}, github.Asset{DownloadURL:(*url.URL)(0xc00022c5a0)}, github.Asset{DownloadURL:(*url.URL)(0xc00022c630)}, github.Asset{DownloadURL:(*url.URL)(0xc00022c6c0)}, github.Asset{DownloadURL:(*url.URL)(0xc00022c750)}, github.Asset{DownloadURL:(*url.URL)(0xc00022c7e0)}, github.Asset{DownloadURL:(*url.URL)(0xc00022c870)}, github.Asset{DownloadURL:(*url.URL)(0xc00022c900)}, github.Asset{DownloadURL:(*url.URL)(0xc00022c990)}, github.Asset{DownloadURL:(*url.URL)(0xc00022ca20)}, github.Asset{DownloadURL:(*url.URL)(0xc00022cab0)}, github.Asset{DownloadURL:(*url.URL)(0xc00022cb40)}, github.Asset{DownloadURL:(*url.URL)(0xc00022cbd0)}, github.Asset{DownloadURL:(*url.URL)(0xc00022cc60)}, github.Asset{DownloadURL:(*url.URL)(0xc00022ccf0)}, github.Asset{DownloadURL:(*url.URL)(0xc00022cd80)}, github.Asset{DownloadURL:(*url.URL)(0xc00022ce10)}, github.Asset{DownloadURL:(*url.URL)(0xc00022cfc0)}, github.Asset{DownloadURL:(*url.URL)(0xc00022d050)}, github.Asset{DownloadURL:(*url.URL)(0xc00022d0e0)}, github.Asset{DownloadURL:(*url.URL)(0xc00022d170)}, github.Asset{DownloadURL:(*url.URL)(0xc00022d200)}, github.Asset{DownloadURL:(*url.URL)(0xc00022d290)}, github.Asset{DownloadURL:(*url.URL)(0xc00022d320)}, github.Asset{DownloadURL:(*url.URL)(0xc00022d8c0)}, github.Asset{DownloadURL:(*url.URL)(0xc00022d950)}, github.Asset{DownloadURL:(*url.URL)(0xc00022d9e0)}, github.Asset{DownloadURL:(*url.URL)(0xc00022da70)}, github.Asset{DownloadURL:(*url.URL)(0xc00022db00)}, github.Asset{DownloadURL:(*url.URL)(0xc00022db90)}, github.Asset{DownloadURL:(*url.URL)(0xc00022dc20)}, github.Asset{DownloadURL:(*url.URL)(0xc00022dcb0)}, github.Asset{DownloadURL:(*url.URL)(0xc00022dd40)}, github.Asset{DownloadURL:(*url.URL)(0xc00022ddd0)}, github.Asset{DownloadURL:(*url.URL)(0xc00022de60)}, github.Asset{DownloadURL:(*url.URL)(0xc00022def0)}, github.Asset{DownloadURL:(*url.URL)(0xc00038c000)}, github.Asset{DownloadURL:(*url.URL)(0xc00038c090)}, github.Asset{DownloadURL:(*url.URL)(0xc00038c120)}, github.Asset{DownloadURL:(*url.URL)(0xc00038c1b0)}, github.Asset{DownloadURL:(*url.URL)(0xc00038c240)}, github.Asset{DownloadURL:(*url.URL)(0xc00038c2d0)}, github.Asset{DownloadURL:(*url.URL)(0xc00038c360)}, github.Asset{DownloadURL:(*url.URL)(0xc00038c3f0)}, github.Asset{DownloadURL:(*url.URL)(0xc00038c480)}, github.Asset{DownloadURL:(*url.URL)(0xc00038c510)}, github.Asset{DownloadURL:(*url.URL)(0xc00038c5a0)}, github.Asset{DownloadURL:(*url.URL)(0xc00038c630)}, github.Asset{DownloadURL:(*url.URL)(0xc00038c6c0)}, github.Asset{DownloadURL:(*url.URL)(0xc00038c750)}, github.Asset{DownloadURL:(*url.URL)(0xc00038c7e0)}, github.Asset{DownloadURL:(*url.URL)(0xc00038c870)}, github.Asset{DownloadURL:(*url.URL)(0xc00038c900)}, github.Asset{DownloadURL:(*url.URL)(0xc00038c990)}, github.Asset{DownloadURL:(*url.URL)(0xc00038ce10)}, github.Asset{DownloadURL:(*url.URL)(0xc00038cea0)}, github.Asset{DownloadURL:(*url.URL)(0xc00038cf30)}, github.Asset{DownloadURL:(*url.URL)(0xc00038cfc0)}, github.Asset{DownloadURL:(*url.URL)(0xc00022c240)}, github.Asset{DownloadURL:(*url.URL)(0xc00022c2d0)}, github.Asset{DownloadURL:(*url.URL)(0xc00022c360)}, github.Asset{DownloadURL:(*url.URL)(0xc00038d050)}, github.Asset{DownloadURL:(*url.URL)(0xc00038d0e0)}, github.Asset{DownloadURL:(*url.URL)(0xc00038d170)}, github.Asset{DownloadURL:(*url.URL)(0xc00038d200)}, github.Asset{DownloadURL:(*url.URL)(0xc00038d290)}, github.Asset{DownloadURL:(*url.URL)(0xc00038d440)}, github.Asset{DownloadURL:(*url.URL)(0xc00038d4d0)}, github.Asset{DownloadURL:(*url.URL)(0xc00038d560)}, github.Asset{DownloadURL:(*url.URL)(0xc00038d5f0)}}

                                Diff:
                                --- Expected
                                +++ Actual
                                @@ -1,2 +1,2 @@
                                -(github.AssetList) (len=74) {
                                +(github.AssetList) (len=73) {
                                  (github.Asset) {
                                @@ -1088,17 +1088,2 @@
                                    Path: (string) (len=80) "/aquasecurity/trivy/releases/download/v0.53.0/trivy_0.53.0_windows-64bit.zip.sig",
                                -   RawPath: (string) "",
                                -   OmitHost: (bool) false,
                                -   ForceQuery: (bool) false,
                                -   RawQuery: (string) "",
                                -   Fragment: (string) "",
                                -   RawFragment: (string) ""
                                -  })
                                - },
                                - (github.Asset) {
                                -  DownloadURL: (*url.URL)({
                                -   Scheme: (string) (len=5) "https",
                                -   Opaque: (string) "",
                                -   User: (*url.Userinfo)(<nil>),
                                -   Host: (string) (len=10) "github.com",
                                -   Path: (string) (len=67) "/argoproj/argo-cd/releases/download/v2.9.18/argocd-cli.intoto.jsonl",
                                    RawPath: (string) "",
                Test:           TestListAssets/aquasecurity/trivy/v0.53.0
    --- FAIL: TestListAssets/argoproj/argo-cd/v2.9.18 (0.60s)
        asset_test.go:343:
                Error Trace:    /home/takao/src/shibataka000/go-get-release/github/asset_test.go:343
                Error:          Not equal:
                                expected: github.AssetList{github.Asset{DownloadURL:(*url.URL)(0xc0000f75f0)}, github.Asset{DownloadURL:(*url.URL)(0xc0000f7680)}, github.Asset{DownloadURL:(*url.URL)(0xc0000f7710)}, github.Asset{DownloadURL:(*url.URL)(0xc0000f77a0)}, github.Asset{DownloadURL:(*url.URL)(0xc0000f7830)}, github.Asset{DownloadURL:(*url.URL)(0xc000224000)}, github.Asset{DownloadURL:(*url.URL)(0xc000224090)}, github.Asset{DownloadURL:(*url.URL)(0xc000224120)}, github.Asset{DownloadURL:(*url.URL)(0xc0002241b0)}, github.Asset{DownloadURL:(*url.URL)(0xc000224240)}}
                                actual  : github.AssetList{github.Asset{DownloadURL:(*url.URL)(0xc0002242d0)}, github.Asset{DownloadURL:(*url.URL)(0xc000224360)}, github.Asset{DownloadURL:(*url.URL)(0xc0002243f0)}, github.Asset{DownloadURL:(*url.URL)(0xc000224480)}, github.Asset{DownloadURL:(*url.URL)(0xc000224510)}, github.Asset{DownloadURL:(*url.URL)(0xc0002246c0)}, github.Asset{DownloadURL:(*url.URL)(0xc000224750)}, github.Asset{DownloadURL:(*url.URL)(0xc0002247e0)}, github.Asset{DownloadURL:(*url.URL)(0xc000224870)}, github.Asset{DownloadURL:(*url.URL)(0xc000224900)}, github.Asset{DownloadURL:(*url.URL)(0xc000224990)}}

                                Diff:
                                --- Expected
                                +++ Actual
                                @@ -1,2 +1,17 @@
                                -(github.AssetList) (len=10) {
                                +(github.AssetList) (len=11) {
                                + (github.Asset) {
                                +  DownloadURL: (*url.URL)({
                                +   Scheme: (string) (len=5) "https",
                                +   Opaque: (string) "",
                                +   User: (*url.Userinfo)(<nil>),
                                +   Host: (string) (len=10) "github.com",
                                +   Path: (string) (len=67) "/argoproj/argo-cd/releases/download/v2.9.18/argocd-cli.intoto.jsonl",
                                +   RawPath: (string) "",
                                +   OmitHost: (bool) false,
                                +   ForceQuery: (bool) false,
                                +   RawQuery: (string) "",
                                +   Fragment: (string) "",
                                +   RawFragment: (string) ""
                                +  })
                                + },
                                  (github.Asset) {
                Test:           TestListAssets/argoproj/argo-cd/v2.9.18
    --- FAIL: TestListAssets/docker/scan-cli-plugin/v026.0 (0.23s)
        asset_test.go:338:
                Error Trace:    /home/takao/src/shibataka000/go-get-release/github/asset_test.go:338
                Error:          Received unexpected error:
                                GET https://api.github.com/repos/docker/scan-cli-plugin/releases/tags/v026.0: 404 Not Found []
                Test:           TestListAssets/docker/scan-cli-plugin/v026.0
    --- FAIL: TestListAssets/istio/istio/v1.22.2 (0.25s)
        asset_test.go:338:
                Error Trace:    /home/takao/src/shibataka000/go-get-release/github/asset_test.go:338
                Error:          Received unexpected error:
                                GET https://api.github.com/repos/istio/istio/releases/tags/v1.22.2: 404 Not Found []
                Test:           TestListAssets/istio/istio/v1.22.2
FAIL
FAIL    github.com/shibataka000/go-get-release/github   23.228s
FAIL
