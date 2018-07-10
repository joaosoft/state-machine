package logger

var internalLogger = NewLoggerEmpty(InfoLevel)

func SetLevel(level Level) {
	internalLogger.SetLevel(level)
}

func With(prefixes, tags, fields map[string]interface{}) ILogger {
	return internalLogger.With(prefixes, tags, fields)
}

func WithPrefixes(prefixes map[string]interface{}) ILogger {
	return internalLogger.WithPrefixes(prefixes)
}

func WithTags(tags map[string]interface{}) ILogger {
	return internalLogger.WithTags(tags)
}

func WithFields(fields map[string]interface{}) ILogger {
	return internalLogger.WithFields(fields)
}

func WithField(key string, value interface{}) ILogger {
	return internalLogger.WithField(key, value)
}

func Debug(message interface{}) IAddition {
	return internalLogger.Debug(message)
}

func Info(message interface{}) IAddition {
	return internalLogger.Info(message)
}

func Warn(message interface{}) IAddition {
	return internalLogger.Warn(message)
}

func Error(message interface{}) IAddition {
	return internalLogger.Error(message)
}

func Debugf(format string, arguments ...interface{}) IAddition {
	return internalLogger.Debugf(format, arguments)
}

func Infof(format string, arguments ...interface{}) IAddition {
	return internalLogger.Infof(format, arguments)
}

func Warnf(format string, arguments ...interface{}) IAddition {
	return internalLogger.Warnf(format, arguments)
}

func Errorf(format string, arguments ...interface{}) IAddition {
	return internalLogger.Errorf(format, arguments)
}

func Reconfigure(options ...LoggerOption) {
	internalLogger.Reconfigure(options...)
}