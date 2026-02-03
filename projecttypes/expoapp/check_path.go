package expoapp

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/charmbracelet/huh"
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

	value := false

	utils.PanicOnError(
		huh.NewConfirm().
			Title(fmt.Sprintf("Is it the desired path: %v", src)).
			Value(&value).
			Run(),
	)

	if value {
		return true
	}

	value = true

	utils.PanicOnError(
		huh.NewConfirm().
			Title(fmt.Sprintf("Should it be: %v", wd)).
			Value(&value).
			Run(),
	)

	if !value {
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
