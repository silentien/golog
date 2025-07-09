package request

import (
	"io"
	"time"

	"github.com/silentien/golog/internal/types"
)

// LogRequest represents the request with the data available to process by the logger.
type LogRequest struct {
	Level     types.LogLevel
	Namespace string
	Message   string
	Writer    io.Writer
	Delay     time.Duration
	AddColor  types.AddColorFunc
}
