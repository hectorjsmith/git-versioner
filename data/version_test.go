package data

import (
	"testing"
)

func Test_GivenVersionDataWhenIncrementMajorThenCorrectNewVersionDataReturned(t *testing.T) {
	// given
	before := Version{Major: 1, Minor: 1, Bugfix: 1}

	// when
	after := before.IncrementMajor()

	// then
	assertEqualInt(t, before.Major+1, after.Major)
	assertEqualInt(t, 0, after.Minor)
	assertEqualInt(t, 0, after.Bugfix)
}

func Test_GivenVersionDataWhenIncrementMinorThenCorrectNewVersionDataReturned(t *testing.T) {
	// given
	before := Version{Major: 1, Minor: 1, Bugfix: 1}

	// when
	after := before.IncrementMinor()

	// then
	assertEqualInt(t, before.Major, after.Major)
	assertEqualInt(t, before.Minor+1, after.Minor)
	assertEqualInt(t, 0, after.Bugfix)
}

func Test_GivenVersionDataWhenIncrementBugfixThenCorrectNewVersionDataReturned(t *testing.T) {
	// given
	before := Version{Major: 1, Minor: 1, Bugfix: 1}

	// when
	after := before.IncrementBugfix()

	// then
	assertEqualInt(t, before.Major, after.Major)
	assertEqualInt(t, before.Minor, after.Minor)
	assertEqualInt(t, before.Bugfix+1, after.Bugfix)
}

func Test_GivenVersionDataWhenGetVersionStringThenCorrectStringReturned(t *testing.T) {
	assertVersionDataStringMatchesExpected(t, Version{0, 0, 0}, "0.0.0")
	assertVersionDataStringMatchesExpected(t, Version{1, 0, 0}, "1.0.0")
	assertVersionDataStringMatchesExpected(t, Version{0, 1, 0}, "0.1.0")
	assertVersionDataStringMatchesExpected(t, Version{0, 0, 1}, "0.0.1")
	assertVersionDataStringMatchesExpected(t, Version{12, 15, 102}, "12.15.102")
}

func assertVersionDataStringMatchesExpected(t *testing.T, versionData Version, expectedString string) {
	// given: versionData

	// when
	versionString := versionData.String()

	// then
	t.Logf("%s == %s", expectedString, versionString)
	assertEqualString(t, expectedString, versionString)
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
