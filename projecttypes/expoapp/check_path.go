package expoapp

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/robinbobin/create-project/utils"
)

func checkPathIsCorrect(appName string) {
	wd, err := os.Getwd()
	utils.PanicOnError(err)

	src := filepath.Join(wd, appName)

	shouldChdir := true

	defer func() {
		if shouldChdir {
			utils.PanicOnError(os.Chdir(src))
		}
	}()

	if appName != filepath.Base(wd) {
		return
	}

	fmt.Println()

	if utils.AskBool(fmt.Sprintf("Is it the desired path: %v", src)) {
		return
	}

	if !utils.AskBool(fmt.Sprintf("Should it be: %v", wd)) {
		return
	}

	shouldChdir = false

	entries, err := os.ReadDir(src)
	utils.PanicOnError(err)

	for _, entry := range entries {
		utils.PanicOnError(os.Rename(
			filepath.Join(src, entry.Name()),
			filepath.Join(wd, entry.Name()),
		))
	}

	utils.PanicOnError(os.Remove(src))
}
