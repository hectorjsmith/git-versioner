package git

import (
	"fmt"
	"gitlab.com/hectorjsmith/git-versioner/util"
	"os/exec"
	"strings"
)

type Repository struct {
	config *config
	Path   string
}

func GetRepositoryForPath(path string) (*Repository, error) {
	repo := &Repository{Path: path, config: defaultGitConfig()}
	if !repo.isValidRepo() {
		return nil, fmt.Errorf("path is not a valid git repository '%s'", path)
	}
	return repo, nil
}

func (r *Repository) isValidRepo() bool {
	_, err := exec.Command(r.config.gitBinaryPath, "status", "--porcelain").Output()
	return err == nil
}

func (r *Repository) RootRepositoryPath() string {
	out, err := exec.Command(r.config.gitBinaryPath, "rev-parse", "--show-toplevel").Output()
	util.CheckIfError(err)
	return strings.TrimSpace(string(out))
}

func (r *Repository) IsClean() bool {
	out, err := exec.Command(r.config.gitBinaryPath, "status", "--porcelain").Output()
	util.CheckIfError(err)

	s := string(out)
	return len(s) == 0
}

func (r *Repository) HeadCommitHash() string {
	out, err := exec.Command(r.config.gitBinaryPath, "rev-parse", "HEAD").Output()
	util.CheckIfError(err)
	return strings.TrimSpace(string(out))
}

func (r *Repository) HeadCommitTag() string {
	out, err := exec.Command(r.config.gitBinaryPath, "tag", "--points-at", "HEAD").Output()
	util.CheckIfError(err)

	return strings.TrimSpace(string(out))
}

func (r *Repository) AllTags() []string {
	out, err := exec.Command(r.config.gitBinaryPath, "tag").Output()
	util.CheckIfError(err)

	tagsRaw := string(out)
	if len(tagsRaw) < 1 {
		return []string{}
	}

	tagsRaw = strings.Replace(tagsRaw, "\r\n", "\n", -1)
	return strings.Split(tagsRaw, "\n")
}

func (r *Repository) TagCurrentCommit(tag string) error {
	_, err := exec.Command(r.config.gitBinaryPath, "tag", tag).Output()
	return err
}

func (r *Repository) TagCurrentCommitWithMessage(tag string, message string) error {
	if message == "" {
		return r.TagCurrentCommit(tag)
	}
	_, err := exec.Command(r.config.gitBinaryPath, "tag", tag, "-m", message).Output()
	return err
}

func (r *Repository) CheckoutTag(tagName string) error {
	cmd := exec.Command(r.config.gitBinaryPath, "checkout", tagName)
	return cmd.Run()
}

func (r *Repository) NewBranch(branchName string) error {
	cmd := exec.Command(r.config.gitBinaryPath, "checkout", "-b", branchName)
	return cmd.Run()
}

func (r *Repository) CurrentBranch() string {
	out, err := exec.Command(r.config.gitBinaryPath, "rev-parse", "--abbrev-ref", "HEAD").Output()
	util.CheckIfError(err)

	return strings.TrimSpace(string(out))
}

func (r *Repository) HardReset() error {
	cmd := exec.Command(r.config.gitBinaryPath, "reset", "--hard")
	return cmd.Run()
}

func (r *Repository) GitDescribe() string {
	return r.GitDescribeWithMatch("*")
}

func (r *Repository) GitDescribeWithMatch(matchString string) string {
	out, err := exec.Command(r.config.gitBinaryPath, "describe", "--tags", "--match", matchString).Output()
	util.CheckIfError(err)

	return strings.TrimSpace(string(out))
}

func (r *Repository) PushTags() error {
	_, err := exec.Command(r.config.gitBinaryPath, "push", "--tags").Output()
	return err
}
