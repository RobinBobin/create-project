package eslintconfig

func Process() {
	addPackages()

	baseConfigs := getBaseConfigs()

	createConfig(baseConfigs)
	copyBaseConfigs(baseConfigs)
}
