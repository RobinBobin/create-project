package tsconfig

import (
	"fmt"
	"slices"
)

func processFiles(fileNames []string, tsconfig map[string]any) {
	if len(fileNames) == 0 {
		return
	}

	const key = "files"

	switch rawFiles := tsconfig[key].(type) {
	case []any:
		for _, rawFile := range rawFiles {
			fileNames = append(fileNames, rawFile.(string))
		}

		slices.Sort(fileNames)

		fileNames = slices.Compact(fileNames)

	case nil:
		slices.Sort(fileNames)

	default:
		panic(fmt.Errorf("\"%v\": \"%v\" is of type \"%T\", equals \"%v\" and can't be parsed", tsconfig_json, key, rawFiles, rawFiles))
	}

	tsconfig[key] = fileNames
}
