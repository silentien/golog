package types

import "fmt"

// InvalidLogLevelError is an error type that represents an invalid log level.
type InvalidLogLevelError struct {
	error_numer int
}

// Error implements the error interface for InvalidLogLevelError.
func (e *InvalidLogLevelError) Error() string {
	return fmt.Sprintf("Invalid log level: %d", e.error_numer)
}

// LogLevel represents the different levels of logging severity.
type LogLevel int

const (
	// LogLevelDebug is the debug log level, used for detailed debugging information.
	LogLevelDebug LogLevel = iota
	// LogLevelInfo is the info log level, used for general information messages.
	LogLevelInfo
	// LogLevelWarn is the warn log level, used for warning messages that indicate potential issues.
	LogLevelWarn
	// LogLevelError is the error log level, used for error messages that indicate failures.
	LogLevelError
)

// String converts a LogLevel to its string representation.
// It returns an error if the LogLevel is invalid.
//
// Example:
//
//	level := types.LogLevelDebug
//	str, err := level.String()
//	if err != nil {
//		panic(err)
//	}
//	fmt.Println(str) // Output: DEBUG
func (ll LogLevel) String() (string, error) {
	switch ll {

	case LogLevelDebug:
		return "DEBUG", nil
	case LogLevelInfo:
		return "INFO", nil
	case LogLevelWarn:
		return "WARN", nil
	case LogLevelError:
		return "ERROR", nil
	default:
		return "", &InvalidLogLevelError{error_numer: int(ll)}
	}
}

// LogLevelFromString converts a string representation of a log level to its LogLevel type.
// It returns an error if the string does not match any valid log level.
//
// Example:
//
//	levelStr := "DEBUG"
//	level, err := types.LogLevelFromString(levelStr)
//	if err != nil {
//		panic(err)
//	}
func LogLevelFromString(levelStr string) (LogLevel, error) {
	switch levelStr {
	case "DEBUG":
		return LogLevelDebug, nil
	case "INFO":
		return LogLevelInfo, nil
	case "WARN":
		return LogLevelWarn, nil
	case "ERROR":
		return LogLevelError, nil
	default:
		return -1, &InvalidLogLevelError{error_numer: -1}
	}
}
