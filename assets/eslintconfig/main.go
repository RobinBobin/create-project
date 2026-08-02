package eslintconfig

func Process() {
	configNames, dependencies := getConfigData()

	addPackages(dependencies)
	createConfig(configNames)
	copyBaseConfigs(configNames)
}
