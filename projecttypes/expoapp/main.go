package expoapp

import (
	"os/exec"

	"github.com/robinbobin/create-project/utils"
)

func Create() bool {
	defer utils.RecoverFromPanic()

	appName, mustApproveBuilds := createApp()

	preRunner := func(cmd *exec.Cmd) {
		cmd.Dir = appName
	}

	if mustApproveBuilds {
		approveBuilds(preRunner)
	}

	utils.UsePNPMInDir(appName)
	utils.AskSortJSONInDir("app.json", appName)
	utils.AskSortJSONInDir("package.json", appName)

	utils.RunCmdWithPreRunner(
		"pnpm config --location project delete node-linker",
		preRunner,
	)

	return true
}
