package logger

import (
	"io"
	"os"
	"time"

	"github.com/rs/zerolog"
)

// New creates a new zerolog.Logger configured based on the log level string.
func New(level string) zerolog.Logger {
	lvl, err := zerolog.ParseLevel(level)
	if err != nil {
		lvl = zerolog.InfoLevel
	}

	var output io.Writer = zerolog.ConsoleWriter{
		Out:        os.Stdout,
		TimeFormat: time.RFC3339,
	}

	return zerolog.New(output).
		Level(lvl).
		With().
		Timestamp().
		Caller().
		Logger()
}
