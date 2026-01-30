package expoapp

import (
	"fmt"
	"path/filepath"
	"slices"

	"github.com/AlecAivazis/survey/v2"
	"github.com/robinbobin/create-project/utils"
)

type expoAppConfig struct {
	Expo appConfig `json:"expo"`
}

type appConfig struct {
	Plugins []any `json:"plugins,omitempty"`
}

func addMissingPlugins(appPath string) {
	jsonData := &expoAppConfig{}

	utils.ReadJSON(&jsonData, filepath.Join(appPath, "app.json"))

	currentPlugins := []string{}
	missingPlugins := []string{}
	pluginsToCheck := []string{"expo-font", "expo-router", "expo-system-ui"}

	for _, raw := range jsonData.Expo.Plugins {
		var pluginName string

		switch plugin := raw.(type) {
		case string:
			pluginName = plugin

		case []any:
			pluginName = plugin[0].(string)

		default:
			fmt.Printf("%T %v", plugin, plugin)
		}

		currentPlugins = append(currentPlugins, pluginName)
	}

	for _, pluginToCheck := range pluginsToCheck {
		if !slices.Contains(currentPlugins, pluginToCheck) {
			missingPlugins = append(missingPlugins, pluginToCheck)
		}
	}

	fmt.Println()

	prompt := &survey.MultiSelect{
		Default: missingPlugins,
		Message: "Would you like to add these missing plugins to 'app.json'",
		Options: missingPlugins,
	}

	selectedPlugins := []string{}

	survey.AskOne(prompt, &selectedPlugins)

	fmt.Println(selectedPlugins)
}
