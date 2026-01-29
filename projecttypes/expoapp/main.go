package expoapp

import (
	"github.com/robinbobin/create-project/utils"
)

func Create() bool {
	defer utils.RecoverFromPanic()

	appName := createApp()

	utils.AskSortJSONInDir("app.json", appName)

	return true
}
