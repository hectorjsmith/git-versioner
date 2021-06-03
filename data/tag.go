package data

import "sort"

type VersionTag struct {
	Version Version
	Tag     string
}

type versionTagSlice []VersionTag

func (s versionTagSlice) Sort() {
	sort.Sort(s)
}

func (s versionTagSlice) Len() int {
	return len(s)
}

func (s versionTagSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s versionTagSlice) Less(i, j int) bool {
	v1 := s[i].Version
	v2 := s[j].Version

	return v1.Compare(v2) < 0
}
