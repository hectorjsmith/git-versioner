package rel

import (
	"fmt"
	"gitlab.com/hectorjsmith/git-versioner/data"
	"gitlab.com/hectorjsmith/git-versioner/git"
	"gitlab.com/hectorjsmith/git-versioner/util"
	"gitlab.com/hectorjsmith/git-versioner/versioner"
	"log"
)

type CommandOptions struct {
	Major   bool
	Minor   bool
	Test    bool
	Message string
}

func (opt CommandOptions) isValid() bool {
	return util.CountTrueValues(opt.Major, opt.Minor, opt.Test) > 1
}

func Run(options CommandOptions) error {
	if options.isValid() {
		return fmt.Errorf("Must select at most one option")
	}

	if options.Major {
		return majorRelease(options.Message)
	} else if options.Minor {
		return minorRelease(options.Message)
	} else if options.Test {
		return testRelease(options.Message)
	} else {
		return currentRelease()
	}
}

func currentRelease() error {
	log.Printf("Attempting to create new release version from branch name")
	return updateLatestTagAndTagCurrentCommit(
		func(r *git.Repository) data.VersionData { return data.NewVersionDataFromString(r.CurrentBranch()) },
		func(version data.VersionData) data.VersionData { return version }, "")
}

func majorRelease(message string) error {
	return updateLatestTagAndTagCurrentCommit(
		func(r *git.Repository) data.VersionData { return versioner.GetLatestVersion().VersionData },
		func(version data.VersionData) data.VersionData { return version.IncrementMajor() },
		message)
}

func minorRelease(message string) error {
	return updateLatestTagAndTagCurrentCommit(
		func(r *git.Repository) data.VersionData { return versioner.GetLatestVersion().VersionData },
		func(version data.VersionData) data.VersionData { return version.IncrementMinor() },
		message)
}

func testRelease(message string) error {
	repo, err := git.GetRepositoryForPath(".")
	util.CheckIfError(err)

	tag := repo.HeadCommitTag()
	if tag != "" {
		return fmt.Errorf("Cannot create a release, head is already tagged as release '%s'", tag)
	}

	describe := repo.GitDescribeWithMatchAndExclude("v*.*.*", "*-*")
	log.Printf("Creating new tag '%s'", describe)
	return repo.TagCurrentCommitWithMessage(describe, message)
}

func updateLatestTagAndTagCurrentCommit(
	versionProvider func(r *git.Repository) data.VersionData,
	versionUpdater func(data data.VersionData) data.VersionData,
	message string) error {

	repo, err := git.GetRepositoryForPath(".")
	util.CheckIfError(err)

	tag := repo.HeadCommitTag()
	if tag != "" {
		return fmt.Errorf("Cannot create a release, head is already tagged as release '%s'", tag)
	}

	before := versionProvider(repo)
	after := versionUpdater(before)
	log.Printf("Updating version from '%s' to '%s'", before.VersionString(), after.VersionString())

	tagName := fmt.Sprintf("v%s", after.VersionString())
	log.Printf("Creating new tag '%s'", tagName)
	err = repo.TagCurrentCommitWithMessage(tagName, message)
	if err != nil {
		return fmt.Errorf("Failed to tag current commit: %v", err)
	}
	return nil
}
