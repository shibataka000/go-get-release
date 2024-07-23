package github

type ExecBinary struct {
	Name string
}

func newExecBinary(name string) ExecBinary {
	return ExecBinary{
		Name: name,
	}
}
