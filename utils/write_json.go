package utils

import (
	"encoding/json"
	"os"
	"strings"
)

func WriteJSON(jsonData map[string]any, jsonFile string) {
	fileInfo, err := os.Stat(jsonFile)
	PanicOnError(err)

	sortedBytes, err := json.MarshalIndent(jsonData, "", strings.Repeat(" ", 2))
	PanicOnError(err)

	PanicOnError(os.WriteFile(jsonFile, sortedBytes, fileInfo.Mode()))
}
