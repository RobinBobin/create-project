package missingplugins

import (
	"fmt"
	"slices"

	"github.com/AlecAivazis/survey/v2"
	"github.com/robinbobin/create-project/utils"
)

func getPluginsToAdd(plugins []any) []string {
	currentPlugins := []string{}
	missingPlugins := []string{}
	pluginsToCheck := []string{"expo-font", "expo-system-ui"}

	for _, rawPlugin := range plugins {
		currentPlugins = append(currentPlugins, getPluginName(rawPlugin))
	}

	for _, pluginToCheck := range pluginsToCheck {
		if !slices.Contains(currentPlugins, pluginToCheck) {
			missingPlugins = append(missingPlugins, pluginToCheck)
		}
	}

	fmt.Println()

	prompt := &survey.MultiSelect{
		Default: missingPlugins,
		Message: "Which missing plugins would you like to add to 'app.json'",
		Options: missingPlugins,
	}

	selectedPlugins := []string{}

	utils.PanicOnError(survey.AskOne(prompt, &selectedPlugins))

	return selectedPlugins
}
