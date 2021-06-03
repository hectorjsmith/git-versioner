package latest

import (
	"fmt"
	"gitlab.com/hectorjsmith/git-versioner/data"
	"gitlab.com/hectorjsmith/git-versioner/git"
	"gitlab.com/hectorjsmith/git-versioner/util"
	"log"
)

func Run(options CommandOptions) error {
	version := git.GetLatestVersion()

	if options.Verbose {
		printVerboseInfo(version)
	} else {
		if options.Tag {
			fmt.Println(version.Tag)
		} else {
			fmt.Println(version.Version.String())
		}
	}
	return nil
}

func printVerboseInfo(version data.VersionTag) {
	log.Printf("latest version: %s (tag: %s)",
		version.Version.String(), version.Tag)

	repo, err := git.NewRepository(".")
	util.CheckIfError(err)

	extraInfoPrefix := "    "
	log.Printf("%sHEAD commit hash: %s", extraInfoPrefix, repo.HeadCommitHash())
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
