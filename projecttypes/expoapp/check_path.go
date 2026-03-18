package expoapp

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/robinbobin/create-project/utils"
)

func checkPathIsCorrect(appName string) bool {
	_, err := os.Stat(appName)

	if errors.Is(err, os.ErrNotExist) {
		return false
	}

	utils.PanicOnError(err)

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
		return true
	}

	if utils.Confirm(fmt.Sprintf("Is it the desired path: %v", src), false) {
		return true
	}

	if !utils.Confirm(fmt.Sprintf("Should it be: %v", wd), true) {
		return true
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

	return true
}
