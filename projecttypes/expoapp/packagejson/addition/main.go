package addition

import (
	"slices"
	"strings"

	"github.com/robinbobin/create-project/utils"
	"github.com/robinbobin/create-project/utils/packagejson"
)

func Run() {
	actions := []*utils.BatchAction{}

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
