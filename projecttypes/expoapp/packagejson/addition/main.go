package addition

import (
	"github.com/charmbracelet/huh"
	"github.com/robinbobin/create-project/projecttypes/expoapp/packagejson/addition/rnpaper"
	"github.com/robinbobin/create-project/utils"
)

func Run() {
	actions := []*utils.Action[func()]{
		{
			Fn:   addMST,
			Name: "MST tooling",
		},
		{
			Fn:   addPrettier,
			Name: "Prettier",
		},
		{
			Fn:   rnpaper.Add,
			Name: "React Native Paper",
		},
	}

	utils.PanicOnError(
		huh.NewMultiSelect[*utils.Action[func()]]().
			Title("Which of the following packages would you like to add?").
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
