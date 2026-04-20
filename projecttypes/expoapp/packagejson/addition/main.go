package addition

import (
	"github.com/robinbobin/create-project/projecttypes/expoapp/packagejson/addition/rnpaper"
	"github.com/robinbobin/create-project/utils"
)

type action struct {
	utils.BatchAction
}

func Run() {
	actions := []*action{
		{
			utils.BatchAction{
				Fn:   addMST,
				Name: "MST tooling",
			},
		},
		{
			utils.BatchAction{
				Fn:   addPrettier,
				Name: "Prettier",
			},
		},
		{
			utils.BatchAction{
				Fn:   addRadashi,
				Name: "Radashi",
			},
		},
		{
			utils.BatchAction{
				Fn:   rnpaper.Add,
				Name: "React Native Paper",
			},
		},
		{
			utils.BatchAction{
				Fn:   addReactUse,
				Name: "react-use",
			},
		},
		{
			utils.BatchAction{
				Fn:   addTypeFest,
				Name: "type-fest",
			},
		},
	}

	utils.BatchActions(actions, "Which of the following packages would you like to add?")
}
