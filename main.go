package main

import (
	"fmt"
	"slices"

	"github.com/robinbobin/create-project/projecttypes/expoapp"
	"github.com/robinbobin/create-project/projecttypes/npmpackage"
	"github.com/robinbobin/create-project/utils"
)

func main() {
	actions := []func() bool{npmpackage.Create, expoapp.Create}

	const optionExit = "Exit"
	options := []string{"Create an npm package", "Create an Expo app", optionExit}

	result := utils.AskOne(options)

	farewell := "Bye."

	if result != optionExit {
		actionIndex := slices.Index(options, result)

		if actions[actionIndex]() {
			farewell = "Done."
		}
	}

	fmt.Println(farewell)
}
