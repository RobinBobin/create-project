package main

import (
	"fmt"
	"os"

	"github.com/charmbracelet/huh"
	"github.com/robinbobin/create-project/projecttypes/expoapp"
	"github.com/robinbobin/create-project/projecttypes/npmpackage"
	"github.com/robinbobin/create-project/utils"
)

func main() {
	actions := []*utils.Action[func() bool]{
		{Fn: npmpackage.Create, Name: "Create an npm package"},
		{Fn: expoapp.Create, Name: "Create an Expo app"},
		{Name: "Exit"},
	}

	var action *utils.Action[func() bool]

	utils.PanicOnError(
		huh.NewSelect[*utils.Action[func() bool]]().
			Title("What would you like:").
			Options(huh.NewOptions(actions...)...).
			Value(&action).
			Run(),
	)

	farewell := "Bye."

	if action.Fn != nil {
		wd, err := os.Getwd()
		utils.PanicOnError(err)

		if action.Fn() {
			farewell = "Done."
		}

		utils.PanicOnError(os.Chdir(wd))
	}

	fmt.Println(farewell)
}
