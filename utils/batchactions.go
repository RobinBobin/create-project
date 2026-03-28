package utils

import "github.com/charmbracelet/huh"

func BatchActions(actions []*Action[func()], title string) {
	PanicOnError(
		huh.NewMultiSelect[*Action[func()]]().
			Title(title).
			Options(huh.NewOptions(actions...)...).
			Value(&actions).
			Run(),
	)

	if len(actions) == 0 {
		return
	}

	for _, action := range actions {
		action.Fn()
	}
}
