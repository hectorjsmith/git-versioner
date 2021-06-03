package git

import (
	"gitlab.com/hectorjsmith/git-versioner/data"
	"reflect"
	"testing"
)

func Test_findLatestBugfixTagForVersion(t *testing.T) {
	type args struct {
		versionToFind data.Version
		tags          []data.VersionTag
	}
	tests := []struct {
		name    string
		args    args
		want    data.VersionTag
		wantErr bool
	}{
		{
			name: "test 1",
			args: args{
				versionToFind: data.Version{Major: 1, Minor: 1, Bugfix: 0},
				tags: []data.VersionTag{
					{
						Version: data.Version{Major: 1, Bugfix: 12},
					},
					{
						Version: data.Version{Major: 1, Minor: 1},
					},
					{
						Version: data.Version{Minor: 9, Bugfix: 13},
					},
					{
						Version: data.Version{Major: 1, Minor: 1, Bugfix: 5},
					},
					{
						Version: data.Version{Major: 1, Minor: 1, Bugfix: 3},
					},
				},
			},
			want: data.VersionTag{
				Version: data.Version{Major: 1, Minor: 1, Bugfix: 5},
			},
			wantErr: false,
		},
		{name: "test 2",
			args: args{
				versionToFind: data.Version{Major: 1, Minor: 5, Bugfix: 5},
				tags: []data.VersionTag{
					{
						Version: data.Version{Major: 1, Bugfix: 12},
					},
					{
						Version: data.Version{Major: 1, Minor: 5, Bugfix: 4},
					},
					{
						Version: data.Version{Major: 1, Minor: 5, Bugfix: 6},
					},
				},
			},
			want: data.VersionTag{
				Version: data.Version{Major: 1, Minor: 5, Bugfix: 6},
			},
			wantErr: false,
		},
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
