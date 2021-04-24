package log

import "sync"

var (
	_mu     sync.RWMutex
	_global = NewSimple()
)

// ReplaceGlobals replaces the global Logger, and returns a
// function to restore the original values. It's safe for concurrent use.
func ReplaceGlobals(logger Logger) func() {
	_mu.Lock()
	prev := _global
	_global = logger
	_mu.Unlock()
	return func() { ReplaceGlobals(prev) }
}

func globalLogger() Logger {
	_mu.RLock()
	l := _global
	_mu.RUnlock()
	return l
}

// Debug log message
func Debug(msg ...interface{}) {
	globalLogger().Debug(msg...)
}

// Info log message
func Info(msg ...interface{}) {
	globalLogger().Info(msg...)
}

// Warn log message
func Warn(msg ...interface{}) {
	globalLogger().Warn(msg...)
}

// Error log message
func Error(msg ...interface{}) {
	globalLogger().Error(msg...)
}

// Fatal log message
func Fatal(msg ...interface{}) {
	globalLogger().Fatal(msg...)
}

// Panic log message
func Panic(msg ...interface{}) {
	globalLogger().Panic(msg...)
}

// Debugf log message with formatting
func Debugf(format string, args ...interface{}) {
	globalLogger().Debugf(format, args...)
}

// Infof log message with formatting
func Infof(format string, args ...interface{}) {
	globalLogger().Infof(format, args...)
}

// Warnf log message with formatting
func Warnf(format string, args ...interface{}) {
	globalLogger().Warnf(format, args...)
}

// Errorf log message with formatting
func Errorf(format string, args ...interface{}) {
	globalLogger().Errorf(format, args...)
}

// Fatalf log message with formatting
func Fatalf(format string, args ...interface{}) {
	globalLogger().Fatalf(format, args...)
}

// Panicf log message with formatting
func Panicf(format string, args ...interface{}) {
	globalLogger().Panicf(format, args...)
}

// WithFields
func WithFields(fields Fields) Logger {
	return globalLogger().WithFields(fields)
}
