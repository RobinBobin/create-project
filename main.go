package main

import (
	"fmt"
	"os"

	"github.com/charmbracelet/huh"
	"github.com/robinbobin/create-project/projecttypes/expoapp"
	"github.com/robinbobin/create-project/projecttypes/npmpackage"
	"github.com/robinbobin/create-project/utils"
)

type wrapper struct {
	action func() bool
	name   string
}

func (action *wrapper) String() string {
	return action.name
}

func main() {
	actions := []*wrapper{
		{action: npmpackage.Create, name: "Create an npm package"},
		{action: expoapp.Create, name: "Create an Expo app"},
		{name: "Exit"},
	}

	var action *wrapper

	utils.PanicOnError(
		huh.NewSelect[*wrapper]().
			Title("What would you like:").
			Options(huh.NewOptions(actions...)...).
			Value(&action).
			Run(),
	)

	farewell := "Bye."

	if action.action != nil {
		wd, err := os.Getwd()
		utils.PanicOnError(err)

		if action.action() {
			farewell = "Done."
		}

		utils.PanicOnError(os.Chdir(wd))
	}

	fmt.Println(farewell)
}
