package eslintconfig

import (
	"fmt"
	"os"

	"github.com/charmbracelet/huh"
	"github.com/robinbobin/create-project/options"
	"github.com/robinbobin/create-project/utils"
)

func handleMissingTypes() {
	extension := ""

	if options.Options.IsESM {
		extension = ".js"
	}

	dummyTypings := []string{
		fmt.Sprint("eslint-config-expo/flat", extension),
	}

	utils.PanicOnError(
		huh.NewMultiSelect[string]().
			Title("Which dummy typings do you need?").
			Options(huh.NewOptions(dummyTypings...)...).
			Value(&dummyTypings).
			Run(),
	)

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

	options.Options.AddFile(utils.CUSTOM_D_TS)
}
