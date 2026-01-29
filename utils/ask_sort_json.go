package utils

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func AskSortJSON(name string) {
	AskSortJSONInDir(name, "")
}

func AskSortJSONInDir(name string, dir string) {
	fmt.Println()

	var jsonFile string

	if len(dir) != 0 {
		jsonFile = fmt.Sprintf("%v", filepath.Join(dir, name))
	} else {
		jsonFile = name
	}

	if !AskBool(fmt.Sprintf("Would you like to sort '%v'", jsonFile)) {
		return
	}

	fileInfo, err := os.Stat(jsonFile)

	PanicOnError(err)

	fileData, err := os.ReadFile(jsonFile)

	PanicOnError(err)

	decoder := json.NewDecoder(strings.NewReader(string(fileData)))
	decoder.UseNumber()

	var jsonData any

	PanicOnError(decoder.Decode(&jsonData))

	sortedBytes, err := json.MarshalIndent(jsonData, "", strings.Repeat(" ", 2))

	PanicOnError(err)

	PanicOnError(os.WriteFile(jsonFile, sortedBytes, fileInfo.Mode()))
}
