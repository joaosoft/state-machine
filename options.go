package logger

import (
	gowriter "github.com/joaosoft/go-writer/app"
	"io"
)

// LoggerOption ...
type LoggerOption func(log *Logger)

// Reconfigure ...
func (logger *Logger) Reconfigure(options ...LoggerOption) {
	for _, option := range options {
		option(logger)
	}
}

// WithWriter ...
func WithWriter(writer io.Writer) LoggerOption {
	return func(logger *Logger) {
		logger.writer = writer
	}
}

// WithSpecialWriter ...
func WithSpecialWriter(writer ISpecialWriter) LoggerOption {
	return func(logger *Logger) {
		logger.specialWriter = writer
	}
}

// WithLevel ...
func WithLevel(level Level) LoggerOption {
	return func(logger *Logger) {
		logger.level = level
	}
}

// WithFormatHandler ...
func WithFormatHandler(formatHandler gowriter.FormatHandler) LoggerOption {
	return func(logger *Logger) {
		logger.formatHandler = formatHandler
	}
}
