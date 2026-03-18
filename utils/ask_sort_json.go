package utils

import (
	"fmt"
)

func AskSortJSON(name string) {
	if Confirm(fmt.Sprintf("Would you like to sort '%v'", name), true) {
		WriteJSON(ReadJSON(name), name)
	}
}
