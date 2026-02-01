package expoapp

import (
	"github.com/robinbobin/create-project/projecttypes/expoapp/appjson"
	"github.com/robinbobin/create-project/projecttypes/expoapp/packagejson"
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

	deleteNodeLinkerHoisted()

	packagejson.Lint()
	appjson.Lint()

	return true
}
