package data

import (
	"testing"
)

func Test_GivenVersionDataWhenIncrementMajorThenCorrectNewVersionDataReturned(t *testing.T) {
	// given
	before := VersionData{Major: 1, Minor: 1, Bugfix: 1}

	// when
	after := before.IncrementMajor()

	// then
	assertEqualInt(t, before.Major+1, after.Major)
	assertEqualInt(t, 0, after.Minor)
	assertEqualInt(t, 0, after.Bugfix)
}

func Test_GivenVersionDataWhenIncrementMinorThenCorrectNewVersionDataReturned(t *testing.T) {
	// given
	before := VersionData{Major: 1, Minor: 1, Bugfix: 1}

	// when
	after := before.IncrementMinor()

	// then
	assertEqualInt(t, before.Major, after.Major)
	assertEqualInt(t, before.Minor+1, after.Minor)
	assertEqualInt(t, 0, after.Bugfix)
}

func Test_GivenVersionDataWhenIncrementBugfixThenCorrectNewVersionDataReturned(t *testing.T) {
	// given
	before := VersionData{Major: 1, Minor: 1, Bugfix: 1}

	// when
	after := before.IncrementBugfix()

	// then
	assertEqualInt(t, before.Major, after.Major)
	assertEqualInt(t, before.Minor, after.Minor)
	assertEqualInt(t, before.Bugfix+1, after.Bugfix)
}

func Test_GivenVersionDataWhenGetVersionStringThenCorrectStringReturned(t *testing.T) {
	assertVersionDataStringMatchesExpected(t, VersionData{0, 0, 0}, "0.0.0")
	assertVersionDataStringMatchesExpected(t, VersionData{1, 0, 0}, "1.0.0")
	assertVersionDataStringMatchesExpected(t, VersionData{0, 1, 0}, "0.1.0")
	assertVersionDataStringMatchesExpected(t, VersionData{0, 0, 1}, "0.0.1")
	assertVersionDataStringMatchesExpected(t, VersionData{12, 15, 102}, "12.15.102")
}

func assertVersionDataStringMatchesExpected(t *testing.T, versionData VersionData, expectedString string) {
	// given: versionData

	// when
	versionString := versionData.VersionString()

	// then
	t.Logf("%s == %s", expectedString, versionString)
	assertEqualString(t, expectedString, versionString)
}

func Test_GivenVersionStringWhenParseIntoVersionDataThenCorrectDataRead(t *testing.T) {
	assertParsedVersionDataFromStringMatchesExpected(t, "1.1.1", VersionData{1, 1, 1})
	assertParsedVersionDataFromStringMatchesExpected(t, "v7.1.2-snapshot", VersionData{7, 1, 2})
	assertParsedVersionDataFromStringMatchesExpected(t, "rel/version4.2.9.12.dirty", VersionData{4, 2, 9})
	assertParsedVersionDataFromStringMatchesExpected(t, "rel/version4.999.121.dirty", VersionData{4, 999, 121})
}

func assertParsedVersionDataFromStringMatchesExpected(t *testing.T, versionString string, expectedData VersionData) {
	// given: versionString

	// when
	versionData := NewVersionDataFromString(versionString)

	// then
	t.Logf("%s == %s", expectedData.VersionString(), versionData.VersionString())
	assertEqualInt(t, expectedData.Major, versionData.Major)
	assertEqualInt(t, expectedData.Minor, versionData.Minor)
	assertEqualInt(t, expectedData.Bugfix, versionData.Bugfix)
}

func assertEqualString(t *testing.T, expected string, actual string) {
	if expected != actual {
		t.Errorf("Assert failed. Expected '%s' but got '%s'", expected, actual)
	}
}

func assertEqualInt(t *testing.T, expected int, actual int) {
	if expected != actual {
		t.Errorf("Assert failed. Expected %d but got %d", expected, actual)
	}
}
