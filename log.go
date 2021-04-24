package log

// Logger interface allows you to maintain a unified interface while using a
// custom logger. This allows you to write log statements without dictating
// the specific underlying library used for logging. You can avoid vendoring
// of logging libraries, which is especially useful when writing shared code
// such as a library. This package contains a simple logger and a no-op logger
// which both implement this interface. It is also supplemented with some
// additional helpers/shims for other common logging libraries such as logrus
type Logger interface {
	Debug(msg ...interface{})
	Info(msg ...interface{})
	Warn(msg ...interface{})
	Error(msg ...interface{})
	Fatal(msg ...interface{})
	Panic(msg ...interface{})

	Debugf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Warnf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	Fatalf(format string, args ...interface{})
	Panicf(format string, args ...interface{})

	WithFields(Fields) Logger
}

// Fields is used to define structured fields which are appended to log messages
type Fields map[string]interface{}
