package packagejson

import (
	"github.com/robinbobin/create-project/options"
	"github.com/robinbobin/create-project/projecttypes/expoapp/packagejson/addition"
)

func Process() {
	resetProject()
	options.Options.IsESM = setType()

	uninstallPackages()
	addition.Run()
}
