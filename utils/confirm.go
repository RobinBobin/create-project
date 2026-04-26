package utils

import (
	"github.com/charmbracelet/huh"
	"github.com/robinbobin/create-project/options"
)

func Confirm(title string, value bool) bool {
	if !options.Options.ShouldUseDefaults {
		PanicOnError(
			huh.NewConfirm().
				Title(title).
				Value(&value).
				Run(),
		)
	}

	return value
}
