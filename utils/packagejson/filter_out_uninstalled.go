package packagejson

import (
	"slices"
)

func FilterOutUninstalled(packageNames []string) []string {
	installedPackages := GetInstalled()

	return slices.DeleteFunc(
		packageNames,
		func(packageName string) bool {

			return !slices.Contains(installedPackages, packageName)
		},
	)
}
