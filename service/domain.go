package golog

import (
	"io"
	"sync"

	"github.com/joaosoft/go-writer/service"
)

type ISpecialWriter interface {
	SWrite(prefixes map[string]interface{}, tags map[string]interface{}, message interface{}, fields map[string]interface{}) (n int, err error)
}

type ILog interface {
	SetLevel(level Level)

	With(prefixes, tags, fields map[string]interface{}) ILog
	WithPrefixes(prefixes map[string]interface{}) ILog
	WithTags(tags map[string]interface{}) ILog
	WithFields(fields map[string]interface{}) ILog

	Debug(message interface{})
	Info(message interface{})
	Warn(message interface{})
	Error(message interface{})

	Debugf(format string, arguments ...interface{})
	Infof(format string, arguments ...interface{})
	Warnf(format string, arguments ...interface{})
	Errorf(format string, arguments ...interface{})
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
