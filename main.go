package logger

import (
	"fmt"
)

var (
	Reset  = "\033[0m"
	Red    = "\033[0;31m"
	Green  = "\033[0;32m"
	Orange = "\033[0;33m"
	Blue   = "\033[0;34m"
)

// Error displays a message in a red color
// Works like Printf, so you can pass arguments, so they are added to the string
func Error(message string, args ...any) {
	fmt.Printf(Red+message+Reset+"\n", args...)
}

// Success displays a message in a green color
// Works like Printf, so you can pass arguments, so they are added to the string
func Success(message string, args ...any) {
	fmt.Printf(Green+message+Reset+"\n", args...)
}

// Warning displays a message in an orange color
// Works like Printf, so you can pass arguments, so they are added to the string
func Warning(message string, args ...any) {
	fmt.Printf(Orange+message+Reset+"\n", args...)
}

// Custom displays a simple message where you add the colors
// Works like Printf, so you can pass arguments, so they are added to the string
func Custom(message string, args ...any) {
	fmt.Printf(message+"\n", args...)
}
