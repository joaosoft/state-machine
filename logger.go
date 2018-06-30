package logger

import (
	"fmt"
	"os"
	"time"
	writer "github.com/joaosoft/writers"
	"net"
	"runtime"
	"runtime/debug"
	"strings"
)

var logger = NewLoggerEmpty(InfoLevel)

// NewLogger ...
func NewLogger(options ...LoggerOption) ILogger {
	logger := &Logger{
		writer:        os.Stdout,
		formatHandler: writer.JsonFormatHandler,
		level:         InfoLevel,
		prefixes:      make(map[string]interface{}),
		tags:          make(map[string]interface{}),
		fields:        make(map[string]interface{}),
	}
	logger.Reconfigure(options...)

	return logger
}

// NewLogDefault
func NewLogDefault(service string, level Level) ILogger {
	return NewLogger(
		WithLevel(level),
		WithFormatHandler(writer.JsonFormatHandler),
		WithWriter(os.Stdout)).
		With(
			map[string]interface{}{"level": LEVEL, "timestamp": TIMESTAMP, "date": DATE, "time": TIME, "ip": IP, "package": PACKAGE, "function": FUNCTION, "stack": STACK, "trace": TRACE},
			map[string]interface{}{"service": service},
			map[string]interface{}{})
}

// NewLoggerEmpty
func NewLoggerEmpty(level Level) ILogger {
	return NewLogger(
		WithLevel(level),
		WithFormatHandler(writer.JsonFormatHandler),
		WithWriter(os.Stdout)).
		WithPrefixes(map[string]interface{}{"level": LEVEL, "timestamp": TIMESTAMP, "date": DATE, "time": TIME, "ip": IP, "package": PACKAGE, "function": FUNCTION, "stack": STACK, "trace": TRACE})
}

func (logger *Logger) SetLevel(level Level) {
	logger.level = level
}

func (logger *Logger) With(prefixes, tags, fields map[string]interface{}) ILogger {
	newLog := logger.clone().WithPrefixes(prefixes).WithTags(tags).WithFields(fields)
	return newLog
}

func (logger *Logger) WithPrefixes(prefixes map[string]interface{}) ILogger {
	newLog := logger.clone()
	newLog.prefixes = prefixes
	return newLog
}

func (logger *Logger) WithTags(tags map[string]interface{}) ILogger {
	newLog := logger.clone()
	newLog.tags = tags
	return newLog
}

func (logger *Logger) WithFields(fields map[string]interface{}) ILogger {
	newLog := logger.clone()
	newLog.fields = fields
	return newLog
}

func (logger *Logger) WithPrefix(key string, value interface{}) ILogger {
	newLog := logger.clone()
	newLog.prefixes[key] = fmt.Sprintf("%s", value)
	return newLog
}

func (logger *Logger) WithTag(key string, value interface{}) ILogger {
	newLog := logger.clone()
	newLog.tags[key] = fmt.Sprintf("%s", value)
	return newLog
}

func (logger *Logger) WithField(key string, value interface{}) ILogger {
	newLog := logger.clone()
	newLog.fields[key] = fmt.Sprintf("%s", value)
	return newLog
}

// Clone ...
func (logger *Logger) clone() *Logger {
	return &Logger{
		level:         logger.level,
		writer:        logger.writer,
		formatHandler: logger.formatHandler,
		specialWriter: logger.specialWriter,
		tags:          logger.tags,
		prefixes:      logger.prefixes,
		fields:        logger.fields,
	}
}

func (logger *Logger) Debug(message interface{}) IAddition {
	msg := fmt.Sprint(message)
	logger.writeLog(DebugLevel, message)

	return NewAddition(msg)
}

func (logger *Logger) Info(message interface{}) IAddition {
	msg := fmt.Sprint(message)
	logger.writeLog(InfoLevel, msg)

	return NewAddition(msg)
}

func (logger *Logger) Warn(message interface{}) IAddition {
	msg := fmt.Sprint(message)
	logger.writeLog(WarnLevel, msg)

	return NewAddition(msg)
}

func (logger *Logger) Error(message interface{}) IAddition {
	msg := fmt.Sprint(message)
	logger.writeLog(ErrorLevel, msg)

	return NewAddition(msg)
}

func (logger *Logger) Debugf(format string, arguments ...interface{}) IAddition {
	msg := fmt.Sprintf(format, arguments...)
	logger.writeLog(DebugLevel, msg)

	return NewAddition(msg)
}

func (logger *Logger) Infof(format string, arguments ...interface{}) IAddition {
	msg := fmt.Sprintf(format, arguments...)
	logger.writeLog(InfoLevel, msg)

	return NewAddition(msg)
}

func (logger *Logger) Warnf(format string, arguments ...interface{}) IAddition {
	msg := fmt.Sprintf(format, arguments...)
	logger.writeLog(WarnLevel, msg)

	return NewAddition(msg)
}

func (logger *Logger) Errorf(format string, arguments ...interface{}) IAddition {
	msg := fmt.Sprintf(format, arguments...)
	logger.writeLog(ErrorLevel, msg)

	return NewAddition(msg)
}

func (logger *Logger) writeLog(level Level, message interface{}) {
	if level > logger.level {
		return
	}

	prefixes := addSystemInfo(level, logger.prefixes)
	if logger.specialWriter == nil {
		if bytes, err := logger.formatHandler(prefixes, logger.tags, message, logger.fields); err != nil {
			return
		} else {
			logger.writer.Write(bytes)
		}
	} else {
		logger.specialWriter.SWrite(prefixes, logger.tags, message, logger.fields)
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
			stack := string(debug.Stack())
			stack = stack[strings.Index(stack, function.Name()):]

			value = struct {
				File     string `json:"file"`
				Line     int    `json:"line"`
				Package  string `json:"package"`
				Function string `json:"function"`
				Stack    string `json:"stack"`
			}{
				File:     file,
				Line:     line,
				Package:  info[0],
				Function: info[1],
				Stack:    stack,
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

		case STACK:
			pc := make([]uintptr, 1)
			runtime.Callers(4, pc)
			function := runtime.FuncForPC(pc[0])
			stack := string(debug.Stack())
			value = stack[strings.Index(stack, function.Name()):]
		}

		newPrefixes[key] = value
	}
	return newPrefixes
}
