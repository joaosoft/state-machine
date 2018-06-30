package logger

import (
	"io"

	"github.com/joaosoft/go-error/app"
	"github.com/joaosoft/go-writer/app"
)

type IAddition interface {
	ToError(err *error) IAddition
	ToErrorData(err *goerror.ErrorData) IAddition
}

type ISpecialWriter interface {
	SWrite(prefixes map[string]interface{}, tags map[string]interface{}, message interface{}, fields map[string]interface{}) (n int, err error)
}

type ILogger interface {
	SetLevel(level Level)

	With(prefixes, tags, fields map[string]interface{}) ILogger
	WithPrefixes(prefixes map[string]interface{}) ILogger
	WithTags(tags map[string]interface{}) ILogger
	WithFields(fields map[string]interface{}) ILogger

	WithPrefix(key string, value interface{}) ILogger
	WithTag(key string, value interface{}) ILogger
	WithField(key string, value interface{}) ILogger

	Debug(message interface{}) IAddition
	Info(message interface{}) IAddition
	Warn(message interface{}) IAddition
	Error(message interface{}) IAddition

	Debugf(format string, arguments ...interface{}) IAddition
	Infof(format string, arguments ...interface{}) IAddition
	Warnf(format string, arguments ...interface{}) IAddition
	Errorf(format string, arguments ...interface{}) IAddition

	Reconfigure(options ...LoggerOption)
}

// Logger ...
type Logger struct {
	level         Level
	writer        io.Writer
	specialWriter ISpecialWriter
	prefixes      map[string]interface{} `json:"prefixes"`
	tags          map[string]interface{} `json:"tags"`
	fields        map[string]interface{} `json:"fields"`
	formatHandler gowriter.FormatHandler
}