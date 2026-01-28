package utils

import (
	"github.com/AlecAivazis/survey/v2"
)

func AskOne(options []string) string {
	return AskOneWithMessage("What would you like:", options)
}

func AskOneWithMessage(message string, options []string) string {
	prompt := &survey.Select{
		Message: message,
		Options: options,
	}

	result := ""

	err := survey.AskOne(prompt, &result)

	VerifyOK(err)

	return result
}
