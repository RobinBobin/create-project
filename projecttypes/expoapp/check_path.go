package expoapp

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/robinbobin/create-project/utils"
)

func checkPathIsCorrect(appName string) (appPath string) {
	wd, err := os.Getwd()
	utils.PanicOnError(err)

	src := filepath.Join(wd, appName)
	appPath = src

	if appName != filepath.Base(wd) {
		return
	}

	fmt.Println()

	if utils.AskBool(fmt.Sprintf("Is it the desired path: %v", src)) {
		return
	}

	dst := filepath.Join(filepath.Dir(wd), appName)

	if !utils.AskBool(fmt.Sprintf("Should it be: %v", dst)) {
		return
	}

	entries, err := os.ReadDir(src)
	utils.PanicOnError(err)

	for _, entry := range entries {
		utils.PanicOnError(os.Rename(
			filepath.Join(src, entry.Name()),
			filepath.Join(dst, entry.Name()),
		))
	}

	os.Remove(src)

	return dst
}
