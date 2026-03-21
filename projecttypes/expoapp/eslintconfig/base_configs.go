package eslintconfig

import (
	"fmt"
	"slices"
	"strings"

	"github.com/charmbracelet/huh"
	"github.com/robinbobin/create-project/utils"
)

type config struct {
	dependencies []string
	filePostfix  string
}

func useBaseConfigs() {
	generalConfig := "General config"
	reactConfig := "React config"
	reactNativeConfig := "React Native config"

	configs := []string{
		generalConfig,
		reactConfig,
		reactNativeConfig,
	}

	utils.PanicOnError(
		huh.NewMultiSelect[string]().
			Title("Would you like to add a custom base eslint config?").
			Options(huh.NewOptions(configs...)...).
			Value(&configs).
			Validate(func(selection []string) error {
				if len(selection) == 0 || slices.Contains(selection, generalConfig) {
					return nil
				}

				return fmt.Errorf("The general config is required.")
			}).
			Run(),
	)

	if len(configs) == 0 {
		return
	}

	data := map[string]config{
		generalConfig: {
			dependencies: []string{
				"@eslint/js",
				"@stylistic/eslint-plugin",
				"@typescript-eslint/parser",
				"eslint-import-resolver-typescript",
				"eslint-plugin-import-x",
				"eslint-plugin-promise",
				"eslint-plugin-simple-import-sort",
				"typescript-eslint",
			},
		},
		reactConfig: {
			dependencies: []string{
				"eslint-plugin-react-hooks",
			},
			filePostfix: "react",
		},
		reactNativeConfig: {
			filePostfix: "react.native",
		},
	}

	sb := strings.Builder{}

	for _, config := range configs {
		sb.Reset()
		sb.WriteString("eslint.config")

		datum := data[config]

		if len(datum.filePostfix) != 0 {
			sb.WriteString(".")
			sb.WriteString(datum.filePostfix)
		}

		sb.WriteString(".js")

		fileName := sb.String()

		fmt.Println(fileName, datum.dependencies)

		// assets.CopyFile(fileName, fileName)
	}
}
