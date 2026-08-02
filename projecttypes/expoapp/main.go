package expoapp

import (
	"github.com/robinbobin/create-project/assets"
	"github.com/robinbobin/create-project/assets/eslintconfig"
	"github.com/robinbobin/create-project/utils"
)

func Create() bool {
	defer utils.RecoverFromPanic()

	appName, mustApproveBuilds := createApp()

	if !checkPathIsCorrect(appName) {
		return false
	}

	if mustApproveBuilds {
		approveBuilds()
	}

	assets.CreateVSCodeWorkspace(appName)

	utils.UsePNPM()

	deleteNodeLinkerHoisted()
	resetProject()
	setPackageType()

	assets.AddPrettier()
	eslintconfig.Process()

	// copySources()

	// tsconfig.Process()
	// appjson.Lint()

	return true
}
