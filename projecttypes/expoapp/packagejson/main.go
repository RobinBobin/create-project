package packagejson

import "github.com/robinbobin/create-project/utils"

func Lint() {
	utils.AskSortJSON("package.json")

	deletePackages()
}
