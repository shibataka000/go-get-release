package github

// ExecBinary in GitHub release asset.
type ExecBinary struct {
	Meta    ExecBinaryMeta
	Content ExecBinaryContent
}

// ExecBinaryMeta is meta of exec binary.
type ExecBinaryMeta struct {
	Name FileName
}

// ExecBinaryContent is content of exec binary.
type ExecBinaryContent []byte
