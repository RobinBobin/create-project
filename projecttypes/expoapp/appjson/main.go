package appjson

import (
	"github.com/robinbobin/create-project/projecttypes/expoapp/appjson/missingplugins"
	"github.com/robinbobin/create-project/utils"
)

func Lint() {
	utils.AskSortJSON(utils.APP_JSON)

	missingplugins.AddMissingPlugins()
}
