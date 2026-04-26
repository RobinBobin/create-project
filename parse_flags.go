package main

import (
	"fmt"
	"strings"

	"github.com/robinbobin/create-project/options"
	"github.com/spf13/pflag"
)

type projectType = string

const project_type_expo = "expo"
const project_type_npm = "npm"

type flags struct {
	projectType *projectType
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
			"project-type",
			"",
			fmt.Sprintf("Project type (%v).", projectTypes),
		),
	}

	shouldUseDefaults := pflag.Bool(
		"use-defaults",
		true,
		"Use defaults for prompts.",
	)

	pflag.Parse()

	options.Options.ShouldUseDefaults = *shouldUseDefaults

	return flags
}
