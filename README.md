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
