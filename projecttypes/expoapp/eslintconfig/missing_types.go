package eslintconfig

import (
	"fmt"
	"os"

	"github.com/charmbracelet/huh"
	"github.com/robinbobin/create-project/options"
	"github.com/robinbobin/create-project/utils"
	"github.com/robinbobin/create-project/utils/packagejson"
)

func handleMissingTypes() {
	extension := ".js"
	dummyTypings := []string{}

	if packagejson.IsInstalled("eslint-config-expo") {
		dummyTypings = append(dummyTypings, fmt.Sprint("eslint-config-expo/flat", extension))
	}

	if len(dummyTypings) == 0 {
		return
	}

	if !options.Options.ShouldUseDefaults {
		utils.PanicOnError(
			huh.NewMultiSelect[string]().
				Title("Which dummy typings do you need?").
				Options(huh.NewOptions(dummyTypings...)...).
				Value(&dummyTypings).
				Run(),
		)
	}

	if len(dummyTypings) == 0 {
		return
	}

	file, err := os.OpenFile(utils.CUSTOM_D_TS, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	utils.PanicOnError(err)

	defer func() {
		_ = file.Close()
	}()

	for _, moduleName := range dummyTypings {
		_, err = file.WriteString(utils.FormatDeclareModule(moduleName))

		utils.PanicOnError(err)
	}

	options.Options.TS.AddFile(utils.CUSTOM_D_TS)
}
