package tsconfig

import "slices"

func addToFiles(fileName string, tsconfig map[string]any) {
	const key = "files"

	if rawValue, ok := tsconfig[key].([]any); !ok {
		tsconfig[key] = []string{fileName}
	} else {
		value := make([]string, len(rawValue))

		shouldAdd := true

		for index := range value {
			value[index] = rawValue[index].(string)

			if value[index] == fileName {
				shouldAdd = false
			}
		}

		if shouldAdd {
			value = append(value, fileName)
		}

		slices.Sort(value)

		tsconfig[key] = value
	}
}
