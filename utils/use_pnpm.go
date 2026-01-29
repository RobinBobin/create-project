package utils

import "os/exec"

func UsePNPM() {
	UsePNPMInDir("")
}

func UsePNPMInDir(dir string) {
	RunCmdWithPreRunner("corepack use pnpm@latest", func(cmd *exec.Cmd) {
		cmd.Dir = dir
	})
}
