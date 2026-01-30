package utils

import (
	"fmt"
	"path/filepath"
)

func AskSortJSON(name string) {
	AskSortJSONInDir(name, "")
}

func AskSortJSONInDir(name string, dir string) {
	jsonFile := filepath.Join(dir, name)

	if !AskBool(fmt.Sprintf("Would you like to sort '%v'", jsonFile)) {
		return
	}

	var jsonData any

	ReadJSON(&jsonData, jsonFile)
	WriteJSON(jsonData, jsonFile)
}
