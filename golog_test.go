package golog

import (
	"bytes"
	"fmt"
	"io"
	"testing"

	"github.com/silentien/golog/internal/types"
)

func TestIsEnabled(t *testing.T) {
	tests := []struct {
		namespace         string
		allowedNamespaces string
		expected          bool
	}{
		{"test", "test", true},
		{"test", "test*", true},
		{"test", "*", true},
		{"test", "other", false},
		{"test:subnamespace", "test:*", true},
		{"test:subnamespace", "test:sub*", true},
		{"test:subnamespace", "test:foo*", false},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("isEnabled(%s, %s)", tt.namespace, tt.allowedNamespaces), func(t *testing.T) {
			result, err := isEnabled(tt.namespace, tt.allowedNamespaces)
			if err != nil {
				t.Errorf("Unexpected error: %v", err)
			}
			if result != tt.expected {
				t.Errorf("[%d] Expected %v, got %v", i, tt.expected, result)
			}
		})
	}
}

func TestLoggerWithWriter(t *testing.T) {

	namespace := "*"
	debugLevel := "DEBUG"
	logMessage := "Hello"

	t.Setenv(types.DEBUG, "*")
	t.Setenv(types.DEBUG_LEVEL, debugLevel)

	var b bytes.Buffer

	writer := io.Writer(&b)

	l, err := New(namespace,
		WithWriter(writer),
	)

	// should not error
	if err != nil {
		t.Error("Should not throw any errors creating a default logger")
	}

	l.Debug(logMessage)

	if b.Len() == 0 {
		t.Error("Should write to the writer")
	}

	expectedLog := fmt.Sprintf("%s [%s] %s 0ms", namespace, debugLevel, logMessage)

	if b.String() != expectedLog {
		t.Errorf("Expected '%s', got '%s'", expectedLog, b.String())
	}

	b.Reset()

	subnamespace := "subnamespace"

	subLogger, err := l.New(subnamespace)
	if err != nil {
		t.Errorf("Expected no error creating sublogger, got %v", err)
	}

	subLogger.Debug(logMessage)
	if b.Len() == 0 {
		t.Error("Sublogger should also write to the writer")
	}

	expectedSubLog := fmt.Sprintf("%s:%s [%s] %s 0ms", namespace, subnamespace, debugLevel, logMessage)

	if b.String() != expectedSubLog {
		t.Errorf("Expected '%s', got '%s'", expectedSubLog, b.String())
	}
}
