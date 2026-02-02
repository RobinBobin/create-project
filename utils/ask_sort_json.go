package utils

import (
	"fmt"

	"github.com/charmbracelet/huh"
)

func AskSortJSON(name string) {
	shouldSort := true

	PanicOnError(
		huh.NewConfirm().
			Title(fmt.Sprintf("Would you like to sort '%v'", name)).
			Value(&shouldSort).
			Run(),
	)

	if !shouldSort {
		return
	}

	WriteJSON(ReadJSON(name), name)
}
