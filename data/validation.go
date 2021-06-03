package data

import "regexp"

func IsValidVersionString(versionString string, includeTestVersions bool) bool {
	var regex *regexp.Regexp
	if includeTestVersions {
		regex = getVersionStringRegex()
	} else {
		regex = getNonTestVersionStringRegex()
	}
	return regex.MatchString(versionString)
}
