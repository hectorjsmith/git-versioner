package data

import (
	"log"
	"strconv"
)

func NewVersionTagSlice(tags []string, includeTestVersions bool) []VersionTag {
	validTags := filterToOnlyValidTags(tags, includeTestVersions)
	versionTagSlice(validTags).Sort()
	return validTags
}

func filterToOnlyValidTags(tags []string, includeTestVersions bool) []VersionTag {
	validTags := []VersionTag{}
	for i := 0; i < len(tags); i++ {
		tag := tags[i]
		if IsValidVersionString(tag, includeTestVersions) {
			validTag := NewVersionTagFromGitTag(tag)
			validTags = append(validTags, validTag)
		}
	}
	return validTags
}

func NewVersionTagFromGitTag(tag string) VersionTag {
	return VersionTag{
		Version: NewVersionFromString(tag),
		Tag:     tag,
	}
}

func NewVersionFromString(versionString string) Version {
	major, minor, bugfix := parseVersionString(versionString)
	return Version{
		Major:  major,
		Minor:  minor,
		Bugfix: bugfix,
	}
}

func parseVersionString(versionString string) (int, int, int) {
	regex := getVersionStringRegex()
	if !regex.MatchString(versionString) {
		log.Fatalf("failed to parse version string '%s'", versionString)
	}
	m := regex.FindStringSubmatch(versionString)
	n := regex.SubexpNames()
	matchMap := mapSubexpNames(m, n)
	return matchMap["major"], matchMap["minor"], matchMap["bugfix"]
}

//non-mvp: Refactor this to return map[string]string and do conversion after, this can then be re-used
func mapSubexpNames(m, n []string) map[string]int {
	m, n = m[1:], n[1:]
	r := make(map[string]int, len(m))
	var err error
	for i, _ := range n {
		r[n[i]], err = strconv.Atoi(m[i])
		if err != nil {
			log.Fatal("failed to parse regex string into integer", err)
		}
	}
	return r
}
