package utils

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func AskSortJSON(name string) {
	if !AskBool(fmt.Sprintf("Would you like to sort '%v'", filepath.Base(name))) {
		return
	}

	fileInfo, err := os.Stat(name)

	PanicOnError(err)

	fileData, err := os.ReadFile(name)

	PanicOnError(err)

	decoder := json.NewDecoder(strings.NewReader(string(fileData)))
	decoder.UseNumber()

	var jsonData any

	PanicOnError(decoder.Decode(&jsonData))

	sortedBytes, err := json.MarshalIndent(jsonData, "", strings.Repeat(" ", 2))

	PanicOnError(err)

	PanicOnError(os.WriteFile(name, sortedBytes, fileInfo.Mode()))
}
