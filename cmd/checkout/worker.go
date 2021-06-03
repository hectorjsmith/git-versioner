package checkout

import (
	"gitlab.com/hectorjsmith/git-versioner/data"
	"gitlab.com/hectorjsmith/git-versioner/git"
	"log"
)

func Run(options CommandOptions) error {
	var versionTag data.VersionTag
	if options.Latest || options.Version == "" {
		versionTag = git.GetLatestVersion()
	} else {
		versionTag = getVersionTagFromOptions(options)
	}

	repo, err := git.NewRepository(".")
	if err != nil {
		return err
	}
	return repo.CheckoutTag(versionTag.Tag)
}

func getVersionTagFromOptions(options CommandOptions) data.VersionTag {
	if !data.IsValidVersionString(options.Version, false) {
		log.Fatalf("invalid version string: %s", options.Version)
	}

	version := data.NewVersionFromString(options.Version)
	return git.FindTagForVersion(version, false)
}
