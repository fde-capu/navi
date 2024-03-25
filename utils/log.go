package utils

import (
	"fmt"
	"strings"
)

func LogSuccess(message string) {
	fmt.Println(strings.Trim(message, " \n\t"))
}

func LogError(message string) {
	fmt.Println(strings.Trim(message, " \n\t"))
}

func LogInfo(message string) {
	fmt.Println(strings.Trim(message, " \n\t"))
}

func LogExplanation(message string) {
	fmt.Println(strings.Trim(message, " \n\t"))
}
