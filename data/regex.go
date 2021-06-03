package data

import "regexp"

const (
	versionStringRegex        = `(?P<major>\d+)\.(?P<minor>\d+)\.(?P<bugfix>\d+)`
	nonTestVersionStringRegex = `(?P<major>\d+)\.(?P<minor>\d+)\.(?P<bugfix>\d+)$`
)

func getVersionStringRegex() *regexp.Regexp {
	return regexp.MustCompile(versionStringRegex)
}

func getNonTestVersionStringRegex() *regexp.Regexp {
	return regexp.MustCompile(nonTestVersionStringRegex)
}
