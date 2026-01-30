package expoapp

import (
	"errors"
	"os/exec"
	"regexp"

	"github.com/robinbobin/create-project/utils"
)

func createApp() (appName string) {
	appNameRe := regexp.MustCompile(`What is your app named\? … (\w+)`)

	mustApproveBuilds := false
	approveBuildsRe := regexp.MustCompile(`Run "pnpm approve-builds" to pick which dependencies should be allowed to run scripts\.`)

	utils.CaptureCmdOutput(
		"pnpm create expo-app --template",
		func(strippedOutput string) (needsMoreStdin bool) {
			matches := appNameRe.FindStringSubmatch(strippedOutput)

			if matches != nil {
				appName = matches[1]

				return false
			}

			if !mustApproveBuilds {
				mustApproveBuilds = approveBuildsRe.FindStringIndex(strippedOutput) != nil
			}

			return len(appName) == 0
		},
	)

	if len(appName) == 0 {
		utils.PanicOnError(errors.New("the app name could not be determined 🙁"))
	}

	preRunner := func(cmd *exec.Cmd) {
		cmd.Dir = appName
	}

	if mustApproveBuilds {
		approveBuilds(preRunner)
	}

	utils.UsePNPMInDir(appName)
	utils.AskSortJSONInDir("package.json", appName)

	utils.RunCmdWithPreRunner(
		"pnpm config --location project delete node-linker",
		preRunnder,
	)

	return appName
}
