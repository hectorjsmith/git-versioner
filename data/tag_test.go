package data

import (
	"reflect"
	"testing"
)

func Test_GivenTagSliceWhenSortingTagsThenTagsSortedInCorrectOrder(t *testing.T) {
	type args struct {
		tags []VersionTag
	}
	tests := []struct {
		name string
		args args
		want []VersionTag
	}{
		{name: "test 1", args: args{[]VersionTag{
			{
				Version: Version{1, 0, 0},
			},
			{
				Version: Version{1, 1, 0},
			},
			{
				Version: Version{0, 9, 0},
			},
			{
				Version: Version{0, 9, 1},
			},
		}}, want: []VersionTag{
			{
				Version: Version{1, 1, 0},
			},
			{
				Version: Version{1, 0, 0},
			},
			{
				Version: Version{0, 9, 1},
			},
			{
				Version: Version{0, 9, 0},
			},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			versionTagSlice(tt.args.tags).Sort()
			got := tt.args.tags

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortTags() = %v, want %v", got, tt.want)
			}
		})
	}
}
