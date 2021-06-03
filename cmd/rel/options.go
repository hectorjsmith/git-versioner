package rel

import "gitlab.com/hectorjsmith/git-versioner/util"

type CommandOptions struct {
	Major   bool
	Minor   bool
	Test    bool
	Message string
}

func (opt CommandOptions) isValid() bool {
	return util.CountTrueValues(opt.Major, opt.Minor, opt.Test) > 1
}
