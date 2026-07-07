package missingplugins

import (
	"fmt"
	"slices"

	"github.com/charmbracelet/huh"
	"github.com/robinbobin/create-project/options"
	"github.com/robinbobin/create-project/utils"
	"github.com/robinbobin/create-project/utils/packagejson"
)

func getPluginsToAdd(plugins []any) []string {

	currentPlugins := []string{}
	missingPlugins := []string{}

	pluginsToCheck := packagejson.FilterOutUninstalled(
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

	if !options.Options.ShouldUseDefaults {
		utils.PanicOnError(
			huh.NewMultiSelect[string]().
				Title(fmt.Sprintf("Which missing plugins would you like to add to '%v'?", utils.APP_JSON)).
				Options(huh.NewOptions(missingPlugins...)...).
				Value(&missingPlugins).
				Run(),
		)
	}

	return missingPlugins
}
