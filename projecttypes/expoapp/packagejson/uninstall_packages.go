package packagejson

import (
	"strings"

	"github.com/charmbracelet/huh"
	"github.com/robinbobin/create-project/utils"
)

func uninstallPackages() {
	packages := utils.FilterOutUninstalled(
		[]string{
			"expo-haptics",
			"expo-symbols",
		},
	)

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
