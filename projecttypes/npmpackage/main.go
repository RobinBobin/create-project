package npmpackage

import (
	"github.com/robinbobin/create-project/utils"
)

func Create() bool {
	defer utils.RecoverFromPanic()

	initPackage()

	return true
}
