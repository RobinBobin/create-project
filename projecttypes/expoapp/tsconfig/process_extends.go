package tsconfig

import (
	"fmt"
	"path/filepath"

	"github.com/robinbobin/create-project/assets"
	"github.com/robinbobin/create-project/utils"
)

func processExtends(tsconfig map[string]any) {
	if !utils.Confirm("Would you like to add a custom base tsconfig?", true) {
		return
	}

	const baseFile = "tsconfig.base"

	assets.CopyFile(fmt.Sprintf("%v.json", baseFile), tsconfig_json)

	extends := []string{}

	switch ext := tsconfig["extends"].(type) {
	case string:
		extends = append(extends, ext)

	case []any:
		for _, baseConfig := range ext {
			extends = append(extends, baseConfig.(string))
		}

	default:
		panic(fmt.Errorf("\"%v\": \"extends\" is of type \"%T\", equals \"%v\" and can't be parsed", tsconfig_json, ext, ext))
	}

	extends = append(extends, fmt.Sprintf(".%c%v", filepath.Separator, baseFile))

	tsconfig["extends"] = extends
}
