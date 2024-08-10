package logger

import (
	"fmt"
)

var (
	Reset    = "\033[0m"
	Gray     = "\033[0;37m"
	Red      = "\033[0;31m"
	Green    = "\033[0;32m"
	Orange   = "\033[0;33m"
	Blue     = "\033[0;34m"
	Purple   = "\033[0;35m"
	Cyan     = "\033[0;36m"
	White    = "\033[0;37m"
	BgGray   = "\033[1;30m"
	BgRed    = "\033[1;31m"
	BgGreen  = "\033[1;32m"
	BgOrange = "\033[1;33m"
	BgBlue   = "\033[1;34m"
	BgPurple = "\033[1;35m"
	BgCyan   = "\033[1;36m"
	BgWhite  = "\033[1;37m"
)

// Info displays a message in a gray color
// Works like Printf, so you can pass arguments, so they are added to the string
func Info(message string, args ...any) {
	fmt.Printf(Gray+message+Reset+"\n", args...)
}

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
