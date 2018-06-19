package golog

import (
	"fmt"
	"os"

	"time"

	gowriter "github.com/joaosoft/go-writer/app"
	"net"
	"runtime"
	"strings"
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
	newLog := log.clone().WithPrefixes(prefixes).WithTags(tags).WithFields(fields)
	return newLog
}

func (log *Log) WithPrefixes(prefixes map[string]interface{}) ILog {
	newLog := log.clone()
	newLog.prefixes = prefixes
	return newLog
}

func (log *Log) WithTags(tags map[string]interface{}) ILog {
	newLog := log.clone()
	newLog.tags = tags
	return newLog
}

func (log *Log) WithFields(fields map[string]interface{}) ILog {
	newLog := log.clone()
	newLog.fields = fields
	return newLog
}

func (log *Log) WithPrefix(key string, value interface{}) ILog {
	newLog := log.clone()
	newLog.prefixes[key] = fmt.Sprintf("%s", value)
	return newLog
}

func (log *Log) WithTag(key string, value interface{}) ILog {
	newLog := log.clone()
	newLog.tags[key] = fmt.Sprintf("%s", value)
	return newLog
}

func (log *Log) WithField(key string, value interface{}) ILog {
	newLog := log.clone()
	newLog.fields[key] = fmt.Sprintf("%s", value)
	return newLog
}

// Clone ...
func (log *Log) clone() *Log {
	return &Log{
		level:         log.level,
		writer:        log.writer,
		formatHandler: log.formatHandler,
		specialWriter: log.specialWriter,
		tags:          log.tags,
		prefixes:      log.prefixes,
		fields:        log.fields,
	}
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

		case IP:
			addresses, _ := net.InterfaceAddrs()
			for _, a := range addresses {
				if ipNet, ok := a.(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
					if ipNet.IP.To4() != nil {
						value = ipNet.IP.String()
					}
				}
			}

		case TRACE:
			pc := make([]uintptr, 1)
			runtime.Callers(4, pc)
			function := runtime.FuncForPC(pc[0])
			file, line := function.FileLine(pc[0])
			info := strings.SplitN(function.Name(), ".", 2)

			value = struct {
				File     string `json:"file"`
				Line     int    `json:"line"`
				Package  string `json:"package"`
				Function string `json:"function"`
			}{
				File:     file,
				Line:     line,
				Package:  info[0],
				Function: info[1],
			}

		case FILE:
			pc := make([]uintptr, 1)
			runtime.Callers(4, pc)
			function := runtime.FuncForPC(pc[0])
			value, _ = function.FileLine(pc[0])

		case PACKAGE:
			pc := make([]uintptr, 1)
			runtime.Callers(4, pc)
			function := runtime.FuncForPC(pc[0])
			value = strings.SplitN(function.Name(), ".", 2)[0]

		case FUNCTION:
			pc := make([]uintptr, 1)
			runtime.Callers(4, pc)
			function := runtime.FuncForPC(pc[0])
			value = strings.SplitN(function.Name(), ".", 2)[1]
		}

		newPrefixes[key] = value
	}
	return newPrefixes
}
