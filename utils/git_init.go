package utils

import (
	"errors"
	"os"
)

func GitInit() {
	_, err := os.Stat(".git")

	if errors.Is(err, os.ErrExist) {
		return
	}

	if errors.Is(err, os.ErrNotExist) {
		RunCmd("git init")

		return
	}

	PanicOnError(err)
}
