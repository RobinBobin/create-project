package eslintconfig

import "github.com/robinbobin/create-project/projecttypes/expoapp/packagejson"

func Process(packageJsonOptions *packagejson.Options) *Options {
	options := &Options{Files: []string{"eslint.config.js"}}

	handleMissingTypes(options)
	useBaseConfigs(options)

	return options
}
