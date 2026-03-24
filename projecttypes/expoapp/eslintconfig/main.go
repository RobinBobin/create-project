package eslintconfig

import "github.com/robinbobin/create-project/projecttypes/expoapp/packagejson"

const eslintConfigJS = "eslint.config.js"

func Process(packageJsonOptions *packagejson.Options) *Options {
	options := &Options{Files: []string{eslintConfigJS}}

	handleMissingTypes(options)
	useBaseConfigs(options)
	replaceRequire(packageJsonOptions.IsModule)
	run()

	return options
}
