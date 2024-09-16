# go-get-release

[![Test](https://github.com/shibataka000/go-get-release/actions/workflows/test.yaml/badge.svg)](https://github.com/shibataka000/go-get-release/actions/workflows/test.yaml)

I want to only install executable binary, don't want to build executable binary from source code.

## Usage

### Install executable binary from GitHub release asset

```
go-get-release <owner>/<repo>=<tag>
```

For example

```
go-get-release shibataka000/go-get-release=v0.0.1
```

If you omit tags, `go-get-release` fetch latest tag.

```
go-get-release shibataka000/go-get-release
```

If you omit owner name, `go-get-release` search repository in GitHub.

```
go-get-release go-get-release
```

`go-get-release` find GitHub release asset which should be installed by `$GOOS` and `$GOARCH` automatically.

`go-get-release` install executable binary to `$GOHOME/bin` by default.

## Install

```
go install github.com/shibataka000/go-get-release@master
```

## To Do

- [ ] application.go
- [ ] application_test.go
- [ ] asset.go
- [ ] asset_test.go
- [ ] error.go
- [ ] exec_binary.go
- [ ] exec_binary_test.go
- [ ] pattern.go
- [ ] pattern_test.go
- [x] release.go
- [x] repository.go
- [x] repository_test.go
- [x] unarchive.go
- [x] unarchive_test.go

# Test

- [x] ~~aquasecurity/tfsec~~
- [x] aquasecurity/trivy
- [ ] argoproj/argo-cd
- [ ] argoproj/argo-rollouts
- [ ] argoproj/argo-workflows
- [ ] aws/amazon-ec2-instance-selector
- [ ] bitnami-labs/sealed-secrets
- [ ] buildpacks/pack
- [ ] CircleCI-Public/circleci-cli
- [ ] cli/cli
- [ ] docker/buildx
- [ ] docker/compose
- [x] ~~docker/machine~~
- [x] ~~docker/scan-cli-plugin~~
- [ ] fluxcd/flux2
- [ ] getsops/sops
- [ ] goodwithtech/dockle
- [x] ~~hashicorp/terraform~~
- [x] ~~helm/helm~~
- [x] ~~istio/istio~~
- [x] ~~kubernetes/kubernetes~~
- [ ] mikefarah/yq
- [ ] open-policy-agent/conftest
- [ ] open-policy-agent/gatekeeper
- [ ] open-policy-agent/opa
- [ ] openshift-pipelines/pipelines-as-code
- [ ] protocolbuffers/protobuf
- [ ] snyk/cli
- [ ] starship/starship
- [ ] tektoncd/cli
- [ ] viaduct-ai/kustomize-sops

```
go test ./...
?       github.com/shibataka000/go-get-release  [no test files]
?       github.com/shibataka000/go-get-release/cmd      [no test files]
--- FAIL: TestFindAndInstallOnLinuxAmd64 (146.55s)
    --- FAIL: TestFindAndInstallOnLinuxAmd64/argoproj/argo-cd (35.26s)
        application_test.go:192:
                Error Trace:    /home/takao/src/shibataka000/go-get-release/github/application_test.go:192
                Error:          Received unexpected error:
                                exit status 1
                Test:           TestFindAndInstallOnLinuxAmd64/argoproj/argo-cd
    --- FAIL: TestFindAndInstallOnLinuxAmd64/argoproj/argo-rollouts (15.52s)
        application_test.go:192:
                Error Trace:    /home/takao/src/shibataka000/go-get-release/github/application_test.go:192
                Error:          Received unexpected error:
                                exit status 1
                Test:           TestFindAndInstallOnLinuxAmd64/argoproj/argo-rollouts
    --- FAIL: TestFindAndInstallOnLinuxAmd64/getsops/sops (5.93s)
        application_test.go:192:
                Error Trace:    /home/takao/src/shibataka000/go-get-release/github/application_test.go:192
                Error:          Received unexpected error:
                                exec: "sops": executable file not found in $PATH
                Test:           TestFindAndInstallOnLinuxAmd64/getsops/sops
    --- FAIL: TestFindAndInstallOnLinuxAmd64/goodwithtech/dockle (2.44s)
        application_test.go:192:
                Error Trace:    /home/takao/src/shibataka000/go-get-release/github/application_test.go:192
                Error:          Received unexpected error:
                                exit status 1
                Test:           TestFindAndInstallOnLinuxAmd64/goodwithtech/dockle
    --- FAIL: TestFindAndInstallOnLinuxAmd64/istio/istio (0.23s)
        application_test.go:185:
                Error Trace:    /home/takao/src/shibataka000/go-get-release/github/application_test.go:185
                Error:          Received unexpected error:
                                GET https://api.github.com/repos/istio/istio/releases/tags/v1.22.2: 404 Not Found []
                Test:           TestFindAndInstallOnLinuxAmd64/istio/istio
    --- FAIL: TestFindAndInstallOnLinuxAmd64/protocolbuffers/protobuf (1.69s)
        application_test.go:190:
                Error Trace:    /home/takao/src/shibataka000/go-get-release/github/application_test.go:190
                Error:          Received unexpected error:
                                extracting exec binary content from release asset content was failed: EOF
                Test:           TestFindAndInstallOnLinuxAmd64/protocolbuffers/protobuf
    --- FAIL: TestFindAndInstallOnLinuxAmd64/snyk/cli (0.57s)
        application_test.go:187:
                Error Trace:    /home/takao/src/shibataka000/go-get-release/github/application_test.go:187
                Error:          Not equal:
                                expected: github.ExecBinary{name:"snyk"}
                                actual  : github.ExecBinary{name:"snyc"}

                                Diff:
                                --- Expected
                                +++ Actual
                                @@ -1,3 +1,3 @@
                                 (github.ExecBinary) {
                                - name: (string) (len=4) "snyk"
                                + name: (string) (len=4) "snyc"
                                 }
                Test:           TestFindAndInstallOnLinuxAmd64/snyk/cli
    --- FAIL: TestFindAndInstallOnLinuxAmd64/starship/starship (1.06s)
        application_test.go:187:
                Error Trace:    /home/takao/src/shibataka000/go-get-release/github/application_test.go:187
                Error:          Not equal:
                                expected: github.ExecBinary{name:"starship"}
                                actual  : github.ExecBinary{name:"ksops_4.3.2_Linux_x86_64.tar.gz"}

                                Diff:
                                --- Expected
                                +++ Actual
                                @@ -1,3 +1,3 @@
                                 (github.ExecBinary) {
                                - name: (string) (len=8) "starship"
                                + name: (string) (len=31) "ksops_4.3.2_Linux_x86_64.tar.gz"
                                 }
                Test:           TestFindAndInstallOnLinuxAmd64/starship/starship
    --- FAIL: TestFindAndInstallOnLinuxAmd64/viaduct-ai/kustomize-sops (0.55s)
        application_test.go:185:
                Error Trace:    /home/takao/src/shibataka000/go-get-release/github/application_test.go:185
                Error:          Received unexpected error:
                                no pattern matched with any release asset name
                Test:           TestFindAndInstallOnLinuxAmd64/viaduct-ai/kustomize-sops
FAIL
FAIL    github.com/shibataka000/go-get-release/github   147.247s
FAIL
make: *** [Makefile:17: test] エラー 1
```
