package expoapp

import (
	"errors"
	"regexp"

	"github.com/robinbobin/create-project/utils"
)

func createApp() (appName string, mustApproveBuilds bool) {
	appNameRe := regexp.MustCompile(`What is your app named\? … (\w+)`)

	approveBuildsRe := regexp.MustCompile(`Run "pnpm approve-builds" to pick which dependencies should be allowed to run scripts\.`)

	utils.CaptureCmdOutput(&utils.CaptureCmdOutputOptions{
		CmdWithArgs: "pnpm create expo-app --template",
		CapturedOutputProcessor: func(strippedOutput string) (needsMoreStdin bool) {
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
	})

	if len(appName) == 0 {
		utils.PanicOnError(errors.New("the app name could not be determined 🙁"))
	}

	return appName, mustApproveBuilds
}
