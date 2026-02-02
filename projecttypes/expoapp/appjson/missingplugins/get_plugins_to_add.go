package missingplugins

import (
	"slices"

	"github.com/charmbracelet/huh"
	"github.com/robinbobin/create-project/utils"
)

func getPluginsToAdd(plugins []any) []string {

	currentPlugins := []string{}
	missingPlugins := []string{}

	pluginsToCheck := utils.FilterOutUninstalled(
		[]string{
			"expo-font",
			"expo-splash-screen",
			"expo-system-ui",
		},
	)

	if len(pluginsToCheck) == 0 {
		return missingPlugins
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
		return missingPlugins
	}

	utils.PanicOnError(
		huh.NewMultiSelect[string]().
			Title("Which missing plugins would you like to add to 'app.json'?").
			Options(huh.NewOptions(missingPlugins...)...).
			Value(&missingPlugins).
			Run(),
	)

	return missingPlugins
}
