package utils

import (
	"encoding/json"
	"os"
	"strings"
)

func ReadJSON(jsonFile string) map[string]any {
	fileData, err := os.ReadFile(jsonFile)

	PanicOnError(err)

	decoder := json.NewDecoder(strings.NewReader(string(fileData)))
	decoder.UseNumber()

	jsonData := make(map[string]any)

	PanicOnError(decoder.Decode(&jsonData))

	return jsonData
}
