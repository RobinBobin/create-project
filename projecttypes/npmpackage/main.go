package npmpackage

import (
	"errors"
	"os"

	"github.com/robinbobin/create-project/utils"
)

func Create() bool {
	defer utils.RecoverFromPanic()

	utils.RunCmd("npm init")

	_, err := os.Stat("package.json")

	if errors.Is(err, os.ErrNotExist) {
		return false
	}

	utils.PanicOnError(err)
	utils.UsePNPM()
	utils.AskSortJSON("package.json")

	return true
}
