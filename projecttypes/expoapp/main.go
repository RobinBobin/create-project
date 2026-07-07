package expoapp

import (
	"github.com/robinbobin/create-project/assets"
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

	assets.AskCreateVSCodeWorkspace(appName)

	utils.UsePNPM()

	deleteNodeLinkerHoisted()

	// packagejson.Process()

	// copySources()

	// eslintconfig.Process()
	// tsconfig.Process()
	// appjson.Lint()

	return true
}
