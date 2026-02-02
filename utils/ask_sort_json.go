package utils

import (
	"fmt"
)

func AskSortJSON(name string) {
	if !AskBool(fmt.Sprintf("Would you like to sort '%v'", name)) {
		return
	}

	WriteJSON(ReadJSON(name), name)
}
