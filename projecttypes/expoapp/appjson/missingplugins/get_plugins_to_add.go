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
	pluginsToAdd := []string{}

	pluginsToCheck := utils.FilterOutUninstalled(
		[]string{
			"expo-font",
			"expo-splash-screen",
			"expo-system-ui",
		},
	)

	if len(pluginsToCheck) == 0 {
		return pluginsToAdd
	}

	for _, rawPlugin := range plugins {
		currentPlugins = append(currentPlugins, getPluginName(rawPlugin))
	}

	for _, pluginToCheck := range pluginsToCheck {
		if !slices.Contains(currentPlugins, pluginToCheck) {
			missingPlugins = append(missingPlugins, pluginToCheck)
		}
	}

	if len(missingPlugins) == 0 {
		return pluginsToAdd
	}

	fmt.Println()

	prompt := &survey.MultiSelect{
		Default: missingPlugins,
		Message: "Which missing plugins would you like to add to 'app.json'",
		Options: missingPlugins,
	}

	utils.PanicOnError(survey.AskOne(prompt, &pluginsToAdd))

	return pluginsToAdd
}
