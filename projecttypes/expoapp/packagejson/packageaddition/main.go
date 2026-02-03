package packageaddition

import (
	"github.com/charmbracelet/huh"
	"github.com/robinbobin/create-project/utils"
)

func Add() {
	actions := []*utils.Action[func()]{
		{
			Fn:   addRobinBobinTSConfig,
			Name: "@robinbobin/ts-config",
		},
		{
			Fn:   addRobinBobinESLintConfigRN,
			Name: "@robinbobin/eslint-config-react-native",
		},
		{
			Fn:   addRobinBobinPrettierConfig,
			Name: "@robinbobin/prettier-config",
		},
	}

	utils.PanicOnError(
		huh.NewMultiSelect[*utils.Action[func()]]().
			Title("Which of the following packages would you like to install?").
			Options(huh.NewOptions(actions...)...).
			Value(&actions).
			Run(),
	)

	for _, action := range actions {
		action.Fn()
	}
}
