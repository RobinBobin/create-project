package packagejson

import (
	"strings"

	"github.com/charmbracelet/huh"
	"github.com/robinbobin/create-project/utils"
	"github.com/robinbobin/create-project/utils/packagejson"
)

func uninstallPackages(isProjectReset bool) {
	packages := []string{}

	if isProjectReset {
		packages = append(
			packages,
			"expo-haptics",
			"expo-symbols",
			"@react-navigation/bottom-tabs",
		)
	}

	packages = packagejson.FilterOutUninstalled(packages)

	if len(packages) == 0 {
		return
	}

	utils.PanicOnError(
		huh.NewMultiSelect[string]().
			Title("Which of the following packages would you like to uninstall?").
			Options(huh.NewOptions(packages...)...).
			Value(&packages).
			Run(),
	)

	if len(packages) == 0 {
		return
	}

	utils.RunCmd("pnpm uninstall " + strings.Join(packages, " "))
}
