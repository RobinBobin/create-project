package packagejson

import (
	"github.com/robinbobin/create-project/projecttypes/expoapp/packagejson/addition"
	"github.com/robinbobin/create-project/utils"
)

func Process() *Options {
	options := &Options{}

	options.IsProjectReset = resetProject()
	options.IsESM = setType()

	isSorted := options.IsESM || options.IsProjectReset

	if !isSorted {
		utils.AskSortJSON(utils.PACKAGE_JSON)
	}

	uninstallPackages(options.IsProjectReset)
	addition.Run()

	return options
}
