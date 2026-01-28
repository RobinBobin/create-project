package main

import (
	"fmt"
	"slices"

	"github.com/robinbobin/create-project/projecttypes/npmpackage"
	"github.com/robinbobin/create-project/utils"
)

func main() {
	actions := []func(){npmpackage.Create}

	const optionExit = "Exit"
	options := []string{"Create an npm package", "Create an Expo app", optionExit}

	result := utils.AskOne(options)

	if result != optionExit {
		actionIndex := slices.Index(options, result)

		actions[actionIndex]()
	}

	fmt.Println("Bye")
}
