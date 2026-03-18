package utils

import (
	"github.com/charmbracelet/huh"
)

func Confirm(title string, value bool) bool {
	PanicOnError(
		huh.NewConfirm().
			Title(title).
			Value(&value).
			Run(),
	)

	return value
}
