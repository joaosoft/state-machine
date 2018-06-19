package golog

var internalLog = NewLogEmpty(InfoLevel)

func SetLevel(level Level) {
	internalLog.SetLevel(level)
}

func With(prefixes, tags, fields map[string]interface{}) ILog {
	return internalLog.With(prefixes, tags, fields)
}

func WithPrefixes(prefixes map[string]interface{}) ILog {
	return internalLog.WithPrefixes(prefixes)
}

func WithTags(tags map[string]interface{}) ILog {
	return internalLog.WithTags(tags)
}

func WithFields(fields map[string]interface{}) ILog {
	return internalLog.WithFields(fields)
}

func WithField(key string, value interface{}) ILog {
	return internalLog.WithField(key, value)
}

func Debug(message interface{}) IAddition {
	return internalLog.Debug(message)
}

func Info(message interface{}) IAddition {
	return internalLog.Info(message)
}

func Warn(message interface{}) IAddition {
	return internalLog.Warn(message)
}

func Error(message interface{}) IAddition {
	return internalLog.Error(message)
}

func Debugf(format string, arguments ...interface{}) IAddition {
	return internalLog.Debugf(format, arguments)
}

func Infof(format string, arguments ...interface{}) IAddition {
	return internalLog.Infof(format, arguments)
}

func Warnf(format string, arguments ...interface{}) IAddition {
	return internalLog.Warnf(format, arguments)
}

func Errorf(format string, arguments ...interface{}) IAddition {
	return internalLog.Errorf(format, arguments)
}
