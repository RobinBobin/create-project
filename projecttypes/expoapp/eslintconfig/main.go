package eslintconfig

func Process(isTypeSet bool) *Options {
	options := &Options{
		CustomDTS: handleMissingTypes(),
	}

	return options
}
