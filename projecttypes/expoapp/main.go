package expoapp

import (
	"github.com/robinbobin/create-project/assets"
	"github.com/robinbobin/create-project/projecttypes/expoapp/appjson"
	"github.com/robinbobin/create-project/projecttypes/expoapp/eslintconfig"
	"github.com/robinbobin/create-project/projecttypes/expoapp/packagejson"
	"github.com/robinbobin/create-project/projecttypes/expoapp/tsconfig"
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

	packageJsonOptions := packagejson.Process()
	eslintConfigOptions := eslintconfig.Process(packageJsonOptions)

	tsconfig.Process(eslintConfigOptions, packageJsonOptions)

	appjson.Lint()

	return true
}
