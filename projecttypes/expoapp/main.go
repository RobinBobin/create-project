package expoapp

import (
	"github.com/robinbobin/create-project/utils"
)

func Create() bool {
	defer utils.RecoverFromPanic()

	initPackage()
	// utils.AskSortJSON("app.json")

	return true
}
