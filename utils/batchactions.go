package utils

import "github.com/charmbracelet/huh"

type batchActionFunc = func()

type batchActionable interface {
	comparable
	GetFn() batchActionFunc
	GetName() string
}

type BatchAction = Action[batchActionFunc]

func BatchActions[A batchActionable](actions []A, title string) {
	PanicOnError(
		huh.NewMultiSelect[A]().
			Title(title).
			Options(huh.NewOptions(actions...)...).
			Value(&actions).
			Run(),
	)

	if len(actions) == 0 {
		return
	}

	for _, action := range actions {
		action.GetFn()()
	}
}
