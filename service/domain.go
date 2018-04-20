package golog

import (
	"io"
	"sync"

	"github.com/joaosoft/go-error/service"
	"github.com/joaosoft/go-writer/service"
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

	Debug(message interface{}) IAddition
	Info(message interface{}) IAddition
	Warn(message interface{}) IAddition
	Error(message interface{}) IAddition

	Debugf(format string, arguments ...interface{}) IAddition
	Infof(format string, arguments ...interface{}) IAddition
	Warnf(format string, arguments ...interface{}) IAddition
	Errorf(format string, arguments ...interface{}) IAddition
}

// GoLog ...
type GoLog struct {
	level         Level
	writer        io.Writer
	specialWriter ISpecialWriter
	prefixes      map[string]interface{} `json:"prefixes"`
	tags          map[string]interface{} `json:"tags"`
	fields        map[string]interface{} `json:"fields"`
	formatHandler gowriter.FormatHandler
	mux           *sync.Mutex
}
