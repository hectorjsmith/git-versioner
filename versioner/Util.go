package versioner

import (
	"fmt"
	"gitlab.com/hectorjsmith/git-versioner/data"
	"gitlab.com/hectorjsmith/git-versioner/git"
	"gitlab.com/hectorjsmith/git-versioner/util"
	workerdata "gitlab.com/hectorjsmith/git-versioner/versioner/data"
	"log"
	"sort"
)

type byVersionNumber []string

func (s byVersionNumber) Len() int {
	return len(s)
}

func (s byVersionNumber) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s byVersionNumber) Less(i, j int) bool {
	v1 := data.NewVersionDataFromString(s[i])
	v2 := data.NewVersionDataFromString(s[j])

	return v1.Compare(v2) < 0
}

func GetLatestVersion() workerdata.TagVersionData {
	validTags := getSortedValidTags()
	currentVersion := validTags[0]

	return workerdata.TagVersionData{
		VersionData: data.NewVersionDataFromString(currentVersion),
		TagName:     currentVersion,
	}
}

func GetMatchingVersion(versionToFind data.VersionData) workerdata.TagVersionData {
	validTags := getSortedValidTags()
	version, err := findLatestBugfixForVersion(versionToFind, validTags)
	util.CheckIfError(err)
	return version
}

func getSortedValidTags() []string {
	repo, err := git.GetRepositoryForPath(".")
	util.CheckIfError(err)

	allTags := repo.AllTags()
	validTags := filterToOnlyValidTags(allTags)
	if len(validTags) == 0 {
		log.Fatal("No valid tags found")
	}
	return sortTags(validTags)
}

func filterToOnlyValidTags(tags []string) []string {
	validTags := []string{}
	for i := 0; i < len(tags); i++ {
		tag := tags[i]
		if data.IsValidVersionString(tag) {
			validTags = append(validTags, tag)
		}
	}
	return validTags
}

func sortTags(tags []string) []string {
	sort.Sort(byVersionNumber(tags))
	return tags
}

func findLatestBugfixForVersion(versionToFind data.VersionData, tags []string) (workerdata.TagVersionData, error) {
	var bestMatch workerdata.TagVersionData
	matchFound := false

	for i := 0; i < len(tags); i++ {
		tag := tags[i]
		tagVersion := data.NewVersionDataFromString(tag)
		if tagVersion.Major == versionToFind.Major &&
			tagVersion.Minor == versionToFind.Minor &&
			tagVersion.Bugfix >= versionToFind.Bugfix &&
			tagVersion.Bugfix >= bestMatch.VersionData.Bugfix {

			bestMatch = workerdata.TagVersionData{
				VersionData: tagVersion,
				TagName:     tag,
			}
			matchFound = true
		}
	}
	if !matchFound {
		return bestMatch, fmt.Errorf("no match found")
	}
	return bestMatch, nil
}
