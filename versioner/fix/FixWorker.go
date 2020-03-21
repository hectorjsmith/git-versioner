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

func Run(version string) error {
	if version == "" {
		return current()
	}
	return specificVersion(version)
}

func current() error {
	v := versioner.GetLatestVersion()
	return fixSpecificVersion(v)
}

func specificVersion(version string) error {
	v := versioner.GetMatchingVersion(data.NewVersionDataFromString(version))
	if v.TagName == "" {
		return fmt.Errorf("no matching tag found for version '%s'", version)
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
	log.Printf("New bugfix version '%s' (based on '%s')", after.VersionString(), before.VersionString())

	branchName := "rel/v" + after.VersionString()
	log.Printf("Creating new fix branch '%s'", branchName)
	return repo.NewBranch(branchName)
}
