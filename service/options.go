package golog

import (
	"io"

	"github.com/joaosoft/go-writer/service"
)

// logOption ...
type logOption func(log *Log)

// Reconfigure ...
func (log *Log) Reconfigure(options ...logOption) {
	for _, option := range options {
		option(log)
	}
}

// WithWriter ...
func WithWriter(writer io.Writer) logOption {
	return func(log *Log) {
		log.writer = writer
	}
}

// WithSpecialWriter ...
func WithSpecialWriter(writer ISpecialWriter) logOption {
	return func(log *Log) {
		log.specialWriter = writer
	}
}

// WithLevel ...
func WithLevel(level Level) logOption {
	return func(log *Log) {
		log.level = level
	}
}

// WithFormatHandler ...
func WithFormatHandler(formatHandler gowriter.FormatHandler) logOption {
	return func(log *Log) {
		log.formatHandler = formatHandler
	}
}
