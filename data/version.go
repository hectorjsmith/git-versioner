package data

import "fmt"

type Version struct {
	Major  int
	Minor  int
	Bugfix int
}

func (v Version) String() string {
	return fmt.Sprintf("%d.%d.%d", v.Major, v.Minor, v.Bugfix)
}

func (v Version) IncrementMajor() Version {
	return Version{
		Major:  v.Major + 1,
		Minor:  0,
		Bugfix: 0,
	}
}

func (v Version) IncrementMinor() Version {
	return Version{
		Major:  v.Major,
		Minor:  v.Minor + 1,
		Bugfix: 0,
	}
}

func (v Version) IncrementBugfix() Version {
	return Version{
		Major:  v.Major,
		Minor:  v.Minor,
		Bugfix: v.Bugfix + 1,
	}
}

func (v Version) Compare(other Version) int {
	if v.Major != other.Major {
		return other.Major - v.Major
	}
	if v.Minor != other.Minor {
		return other.Minor - v.Minor
	}
	return other.Bugfix - v.Bugfix
}
