package logger_test

import (
	"github.com/TomasBorquez/logger"
)

func ExampleInfo() {
	logger.Info("Just some information being displayed %s", "here")
	// Output: Just some information being displayed
}

func ExampleError() {
	logger.Error("An error occurred: %s", "file not found")
	// Output: An error occurred: file not found
}

func ExampleSuccess() {
	logger.Success("Operation completed: %d items processed", 100)
	// Output: Operation completed: 100 items processed
}

func ExampleWarning() {
	logger.Warning("Low disk space: %d%% remaining", 10)
	// Output: Low disk space: 10% remaining
}

func ExampleCustom() {
	logger.Custom(logger.Blue+"Processing: "+logger.Reset+"%s", "data.txt")
	// Output: Processing: data.txt
}
