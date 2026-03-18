package expoapp

import (
	"fmt"
	"os"
	"slices"
	"strings"

	"github.com/robinbobin/create-project/assets"
	"github.com/robinbobin/create-project/utils"
)

func useTSConfig(isProjectReset bool) {
	if !utils.Confirm("Would you like to use a custom tsconfig?", true) {
		return
	}

	// Copy assets/tsconfig.json as a base config.
	tsconfigName := "tsconfig.json"

	assets.CopyFile("tsconfig.base.json", tsconfigName)

	// Backup the current tsconfig.
	tsconfigFile, err := os.Open(tsconfigName)
	utils.CopyFile(fmt.Sprintf("%v.bak", tsconfigName), tsconfigFile, err)

	// Read the current tsconfig.
	tsconfig := utils.ReadJSON(tsconfigName)

	// Delete `compilerOptions` if the project was reset.
	if isProjectReset {
		delete(tsconfig, "compilerOptions")
	}

	// Modify `extends`.
	extends := []string{}

	switch ext := tsconfig["extends"].(type) {
	case string:
		extends = append(extends, ext)

	case []any:
		for _, baseConfig := range ext {
			extends = append(extends, baseConfig.(string))
		}

	default:
		panic(fmt.Errorf("\"%v\": \"extends\" is of type \"%T\", equals \"%v\" and can't be parsed", tsconfigName, ext, ext))
	}

	extends = append(extends, "./tsconfig.base")

	tsconfig["extends"] = extends

	// Modify `include`.
	rawInclude := tsconfig["include"].([]interface{})
	include := make([]string, len(rawInclude))

	for index := range include {
		pattern := rawInclude[index].(string)

		if isProjectReset && strings.HasPrefix(pattern, "**") {
			pattern = fmt.Sprint("src/", pattern)
		}

		include[index] = pattern
	}

	slices.Sort(include)

	tsconfig["include"] = include

	// Write tsconfig back.
	utils.WriteJSON(tsconfig, tsconfigName)
}
