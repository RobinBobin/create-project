package addition

import (
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
		{
			Fn:   addReactUse,
			Name: "react-use",
		},
		{
			Fn:   addTypeFest,
			Name: "type-fest",
		},
	}

	utils.BatchActions(actions, "Which of the following packages would you like to add?")
}
