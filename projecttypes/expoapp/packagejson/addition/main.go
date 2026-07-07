package addition

import (
	"slices"
	"strings"

	"github.com/robinbobin/create-project/projecttypes/expoapp/packagejson/addition/rnpaper"
	"github.com/robinbobin/create-project/utils"
	"github.com/robinbobin/create-project/utils/packagejson"
)

func Run() {
	actions := []*utils.BatchAction{
		{
			Fn:   addMST,
			Name: "MST tooling",
		},
		{
			Fn:   addPrettier,
			Name: "Prettier",
		},
		{
			Fn:   addRadashi,
			Name: "Radashi",
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

	if !packagejson.IsInstalled(utils.ESLINT) {
		actions = append(actions, &utils.BatchAction{
			Fn:   addESLint,
			Name: "ESLint",
		})
	}

	slices.SortFunc(actions, func(a, b *utils.BatchAction) int {
		return strings.Compare(a.Name, b.Name)
	})

	utils.BatchActions(actions, "Which of the following packages would you like to add?")
}
