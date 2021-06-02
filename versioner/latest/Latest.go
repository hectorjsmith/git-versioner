package latest

import (
	"fmt"
	"gitlab.com/hectorjsmith/git-versioner/git"
	"gitlab.com/hectorjsmith/git-versioner/util"
	"gitlab.com/hectorjsmith/git-versioner/versioner"
	"gitlab.com/hectorjsmith/git-versioner/versioner/data"
	"log"
)

type CommandOptions struct {
	Verbose bool
	Tag     bool
}

func Run(options CommandOptions) error {
	version := versioner.GetLatestVersion()

	if options.Verbose {
		printVerboseInfo(version)
	} else {
		if options.Tag {
			fmt.Println(version.TagName)
		} else {
			fmt.Println(version.VersionData.VersionString())
		}
	}
	return nil
}

func printVerboseInfo(version versionerdata.TagVersionData) {
	log.Printf("Latest version: %s (tag: %s)",
		version.VersionData.VersionString(), version.TagName)

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
