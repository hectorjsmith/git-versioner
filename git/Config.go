package git

type config struct {
	gitBinaryPath string
}

func defaultGitConfig() *config {
	return &config{gitBinaryPath: "git"}
}
