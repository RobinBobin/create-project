package packagejson

import (
	"github.com/robinbobin/create-project/utils"
)

func Lint() *Options {
	isTypeSet := setModuleType()

	options := &Options{
		IsProjectReset: resetProject(isTypeSet),
		IsTypeSet:      isTypeSet,
	}

	if !options.IsTypeSet {
		utils.AskSortJSON(utils.PACKAGE_JSON)
	}

	uninstallPackages()

	return options
}
