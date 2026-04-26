package main

import (
	"maps"
	"slices"

	"github.com/charmbracelet/huh"
	"github.com/robinbobin/create-project/projecttypes/expoapp"
	"github.com/robinbobin/create-project/projecttypes/npmpackage"
	"github.com/robinbobin/create-project/utils"
)

func getAction(projectType projectType) *utils.Action[func() bool] {
	actions := map[string]*utils.Action[func() bool]{
		project_type_expo: {Fn: expoapp.Create, Name: "Create an Expo app"},
		project_type_npm:  {Fn: npmpackage.Create, Name: "Create an npm package"},
		"exit":            {Name: "Exit"},
	}

	action := actions[projectType]

	if action != nil {
		return action
	}

	utils.PanicOnError(
		huh.NewSelect[*utils.Action[func() bool]]().
			Title("What would you like:").
			Options(huh.NewOptions(slices.Collect(maps.Values(actions))...)...).
			Value(&action).
			Run(),
	)

	return action
}
