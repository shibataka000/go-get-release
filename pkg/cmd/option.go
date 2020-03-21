package cmd

// Option to execute commands
type Option struct {
	GithubToken string
	Goos        string
	Goarch      string
	InstallDir  string
	ShowPrompt  bool
}
