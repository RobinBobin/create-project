package packagejson

func IsInstalled(packageName string) bool {
	return AreInstalled([]string{packageName})
}
