package text

import (
	"fmt"

	"github.com/silentien/golog/request"
)

// TextLogger is a simple logger implementation.
type TextLogger struct{}

// Function Log implements the LoggerImpl interface for TextLogger.
func (tl *TextLogger) Log(lr request.LogRequest) {
	levelStr, _ := lr.Level.String()
	fmt.Fprintf(lr.Writer, "%s [%s] %s %s", lr.AddColor(lr.Namespace), levelStr, lr.Message, lr.AddColor(lr.Delay.Milliseconds(), "ms"))
}
