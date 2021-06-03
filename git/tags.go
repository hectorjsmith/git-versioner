package git

import (
	"fmt"
	"gitlab.com/hectorjsmith/git-versioner/data"
	"gitlab.com/hectorjsmith/git-versioner/util"
	"log"
)

func GetLatestVersion() data.VersionTag {
	validTags := GetSortedValidTags(false)
	currentVersion := validTags[0]

	return currentVersion
}

func GetMatchingVersion(versionToFind data.Version, includeTestVersions bool) data.VersionTag {
	validTags := GetSortedValidTags(includeTestVersions)
	version, err := findLatestBugfixForVersion(versionToFind, validTags)
	util.CheckIfError(err)
	return version
}

func GetSortedValidTags(includeTestVersions bool) []data.VersionTag {
	repo, err := NewRepository(".")
	util.CheckIfError(err)

	allTags := repo.AllTags()
	validTags := data.NewVersionTagSlice(allTags, includeTestVersions)
	if len(validTags) == 0 {
		log.Fatal("no valid tags found")
	}
	return validTags
}

func findLatestBugfixForVersion(versionToFind data.Version, tags []data.VersionTag) (data.VersionTag, error) {
	var bestMatch data.VersionTag
	matchFound := false

	for i := 0; i < len(tags); i++ {
		tag := tags[i]
		if tag.Version.Major == versionToFind.Major &&
			tag.Version.Minor == versionToFind.Minor &&
			tag.Version.Bugfix >= versionToFind.Bugfix &&
			tag.Version.Bugfix >= bestMatch.Version.Bugfix {

			bestMatch = tag
			matchFound = true
		}
	}
	if !matchFound {
		return bestMatch, fmt.Errorf("no match found")
	}
	return bestMatch, nil
}
