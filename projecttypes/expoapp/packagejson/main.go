package packagejson

import (
	"github.com/robinbobin/create-project/projecttypes/expoapp/packagejson/addition"
)

func Process() {
	resetProject()
	setType()
	uninstallPackages()

	addition.Run()
}
