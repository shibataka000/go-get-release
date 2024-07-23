package github

import (
	"fmt"

	dist "github.com/shibataka000/go-get-release/distribution"
)

type ExecBinary struct {
	name string
	os   dist.OS
}

func newExecBinary(name string, os dist.OS) ExecBinary {
	return ExecBinary{
		name: name,
		os:   os,
	}
}

func (b ExecBinary) Name() string {
	if b.os == "windows" {
		return fmt.Sprintf("%s.exe", b.name)
	}
	return b.name
}

type ExecBinaryTemplate struct {
	name string
}

func newExecBinaryTemplate(name string) ExecBinaryTemplate {
	return ExecBinaryTemplate{
		name: name,
	}
}

func (b ExecBinaryTemplate) execute(os dist.OS) ExecBinary {
	return newExecBinary(b.name, os)
}
