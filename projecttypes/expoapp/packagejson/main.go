package packagejson

import (
	"github.com/robinbobin/create-project/utils"
)

func Lint() *Options {
	options := &Options{}

	setType(options)
	resetProject(options)

	if !options.IsModule {
		utils.AskSortJSON(utils.PACKAGE_JSON)
	}

	uninstallPackages(options.IsProjectReset)

	return options
}
