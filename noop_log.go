package log

type noop struct{}

// NewNoop creates a no-op logger that can be used to silence
// all logging from this library. Also useful in tests.
func NewNoop() Logger {
	return &noop{}
}

// Debug log message no-op
func (n *noop) Debug(msg ...interface{}) {}

// Info log message no-op
func (n *noop) Info(msg ...interface{}) {}

// Warn log message no-op
func (n *noop) Warn(msg ...interface{}) {}

// Error log message no-op
func (n *noop) Error(msg ...interface{}) {}

// Fatal log message no-op
func (n *noop) Fatal(msg ...interface{}) {}

// Panic log message no-op
func (n *noop) Panic(msg ...interface{}) {}

// Debugf log message with formatting no-op
func (n *noop) Debugf(format string, args ...interface{}) {}

// Infof log message with formatting no-op
func (n *noop) Infof(format string, args ...interface{}) {}

// Warnf log message with formatting no-op
func (n *noop) Warnf(format string, args ...interface{}) {}

// Errorf log message with formatting no-op
func (n *noop) Errorf(format string, args ...interface{}) {}

// Fatalf log message with formatting no-op
func (n *noop) Fatalf(format string, args ...interface{}) {}

// Panicf log message with formatting no-op
func (n *noop) Panicf(format string, args ...interface{}) {}

// WithFields no-op
func (n *noop) WithFields(fields Fields) Logger { return n }
