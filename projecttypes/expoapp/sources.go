package expoapp

// import (
// 	"os"

// 	"github.com/robinbobin/create-project/assets"
// 	"github.com/robinbobin/create-project/options"
// 	"github.com/robinbobin/create-project/utils"
// 	"github.com/robinbobin/create-project/utils/packagejson"
// )

// func copySources() {
// 	requiredPackages := []string{
// 		"mobx-state-tree",
// 		"radashi",
// 		"react-native-paper",
// 		"type-fest",
// 	}

// 	if !packagejson.AreInstalled(requiredPackages) {
// 		return
// 	}

// 	if !utils.Confirm("Would you like to copy some common sources?", true) {
// 		return
// 	}

// 	utils.PanicOnError(os.RemoveAll(utils.SRC))
// 	utils.PanicOnError(assets.CopyFS(utils.SRC, utils.SRC))

// 	options.Options.AreSourcesCopied = true
// }
