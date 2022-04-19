package nbf

import (
	"fmt"
	"io"
	"os"
)

type Logger struct{
	stream io.Writer
}

type LoggerOption func(*Logger)

func WithStream(stream io.Writer) LoggerOption {
	return func(l *Logger) {
		l.stream = stream
	}
}

func NewLogger(opts ...LoggerOption) *Logger {
	logger := &Logger{
		stream: os.Stdout,
	}
	for _, opt := range opts {
		opt(logger)
	}

	return logger
}

func (l *Logger) debug(msg string) {
	fmt.Fprintf(l.stream, "[debug] %s", msg)
}

func (l *Logger) fatal(msg string) {
	fmt.Fprintf(l.stream, "[fatal] %s", msg)
	os.Exit(1)
}