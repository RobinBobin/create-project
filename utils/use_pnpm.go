package utils

import (
	"os"
)

func UsePNPM() {
	PanicOnError(os.Remove("pnpm-lock.yaml"))

	RunCmd("corepack use pnpm@latest")
}
