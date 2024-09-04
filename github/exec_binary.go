package github

type ExecBinary struct {
	name string
}

type ExecBinaryContent []byte

type ExecBinaryRepository struct{}

func NewExecBinaryRepository() *ExecBinaryRepository {
	return &ExecBinaryRepository{}
}

func (r *ExecBinaryRepository) write(execBinary ExecBinary, content ExecBinaryContent) error {
	return nil
}
