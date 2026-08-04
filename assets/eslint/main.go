package eslint

func Add() {
	configNames, dependencies := getConfigData()

	addPackages(dependencies)
	createConfig(configNames)
	copyBaseConfigs(configNames)
}
