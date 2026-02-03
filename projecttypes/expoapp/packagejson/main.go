package packagejson

import (
	"github.com/robinbobin/create-project/projecttypes/expoapp/packagejson/packageaddition"
	"github.com/robinbobin/create-project/utils"
)

func Lint() {
	utils.AskSortJSON("package.json")

	packageaddition.Add()

	uninstallPackages()
}
