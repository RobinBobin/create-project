package main

import (
	"fmt"
	"strings"

	"github.com/spf13/pflag"
)

type projectType = string

const project_type_expo = "expo"
const project_type_npm = "npm"

type flags struct {
	projectType *projectType
	// shouldUseDefaults *bool
}

func parseFlags() flags {
	projectTypes := strings.Join(
		[]string{
			project_type_expo,
			project_type_npm,
		},
		" / ")

	flags := flags{
		projectType: pflag.String(
			"projectType",
			"",
			fmt.Sprintf("Project type (%v).", projectTypes),
		),
		// shouldUseDefaults: pflag.Bool(
		// 	"shouldUseDefaults",
		// 	true,
		// 	"Use defaults for prompts.",
		// ),
	}

	pflag.Parse()

	return flags
}
