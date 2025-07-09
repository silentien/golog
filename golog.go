// Package golog implements extending and customizable logging functionalities.
package golog

import (
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/silentien/golog/colors"
	"github.com/silentien/golog/internal/types"
	"github.com/silentien/golog/request"

	"github.com/silentien/golog/loggers/text"
)

// InvalidNamespaceError is an error type that represents an invalid namespace.
type InvalidNamespaceError struct {
}

// Error implements the error interface for InvalidNamespaceError.
func (e *InvalidNamespaceError) Error() string {
	return "Invalid namespace: it must not be empty"
}

// LoggerImpl is an interface that defines the methods for a logger implementation.
type LoggerImpl interface {
	Log(req request.LogRequest)
}

// isEnabled checks if the given namespace is allowed based on the allowedNamespaces pattern.
//
// Example:
//
//	isEnabled("test", "test") // true
//	isEnabled("test", "test*") // true
//	isEnabled("test", "*") // true
//	isEnabled("test", "other") // false
//	isEnabled("test:subnamespace", "test:*") // true
//	isEnabled("test:subnamespace", "test:foo*") // false
func isEnabled(namespace string, allowedNamespaces string) (bool, error) {
	r := regexp.QuoteMeta(allowedNamespaces)
	r = strings.ReplaceAll(r, "\\*", ".*")
	return regexp.Match(r, []byte(namespace))
}

// Logger is the main struct that implements the logging functionalities.
type Logger struct {
	namespace  string
	enabled    bool
	writer     io.Writer
	loggerImpl LoggerImpl
	logLevel   types.LogLevel
	addColor   types.AddColorFunc
	lastCall   time.Time
}

// New creates a new Logger instance with the specified namespace and options.
//
// Basic example:
//
//	l, err := New(namespace)
//
// With custom writer:
//
//	l, err := New(namespace, WithWriter(os.Stderr))
func New(namespace string, options ...func(*Logger)) (*Logger, error) {

	if namespace == "" {
		err := &InvalidNamespaceError{}
		return nil, err
	}

	allowedNamespaces, exists := os.LookupEnv(types.DEBUG)

	if !exists || allowedNamespaces == "" {
		allowedNamespaces = "*"
	}

	levelStr, exists := os.LookupEnv(types.DEBUG_LEVEL)
	if !exists || levelStr == "" {
		levelStr = "WARN"
	}

	_, colorEnabled := os.LookupEnv(types.DEBUG_COLOR)

	var addColor types.AddColorFunc
	if colorEnabled {
		addColor = colors.RandomColor(namespace).Sprint
	} else {
		addColor = fmt.Sprint
	}

	ll, err := types.LogLevelFromString(levelStr)
	if err != nil {
		return nil, err
	}

	enabled, err := isEnabled(namespace, allowedNamespaces)

	if err != nil {
		return nil, err
	}

	l := &Logger{
		namespace:  namespace,
		writer:     os.Stdout,
		enabled:    enabled,
		logLevel:   ll,
		loggerImpl: &text.TextLogger{},
		addColor:   addColor,
	}
	for _, o := range options {
		o(l)
	}
	return l, nil
}

// WithWriter is an option function to set a custom writer for the logger.
// If the writer is nil, it defaults to io.Discard.
// This allows for flexible output handling, such as writing to a file or network stream.
// Example:
//
//	l, err := New("my-namespace", WithWriter(os.Stderr))
func WithWriter(writer io.Writer) func(*Logger) {
	return func(l *Logger) {
		if writer == nil {
			writer = io.Discard
		}
		l.writer = writer
	}
}

// WithLoggerImpl is an option function to set a custom logger implementation.
// This allows for custom logging behavior by providing a different implementation of the LoggerImpl interface.
// If not set, it defaults to a text logger.
//
// Example:
//
//	l, err := New("my-namespace", WithLoggerImpl(&myCustomLogger{}))
func WithLoggerImpl(loggerImpl LoggerImpl) func(*Logger) {
	return func(l *Logger) {
		l.loggerImpl = loggerImpl
	}
}

// getNamespace returns the namespace of the logger.
func (l *Logger) getNamespace() string {
	return l.namespace
}

// New creates a new Logger instance with a new namespace based on the current logger's namespace.
// This allows for creating child loggers with a specific sub-namespace.
//
// Example:
//
//	childLogger, err := logger.New("subnamespace")
func (l *Logger) New(namespace string) (*Logger, error) {
	if namespace == "" {
		err := &InvalidNamespaceError{}
		return nil, err
	}
	return New(fmt.Sprintf("%s:%s", l.getNamespace(), namespace),
		WithWriter(l.writer),
		WithLoggerImpl(l.loggerImpl),
	)

}

// log is a private method that handles the actual logging and delegates the
// log request to the logger implementation.
//
// Example:
//
//	l.log(types.LogLevelDebug, "This is a debug message")
func (l *Logger) log(ll types.LogLevel, message string) {
	if l.enabled && l.logLevel <= ll {

		if l.lastCall.IsZero() {
			l.lastCall = time.Now()
		}

		l.loggerImpl.Log(request.LogRequest{
			Level:     ll,
			Namespace: l.namespace,
			Message:   message,
			Writer:    l.writer,
			Delay:     time.Since(l.lastCall),
			AddColor:  l.addColor,
		})
		l.lastCall = time.Now()
	}
}

// Debug logs a message at the debug level.
//
// Example:
//
//	l.Debug("This is a debug message")
func (l *Logger) Debug(message string) {
	l.log(types.LogLevelDebug, message)
}

// Info logs a message at the info level.
//
// Example:
//
//	l.Info("This is an info message")
func (l *Logger) Info(message string) {
	l.log(types.LogLevelInfo, message)
}

// Warn logs a message at the warning level.
//
// Example:
//
//	l.Warn("This is a warning message")
func (l *Logger) Warn(message string) {
	l.log(types.LogLevelWarn, message)
}

// Error logs a message at the error level.
//
// Example:
//
//	l.Error("This is an error message")
func (l *Logger) Error(message string) {
	l.log(types.LogLevelError, message)
}
