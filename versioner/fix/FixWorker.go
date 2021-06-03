package fix

import (
	"fmt"
	"gitlab.com/hectorjsmith/git-versioner/data"
	"gitlab.com/hectorjsmith/git-versioner/git"
	"gitlab.com/hectorjsmith/git-versioner/util"
	"gitlab.com/hectorjsmith/git-versioner/versioner"
	workerdata "gitlab.com/hectorjsmith/git-versioner/versioner/data"
	"log"
)

type CommandOptions struct {
	Version string
}

func Run(options CommandOptions) error {
	if options.Version == "" {
		return current()
	}
	return specificVersion(options.Version)
}

func current() error {
	v := versioner.GetLatestVersion()
	return fixSpecificVersion(v)
}

func specificVersion(version string) error {
	v := versioner.GetMatchingVersion(data.NewVersionDataFromString(version), false)
	if v.TagName == "" {
		return fmt.Errorf("No matching tag found for version '%s'", version)
	}
	return fixSpecificVersion(v)
}

func fixSpecificVersion(version workerdata.TagVersionData) error {
	repo, err := git.GetRepositoryForPath(".")
	util.CheckIfError(err)

	err = repo.CheckoutTag(version.TagName)
	util.CheckIfError(err)

	before := version.VersionData
	after := before.IncrementBugfix()
	log.Printf("New bugfix version '%s' (incremented from: '%s')", after.VersionString(), before.VersionString())

	branchName := "rel/v" + after.VersionString()
	log.Printf("Creating new fix branch '%s' (based on tag: '%s')", branchName, version.TagName)
	return repo.NewBranch(branchName)
}
