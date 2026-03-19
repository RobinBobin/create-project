package tsconfig

import (
	"fmt"
	"slices"
)

func addToFiles(fileName string, tsconfig map[string]any) {
	if len(fileName) == 0 {
		return
	}

	const key = "files"

	var files []string

	switch rawFiles := tsconfig[key].(type) {
	case []any:
		files = make([]string, len(rawFiles))

		for index := range files {
			files[index] = rawFiles[index].(string)
		}

	case []string:
		files = rawFiles

	case nil:
		// Nothing to do.

	default:
		panic(fmt.Errorf("\"%v\": \"%v\" is of type \"%T\", equals \"%v\" and can't be parsed", tsconfig_json, key, rawFiles, rawFiles))
	}

	if files == nil {
		tsconfig[key] = []string{fileName}
	} else {
		if !slices.Contains(files, fileName) {
			files = append(files, fileName)
		}

		slices.Sort(files)

		tsconfig[key] = files
	}
}
