# go-get-release

[![CircleCI](https://circleci.com/gh/shibataka000/go-get-release.svg?style=shield)](https://circleci.com/gh/shibataka000/go-get-release)

I want to only install golang release binary, don't want to build golang source code.

## Usage

### Install golang release binary
```
go-get-release install <owner>/<repo>=<tag>
```

For example

```
go-get-release install shibataka000/go-get-release=v0.0.1
```

If you omit tags, `go-get-release` fetch latest tag.

```
go-get-release install shibataka000/go-get-release
```

If you omit owner name, `go-get-release` search repository in GitHub.

```
go-get-release install go-get-release
```

`go-get-release` find asset which should be installed by `$GOOS` and `$GOARCH` automatically.

`go-get-release` install release binary to `$GOHOME/bin` by default.

### Search GitHub repository
```
go-get-release search <keyword>
```

### Show tags of GitHub repository
TBD

## Install
```
go install github.com/shibataka000/go-get-release/cmd/go-get-release@master
```
