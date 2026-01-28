package utils

const (
	no  = "No"
	yes = "Yes"
)

var options = []string{yes, no}

func AskBool(message string) bool {
	return AskOneWithMessage(message, options) == yes
}
