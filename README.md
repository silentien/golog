# Silentien - Golog

A simple and efficient logging library for Go, designed to be lightweight, easy to use and configurable.

## Installation

```bash
go get github.com/silentien/golog
```

## Usage

`golog` exposes a simple interface for logging. You can create a logger instance and use it to log messages at different levels.

Example *app.go*:

```go
package main

import (
    "github.com/silentien/golog"
)

func main() {
    // Create a new logger instance
    logger := golog.NewLogger("my-app")

    // Log messages at different levels
    logger.Debug("This is a debug message")
    logger.Info("This is an info message")
    logger.Warn("This is a warning message")
    logger.Error("This is an error message")
}
```

## Configuration

You can configure the logs dinamically with the following environment variables:

- `DEBUG`: Specify the *namespace* for which you want to enable debug logging. Allowing wildcard (`*`). This way you can enable debug logging for specific parts of your application. For instance `my-app` or `my-app:*:component`
- `DEBUG_LEVEL`: Set the minimum log level to output. Possible values are:
    * `debug`
    * `info`
    * `warn`
    * `error`
- `DEBUG_COLOR`: Enable or disable colored output. Set the environment variable to enabled color output.

## Customization

You can customize the logger by setting options when creating a new logger instance. The available options include:
- `golog.WithWriter(io.Writer)`: Set the output writer for the logger, in case you want to log to a file or another destination.
- `golog.WithLoggerImpl(golog.LoggerImpl)`: Set a custom logger implementation if you want to change the format or behavior of the logs.
