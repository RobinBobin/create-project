package rnpaper

import (
	"fmt"

	"github.com/robinbobin/create-project/assets"
	"github.com/robinbobin/create-project/utils"
)

func addBabelPlugin() {
	if !utils.Confirm("Would you like to add the 'react-native-paper/babel' plugin?", true) {
		return
	}

	assets.CreateBabelConfig()

	fmt.Printf(
		"Please add\n%v\nto your '%v'.\n",
		`env: {
  production: {
    plugins: ['react-native-paper/babel']
  }
}`,
		assets.BABEL_CONFIG_JS,
	)
}
