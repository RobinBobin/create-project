package expoapp

import (
	"fmt"
	"path/filepath"
	"slices"

	"github.com/AlecAivazis/survey/v2"
	"github.com/robinbobin/create-project/utils"
)

func addMissingPlugins(appPath string) {
	jsonData := utils.ReadJSON(filepath.Join(appPath, "app.json"))

	plugins := jsonData["expo"].(map[string]any)["plugins"].([]any)

	currentPlugins := []string{}
	missingPlugins := []string{}
	pluginsToCheck := []string{"expo-font", "expo-system-ui"}

	for _, raw := range plugins {
		var pluginName string

		switch plugin := raw.(type) {
		case string:
			pluginName = plugin

		case []any:
			pluginName = plugin[0].(string)
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
		Message: "Which missing plugins would you like to add to 'app.json'",
		Options: missingPlugins,
	}

	selectedPlugins := []string{}

	survey.AskOne(prompt, &selectedPlugins)

	fmt.Println(selectedPlugins)
}
