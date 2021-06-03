package rel

import (
	"fmt"
	"gitlab.com/hectorjsmith/git-versioner/data"
	"gitlab.com/hectorjsmith/git-versioner/git"
	"gitlab.com/hectorjsmith/git-versioner/util"
	"log"
)

func Run(options CommandOptions) error {
	if options.isValid() {
		return fmt.Errorf("must select at most one option")
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
	log.Printf("attempting to create new release version from branch name")
	return updateLatestTagAndTagCurrentCommit(
		func(r *git.Repository) data.Version { return data.NewVersionFromString(r.CurrentBranch()) },
		func(version data.Version) data.Version { return version }, "")
}

func majorRelease(message string) error {
	return updateLatestTagAndTagCurrentCommit(
		func(r *git.Repository) data.Version { return git.GetLatestVersion().Version },
		func(version data.Version) data.Version { return version.IncrementMajor() },
		message)
}

func minorRelease(message string) error {
	return updateLatestTagAndTagCurrentCommit(
		func(r *git.Repository) data.Version { return git.GetLatestVersion().Version },
		func(version data.Version) data.Version { return version.IncrementMinor() },
		message)
}

func testRelease(message string) error {
	repo, err := git.NewRepository(".")
	util.CheckIfError(err)

	tag := repo.HeadCommitTag()
	if tag != "" {
		return fmt.Errorf("cannot create a release, head is already tagged as release '%s'", tag)
	}

	describe := repo.GitDescribeWithMatchAndExclude("v*.*.*", "*-*")
	log.Printf("creating new tag '%s'", describe)
	return repo.TagCurrentCommitWithMessage(describe, message)
}

func updateLatestTagAndTagCurrentCommit(
	versionProvider func(r *git.Repository) data.Version,
	versionUpdater func(data data.Version) data.Version,
	message string) error {

	repo, err := git.NewRepository(".")
	util.CheckIfError(err)

	tag := repo.HeadCommitTag()
	if tag != "" {
		return fmt.Errorf("cannot create a release, head is already tagged as release '%s'", tag)
	}

	before := versionProvider(repo)
	after := versionUpdater(before)
	log.Printf("updating version from '%s' to '%s'", before.String(), after.String())

	tagName := fmt.Sprintf("v%s", after.String())
	log.Printf("creating new tag '%s'", tagName)
	err = repo.TagCurrentCommitWithMessage(tagName, message)
	if err != nil {
		return fmt.Errorf("failed to tag current commit: %v", err)
	}
	return nil
}
