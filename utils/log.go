package utils

import (
	"fmt"
	"strings"
)

func LogSuccess(message string) {
	fmt.Printf("%s\n", strings.Trim(message, " \n\t")) // Why is this line not outputting a newline at the end of the string?
}

func LogError(message string) {
	fmt.Printf("%s\n", strings.Trim(message, " \n\t"))
}

func LogInfo(message string) {
	fmt.Printf("%s\n", strings.Trim(message, " \n\t"))
}

func LogExplanation(message string) {
	fmt.Printf("%s\n", strings.Trim(message, " \n\t"))
}
