package latest

import (
	"gitlab.com/hectorjsmith/git-versioner/git"
	"gitlab.com/hectorjsmith/git-versioner/util"
	"gitlab.com/hectorjsmith/git-versioner/versioner"
	"log"
)

func Run(verbose bool) error {
	version := versioner.GetLatestVersion()
	log.Printf("Latest version: %s (tag: %s)",
		version.VersionData.VersionString(), version.TagName)

	if verbose {
		printVerboseInfo()
	}
	return nil
}

func printVerboseInfo() {
	extraInfoPrefix := "    "
	repo, err := git.GetRepositoryForPath(".")
	log.Printf("%sHEAD commit hash: %s", extraInfoPrefix, repo.HeadCommitHash())

	util.CheckIfError(err)
	if headTag := repo.HeadCommitTag(); headTag != "" {
		log.Printf("%sHEAD tag: %s", extraInfoPrefix, headTag)
	} else {
		log.Printf("%sHEAD is not tagged", extraInfoPrefix)
	}
	if repo.IsClean() {
		log.Printf("%sRepo is clean", extraInfoPrefix)
	} else {
		log.Printf("%sRepo is DIRTY", extraInfoPrefix)
	}
}
