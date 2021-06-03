package fix

import (
	"fmt"
	"gitlab.com/hectorjsmith/git-versioner/data"
	"gitlab.com/hectorjsmith/git-versioner/git"
	"gitlab.com/hectorjsmith/git-versioner/util"
	"log"
)

func Run(options CommandOptions) error {
	if options.Version == "" {
		return current()
	}
	return specificVersion(options.Version)
}

func current() error {
	v := git.GetLatestVersion()
	return fixSpecificVersion(v)
}

func specificVersion(version string) error {
	v := git.GetMatchingVersion(data.NewVersionFromString(version), false)
	if v.Tag == "" {
		return fmt.Errorf("no matching tag found for version '%s'", version)
	}
	return fixSpecificVersion(v)
}

func fixSpecificVersion(version data.VersionTag) error {
	repo, err := git.NewRepository(".")
	util.CheckIfError(err)

	err = repo.CheckoutTag(version.Tag)
	util.CheckIfError(err)

	before := version.Version
	after := before.IncrementBugfix()
	log.Printf("new bugfix version '%s' (incremented from: '%s')", after.String(), before.String())

	branchName := "rel/v" + after.String()
	log.Printf("creating new fix branch '%s' (based on tag: '%s')", branchName, version.Tag)
	return repo.NewBranch(branchName)
}
