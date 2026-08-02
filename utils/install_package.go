package utils

import (
	"strings"

	"github.com/robinbobin/create-project/options"
)

func InstallPackage(name string, isDevDependency ...bool) {
	isSaveDev := (len(isDevDependency) == 1) && isDevDependency[0]

	sb := strings.Builder{}

	sb.WriteString("pnpm i")

	if isSaveDev {
		sb.WriteString(" --save-dev")
	}

	sb.WriteRune(' ')
	sb.WriteString(name)

	cmd := sb.String()

	if options.Options.CanInstallPackages {
		RunCmd(cmd)
	} else {
		options.Options.Hints.Add(cmd)
	}
}
