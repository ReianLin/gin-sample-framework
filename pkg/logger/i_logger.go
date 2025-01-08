package logger

import "context"

var (
	globalLogger Logger
)

func SetGlobalLogger(logger Logger) {
	globalLogger = logger
}

func GetGlobalLogger() Logger {
	return globalLogger
}

type Fields map[string]interface{}

type Logger interface {
	WithField(key string, value interface{}) Logger
	WithFields(fields Fields) Logger
	WithContext(ctx context.Context) Logger
	Debug(args ...interface{})
	Debugf(format string, args ...interface{})
	Info(args ...interface{})
	Infof(format string, args ...interface{})
	Warn(args ...interface{})
	Warnf(format string, args ...interface{})
	Error(args ...interface{})
	Errorf(format string, args ...interface{})
	Panic(args ...interface{})
	Panicf(format string, args ...interface{})
	Fatal(args ...interface{})
	Fatalf(format string, args ...interface{})
}
