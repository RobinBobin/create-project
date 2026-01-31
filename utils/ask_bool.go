package utils

import (
	"github.com/AlecAivazis/survey/v2"
)

func AskBool(message string) bool {
	prompt := &survey.Confirm{
		Default: true,
		Message: message,
	}

	var response bool

	PanicOnError(survey.AskOne(prompt, &response))

	return response
}
