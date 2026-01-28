package npmpackage

import (
	"os"

	"github.com/robinbobin/create-project/utils"
)

func initPackage() {
	utils.RunCmd("npm", "init")

	_, err := os.Stat("package.json")

	utils.PanicOnError(err)

	utils.UsePNPM()

	utils.AskSortJSON("package.json")
}
