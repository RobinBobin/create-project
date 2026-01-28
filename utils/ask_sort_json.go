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

	VerifyOK(err)

	fileData, err := os.ReadFile(name)

	VerifyOK(err)

	decoder := json.NewDecoder(strings.NewReader(string(fileData)))
	decoder.UseNumber()

	var jsonData any

	VerifyOK(decoder.Decode(&jsonData))

	sortedBytes, err := json.MarshalIndent(jsonData, "", strings.Repeat(" ", 2))

	VerifyOK(err)

	VerifyOK(os.WriteFile(name, sortedBytes, fileInfo.Mode()))
}
