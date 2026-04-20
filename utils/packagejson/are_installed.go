package packagejson

import "slices"

func AreInstalled(packageNames []string) bool {
	installedPackages := GetInstalled()

	for _, packageName := range packageNames {
		if !slices.Contains(installedPackages, packageName) {
			return false
		}
	}

	return true
}
