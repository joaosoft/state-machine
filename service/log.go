package golog

import (
	"fmt"
	"os"

	"time"

	"github.com/joaosoft/go-writer/service"
)

var log = NewLogEmpty(InfoLevel)

// NewLog ...
func NewLog(options ...logOption) ILog {
	golog := &Log{
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

// NewLogEmpty
func NewLogEmpty(level Level) ILog {
	return NewLog(
		WithLevel(level),
		WithFormatHandler(gowriter.JsonFormatHandler),
		WithWriter(os.Stdout)).
		WithPrefixes(map[string]interface{}{"level": LEVEL, "timestamp": TIMESTAMP})
}

func (log *Log) SetLevel(level Level) {
	log.level = level
}

func (log *Log) With(prefixes, tags, fields map[string]interface{}) ILog {
	log.WithPrefixes(prefixes)
	log.WithTags(tags)
	log.WithFields(fields)
	return log
}

func (log *Log) WithPrefixes(prefixes map[string]interface{}) ILog {
	log.prefixes = prefixes
	return log
}

func (log *Log) WithTags(tags map[string]interface{}) ILog {
	log.tags = tags
	return log
}

func (log *Log) WithFields(fields map[string]interface{}) ILog {
	log.fields = fields
	return log
}

func (log *Log) WithPrefix(key string, value interface{}) ILog {
	log.prefixes[key] = fmt.Sprintf("%s", value)
	return log
}

func (log *Log) WithTag(key string, value interface{}) ILog {
	log.tags[key] = fmt.Sprintf("%s", value)
	return log
}

func (log *Log) WithField(key string, value interface{}) ILog {
	log.fields[key] = fmt.Sprintf("%s", value)
	return log
}

func (log *Log) Debug(message interface{}) IAddition {
	msg := fmt.Sprint(message)
	log.writeLog(DebugLevel, message)

	return newAddition(msg)
}

func (log *Log) Info(message interface{}) IAddition {
	msg := fmt.Sprint(message)
	log.writeLog(InfoLevel, msg)

	return newAddition(msg)
}

func (log *Log) Warn(message interface{}) IAddition {
	msg := fmt.Sprint(message)
	log.writeLog(WarnLevel, msg)

	return newAddition(msg)
}

func (log *Log) Error(message interface{}) IAddition {
	msg := fmt.Sprint(message)
	log.writeLog(ErrorLevel, msg)

	return newAddition(msg)
}

func (log *Log) Debugf(format string, arguments ...interface{}) IAddition {
	msg := fmt.Sprintf(format, arguments...)
	log.writeLog(DebugLevel, msg)

	return newAddition(msg)
}

func (log *Log) Infof(format string, arguments ...interface{}) IAddition {
	msg := fmt.Sprintf(format, arguments...)
	log.writeLog(InfoLevel, msg)

	return newAddition(msg)
}

func (log *Log) Warnf(format string, arguments ...interface{}) IAddition {
	msg := fmt.Sprintf(format, arguments...)
	log.writeLog(WarnLevel, msg)

	return newAddition(msg)
}

func (log *Log) Errorf(format string, arguments ...interface{}) IAddition {
	msg := fmt.Sprintf(format, arguments...)
	log.writeLog(ErrorLevel, msg)

	return newAddition(msg)
}

func (log *Log) writeLog(level Level, message interface{}) {
	if level > log.level {
		return
	}

	prefixes := addSystemInfo(level, log.prefixes)
	if log.specialWriter == nil {
		if bytes, err := log.formatHandler(prefixes, log.tags, message, log.fields); err != nil {
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
