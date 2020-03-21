package versionerdata

import "gitlab.com/hectorjsmith/git-versioner/data"

type TagVersionData struct {
	VersionData data.VersionData
	TagName     string
}
