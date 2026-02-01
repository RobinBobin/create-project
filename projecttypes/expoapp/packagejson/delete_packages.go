package packagejson

import (
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"github.com/robinbobin/create-project/utils"
)

func deletePackages() {
	packages := utils.FilterOutUninstalled(
		[]string{
			"expo-haptics",
			"expo-symbols",
		},
	)

	if len(packages) == 0 {
		return
	}

	prompt := survey.MultiSelect{
		Default: packages,
		Message: "Which of the following packages would you like to delete",
		Options: packages,
	}

	selectedPackages := []string{}

	utils.PanicOnError(survey.AskOne(&prompt, &selectedPackages))

	if len(selectedPackages) == 0 {
		return
	}

	utils.RunCmd("pnpm uninstall " + strings.Join(selectedPackages, " "))
}
