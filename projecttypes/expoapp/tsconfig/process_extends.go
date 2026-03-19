package tsconfig

import (
	"fmt"
	"slices"
)

func processExtends(tsconfig map[string]any) {
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

	extends = append(extends, "./tsconfig.base")

	slices.Sort(extends)

	tsconfig["extends"] = extends
}
