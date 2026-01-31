package expoapp

import (
	"fmt"
	"os/exec"

	"github.com/robinbobin/create-project/projecttypes/expoapp/missingplugins"
	"github.com/robinbobin/create-project/utils"
)

func Create() bool {
	defer utils.RecoverFromPanic()

	appName, mustApproveBuilds := createApp()

	setDir := func(cmd *exec.Cmd) {
		cmd.Dir = appName
	}

	if mustApproveBuilds {
		approveBuilds(setDir)
	}

	utils.UsePNPMInDir(appName)

	fmt.Println()

	utils.AskSortJSONInDir("app.json", appName)
	utils.AskSortJSONInDir("package.json", appName)

	deleteNodeLinkerHoisted(setDir)

	appPath := checkPathIsCorrect("my-app")

	missingplugins.AddMissingPlugins(appPath)

	return true
}
