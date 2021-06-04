package list

import (
	"fmt"
	"gitlab.com/hectorjsmith/git-versioner/git"
)

func Run(options CommandOptions) error {
	includeTestVersions := options.Test && (options.Verbose || options.Tag)
	tags := git.GetSortedValidTags(includeTestVersions)

	for i := 0; i < len(tags); i++ {
		tag := tags[i]
		if options.Tag {
			fmt.Println(tag.Tag)
		} else {
			fmt.Println(tag.Version)
		}
	}

	return nil
}
