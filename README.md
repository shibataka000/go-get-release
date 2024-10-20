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
- [x] error.go
- [x] helper_test.go
- [ ] linux_amd64_test.go
- [ ] linux_amd64.go
- [x] exec_binary.go
- [x] exec_binary_test.go
- [x] pattern.go
- [x] pattern_test.go
- [x] release.go
- [x] repository.go
- [x] repository_test.go
- [x] unarchive.go
- [x] unarchive_test.go
