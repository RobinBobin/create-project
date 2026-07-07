package expoapp

import (
	"github.com/robinbobin/create-project/utils"
)

func setPackageType() {
	json := utils.ReadJSON(utils.PACKAGE_JSON)

	json["type"] = "module"

	utils.WriteJSON(json, utils.PACKAGE_JSON)
}
