package npmpackage

import (
	"github.com/robinbobin/create-project/utils"
)

func Create() {
	defer utils.RecoverFromPanic()

	initPackage()
}
