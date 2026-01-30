package utils

import (
	"encoding/json"
	"os"
	"strings"
)

func ReadJSON(jsonFile string) (jsonData any) {
	fileData, err := os.ReadFile(jsonFile)

	PanicOnError(err)

	decoder := json.NewDecoder(strings.NewReader(string(fileData)))
	decoder.UseNumber()

	PanicOnError(decoder.Decode(&jsonData))

	return jsonData
}
