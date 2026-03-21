package eslintconfig

func Process(isTypeSet bool) *Options {
	options := &Options{
		CustomDTS: handleMissingTypes(),
	}

	useBaseConfigs()

	return options
}
