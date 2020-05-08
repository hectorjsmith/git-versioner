package data

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
)

type VersionData struct {
	Major  int
	Minor  int
	Bugfix int
}

func (v VersionData) VersionString() string {
	return fmt.Sprintf("%d.%d.%d", v.Major, v.Minor, v.Bugfix)
}

func (v VersionData) IncrementMajor() VersionData {
	return newVersionData(v.Major+1, 0, 0)
}

func (v VersionData) IncrementMinor() VersionData {
	return newVersionData(v.Major, v.Minor+1, 0)
}

func (v VersionData) IncrementBugfix() VersionData {
	return newVersionData(v.Major, v.Minor, v.Bugfix+1)
}

func (v VersionData) Compare(other VersionData) int {
	if v.Major != other.Major {
		return other.Major - v.Major
	}
	if v.Minor != other.Minor {
		return other.Minor - v.Minor
	}
	return other.Bugfix - v.Bugfix
}

func IsValidVersionString(versionString string, includeTestVersions bool) bool {
	var regex *regexp.Regexp
	if includeTestVersions {
		regex = getVersionStringRegex()
	} else {
		regex = getNonTestVersionStringRegex()
	}
	return regex.MatchString(versionString)
}

func NewVersionDataFromString(versionString string) VersionData {
	major, minor, bugfix := parseVersionString(versionString)
	return newVersionData(major, minor, bugfix)
}

func newVersionData(major int, minor int, bugfix int) VersionData {
	return VersionData{
		Major:  major,
		Minor:  minor,
		Bugfix: bugfix,
	}
}

func parseVersionString(versionString string) (int, int, int) {
	regex := getVersionStringRegex()
	if !regex.MatchString(versionString) {
		log.Fatalf("Failed to parse version string '%s'", versionString)
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
			log.Fatal("Failed to parse regex string into integer", err)
		}
	}
	return r
}

func getVersionStringRegex() *regexp.Regexp {
	return regexp.MustCompile(`(?P<major>\d+)\.(?P<minor>\d+)\.(?P<bugfix>\d+)`)
}

func getNonTestVersionStringRegex() *regexp.Regexp {
	return regexp.MustCompile(`(?P<major>\d+)\.(?P<minor>\d+)\.(?P<bugfix>\d+)$`)
}
