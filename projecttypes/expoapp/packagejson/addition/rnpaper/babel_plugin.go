package rnpaper

import (
	"fmt"

	"github.com/robinbobin/create-project/assets"
	"github.com/robinbobin/create-project/options"
	"github.com/robinbobin/create-project/utils"
)

func addBabelPlugin() {
	if !utils.Confirm("Would you like to add the 'react-native-paper/babel' plugin?", true) {
		return
	}

	assets.CreateBabelConfig()

	hint := fmt.Sprintf(
		"The following should be added to your '%v':\n%v",
		assets.BABEL_CONFIG_JS,
		`env: {
  production: {
    plugins: ['react-native-paper/babel']
  }
}`,
	)

	options.Options.Hints = append(options.Options.Hints, hint)
}
