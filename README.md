# Simple Color Logger for Go

This is a lightweight, zero-dependency color logging library for Go. It provides an easy way to add color-coded log messages to your Go applications.

## Features

- Color-coded log messages (Red, Green, Orange, Blue)
- Zero external dependencies
- Single file implementation
- Printf-like functionality for easy formatting

## Installation

To use this logger in your Go project, simply copy the `logger.go` file into your project directory.

## Usage

Import the package in your Go file and use as such:

```go
import "github.com/TomasBorquez/logger"

func main() {
    logger.Info("Some message")
    logger.Error("An error occurred: %v", err)
    logger.Success("Operation completed successfully")
    logger.Warning("Warning: %s", warningMessage)
    logger.Custom(logger.Blue + "Custom colored message" + logger.Reset)
}
```

## Available Functions

`logger.Info(message string, args ...any)`: Displays a message in gray

`logger.Error(message string, args ...any)`: Displays a message in red

`logger.Success(message string, args ...any)`: Displays a message in green

`logger.Warning(message string, args ...any)`: Displays a message in orange

`logger.Custom(message string, args ...any)`: Displays a message with custom formatting

## Color Constants
The following color constants are available for use with the Custom function:

- **Reset**: Resets the color
- **Red**: Red color
- **Green**: Green color
- **Orange**: Orange color
- **Blue**: Blue color


## License

MIT