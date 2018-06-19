package golog

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

type ILog interface {
	SetLevel(level Level)

	With(prefixes, tags, fields map[string]interface{}) ILog
	WithPrefixes(prefixes map[string]interface{}) ILog
	WithTags(tags map[string]interface{}) ILog
	WithFields(fields map[string]interface{}) ILog

	WithPrefix(key string, value interface{}) ILog
	WithTag(key string, value interface{}) ILog
	WithField(key string, value interface{}) ILog

	Debug(message interface{}) IAddition
	Info(message interface{}) IAddition
	Warn(message interface{}) IAddition
	Error(message interface{}) IAddition

	Debugf(format string, arguments ...interface{}) IAddition
	Infof(format string, arguments ...interface{}) IAddition
	Warnf(format string, arguments ...interface{}) IAddition
	Errorf(format string, arguments ...interface{}) IAddition

	Reconfigure(options ...logOption)
}

// Log ...
type Log struct {
	level         Level
	writer        io.Writer
	specialWriter ISpecialWriter
	prefixes      map[string]interface{} `json:"prefixes"`
	tags          map[string]interface{} `json:"tags"`
	fields        map[string]interface{} `json:"fields"`
	formatHandler gowriter.FormatHandler
}
