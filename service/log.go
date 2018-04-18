package golog

import (
	"fmt"
	"os"

	"time"

	"github.com/joaosoft/go-writer/service"
)

// NewLog ...
func NewLog(options ...GoLogOption) ILog {
	golog := &GoLog{
		writer:        os.Stdout,
		formatHandler: gowriter.JsonFormatHandler,
		level:         InfoLevel,
		prefixes:      make(map[string]interface{}),
		tags:          make(map[string]interface{}),
		fields:        make(map[string]interface{}),
	}
	golog.Reconfigure(options...)

	return golog
}

// NewLogDefault
func NewLogDefault(service string, level Level) ILog {
	return NewLog(
		WithLevel(level),
		WithFormatHandler(gowriter.JsonFormatHandler),
		WithWriter(os.Stdout)).
		With(
			map[string]interface{}{"level": LEVEL, "timestamp": TIMESTAMP},
			map[string]interface{}{"service": service},
			map[string]interface{}{})
}

func (log *GoLog) SetLevel(level Level) {
	log.level = level
}

func (log *GoLog) With(prefixes, tags, fields map[string]interface{}) ILog {
	log.WithPrefixes(prefixes)
	log.WithTags(tags)
	log.WithFields(fields)
	return log
}

func (log *GoLog) WithPrefixes(prefixes map[string]interface{}) ILog {
	log.prefixes = prefixes
	return log
}

func (log *GoLog) WithTags(tags map[string]interface{}) ILog {
	log.tags = tags
	return log
}

func (log *GoLog) WithFields(fields map[string]interface{}) ILog {
	log.fields = fields
	return log
}

func (log *GoLog) WithField(key string, value interface{}) ILog {
	log.fields[key] = fmt.Sprintf("%s", value)
	return log
}

func (log *GoLog) Debug(message interface{}) {
	log.writeLog(DebugLevel, fmt.Sprint(message))
}

func (log *GoLog) Info(message interface{}) {
	log.writeLog(InfoLevel, fmt.Sprint(message))
}

func (log *GoLog) Warn(message interface{}) {
	log.writeLog(WarnLevel, fmt.Sprint(message))
}

func (log *GoLog) Error(message interface{}) {
	log.writeLog(ErrorLevel, fmt.Sprint(message))
}

func (log *GoLog) Debugf(format string, arguments ...interface{}) {
	log.writeLog(DebugLevel, fmt.Sprintf(format, arguments...))
}

func (log *GoLog) Infof(format string, arguments ...interface{}) {
	log.writeLog(InfoLevel, fmt.Sprintf(format, arguments...))
}

func (log *GoLog) Warnf(format string, arguments ...interface{}) {
	log.writeLog(WarnLevel, fmt.Sprintf(format, arguments...))
}

func (log *GoLog) Errorf(format string, arguments ...interface{}) {
	log.writeLog(ErrorLevel, fmt.Sprintf(format, arguments...))
}

func (log *GoLog) writeLog(level Level, message interface{}) {
	if level > log.level {
		return
	}

	prefixes := addSystemInfo(level, log.prefixes)
	if log.specialWriter == nil {
		if bytes, err := log.formatHandler(prefixes, log.tags, fmt.Sprint(message), log.fields); err != nil {
			return
		} else {
			log.writer.Write(bytes)
		}
	} else {
		log.specialWriter.SWrite(prefixes, log.tags, message, log.fields)
	}
}

func addSystemInfo(level Level, prefixes map[string]interface{}) map[string]interface{} {
	newPrefixes := make(map[string]interface{})
	for key, value := range prefixes {
		switch value {
		case LEVEL:
			value = level.String()
		case TIMESTAMP:
			value = time.Now().Format("2006-01-02 15:04:05:06")
		case DATE:
			value = time.Now().Format("2006-01-02")
		case TIME:
			value = time.Now().Format("15:04:05:06")
		}

		newPrefixes[key] = value
	}
	return newPrefixes
}
