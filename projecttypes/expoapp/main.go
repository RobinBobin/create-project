package expoapp

import (
	"fmt"

	"github.com/robinbobin/create-project/projecttypes/expoapp/missingplugins"
	"github.com/robinbobin/create-project/utils"
)

func Create() bool {
	defer utils.RecoverFromPanic()

	appName, mustApproveBuilds := createApp()

	checkPathIsCorrect(appName)

	if mustApproveBuilds {
		approveBuilds()
	}

	utils.UsePNPM()

	fmt.Println()

	utils.AskSortJSON("app.json")
	utils.AskSortJSON("package.json")

	deleteNodeLinkerHoisted()

	missingplugins.AddMissingPlugins()

	return true
}
