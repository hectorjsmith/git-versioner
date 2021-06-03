package data

import (
	"reflect"
	"testing"
)

func Test_GivenVersionStringWhenParseIntoVersionDataThenCorrectDataRead(t *testing.T) {
	assertParsedVersionDataFromStringMatchesExpected(t, "1.1.1", Version{1, 1, 1})
	assertParsedVersionDataFromStringMatchesExpected(t, "v7.1.2-snapshot", Version{7, 1, 2})
	assertParsedVersionDataFromStringMatchesExpected(t, "rel/version4.2.9.12.dirty", Version{4, 2, 9})
	assertParsedVersionDataFromStringMatchesExpected(t, "rel/version4.999.121.dirty", Version{4, 999, 121})
}

func assertParsedVersionDataFromStringMatchesExpected(t *testing.T, versionString string, expectedData Version) {
	// given: versionString

	// when
	versionData := NewVersionFromString(versionString)

	// then
	t.Logf("%s == %s", expectedData.String(), versionData.String())
	assertEqualInt(t, expectedData.Major, versionData.Major)
	assertEqualInt(t, expectedData.Minor, versionData.Minor)
	assertEqualInt(t, expectedData.Bugfix, versionData.Bugfix)
}

func Test_GivenTagSliceWhenFilterToValidTagsThenOnlyValidTagsReturned(t *testing.T) {
	type args struct {
		tags []string
	}
	tests := []struct {
		name string
		args args
		want []VersionTag
	}{
		{
			name: "test: 2 valid",
			args: args{[]string{"1.1.1", "v1.1.1"}},
			want: []VersionTag{
				{
					Version: Version{1, 1, 1},
					Tag:     "1.1.1",
				},
				{
					Version: Version{1, 1, 1},
					Tag:     "v1.1.1",
				},
			},
		},
		{
			name: "test: 1 valid (a)",
			args: args{[]string{"1.1.1", "v1.X.1"}},
			want: []VersionTag{
				{
					Version: Version{1, 1, 1},
					Tag:     "1.1.1",
				},
			},
		},
		{
			name: "test: 1 valid (b)",
			args: args{[]string{"1.1", "v1.1.1"}},
			want: []VersionTag{
				{
					Version: Version{1, 1, 1},
					Tag:     "v1.1.1",
				},
			},
		},
		{
			name: "test: none valid",
			args: args{[]string{"1.1.Y", "v.1.1"}},
			want: []VersionTag{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := filterToOnlyValidTags(tt.args.tags, true)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("filterToOnlyValidTags() = %v, want %v", got, tt.want)
			}
		})
	}
}
