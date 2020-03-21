package versioner

import (
	"gitlab.com/hectorjsmith/git-versioner/data"
	"gitlab.com/hectorjsmith/git-versioner/versioner/data"
	"reflect"
	"testing"
)

func Test_GivenTagSliceWhenFilterToValidTagsThenOnlyValidTagsReturned(t *testing.T) {
	type args struct {
		tags []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{name: "test: 2 valid", args: args{[]string{"1.1.1", "v1.1.1"}}, want: []string{"1.1.1", "v1.1.1"}},
		{name: "test: 1 valid (a)", args: args{[]string{"1.1.1", "v1.X.1"}}, want: []string{"1.1.1"}},
		{name: "test: 1 valid (b)", args: args{[]string{"1.1", "v1.1.1"}}, want: []string{"v1.1.1"}},
		{name: "test: none valid", args: args{[]string{"1.1.Y", "v.1.1"}}, want: []string{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := filterToOnlyValidTags(tt.args.tags); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("filterToOnlyValidTags() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_GivenTagSliceWhenSortingTagsThenTagsSortedInCorrectOrder(t *testing.T) {
	type args struct {
		tags []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{name: "test 1", args: args{[]string{"1.0.0", "1.1.0", "0.9.0"}}, want: []string{"1.1.0", "1.0.0", "0.9.0"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortTags(tt.args.tags); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortTags() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findLatestBugfixForVersion(t *testing.T) {
	type args struct {
		versionToFind data.VersionData
		tags          []string
	}
	tests := []struct {
		name    string
		args    args
		want    versionerdata.TagVersionData
		wantErr bool
	}{
		{name: "test 1",
			args:    args{data.VersionData{1, 1, 0}, []string{"1.1.0", "1.1.1", "2.0.0", "1.1.9"}},
			want:    versionerdata.TagVersionData{data.VersionData{1, 1, 9}, "1.1.9"},
			wantErr: false},
		{name: "test 2",
			args:    args{data.VersionData{1, 5, 5}, []string{"1.1.0", "1.5.4", "1.5.6"}},
			want:    versionerdata.TagVersionData{data.VersionData{1, 5, 6}, "1.5.6"},
			wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := findLatestBugfixForVersion(tt.args.versionToFind, tt.args.tags)
			if (err != nil) != tt.wantErr {
				t.Errorf("findLatestBugfixForVersion() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findLatestBugfixForVersion() got = %v, want %v", got, tt.want)
			}
		})
	}
}
