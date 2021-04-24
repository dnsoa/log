package log

import (
	"fmt"
	stdlog "log"
	"os"
)

type simple struct {
	fields map[string]interface{}
}

// NewSimple creates a basic logger that wraps the core log library.
func NewSimple() Logger {
	return &simple{}
}

// WithFields will return a new logger based on the original logger
// with the additional supplied fields
func (b *simple) WithFields(fields Fields) Logger {
	cp := &simple{}

	if b.fields == nil {
		cp.fields = fields
		return cp
	}

	cp.fields = make(map[string]interface{}, len(b.fields)+len(fields))
	for k, v := range b.fields {
		cp.fields[k] = v
	}

	for k, v := range fields {
		cp.fields[k] = v
	}

	return cp
}

// Debug log message
func (b *simple) Debug(msg ...interface{}) {
	stdlog.Printf("[DEBUG] %s %s\n", fmt.Sprint(msg...), pretty(b.fields))
}

// Info log message
func (b *simple) Info(msg ...interface{}) {
	stdlog.Printf("[INFO] %s %s\n", fmt.Sprint(msg...), pretty(b.fields))
}

// Warn log message
func (b *simple) Warn(msg ...interface{}) {
	stdlog.Printf("[WARN] %s %s\n", fmt.Sprint(msg...), pretty(b.fields))
}

// Error log message
func (b *simple) Error(msg ...interface{}) {
	stdlog.Printf("[ERROR] %s %s\n", fmt.Sprint(msg...), pretty(b.fields))
}

// Fatal log message (and exit)
func (b *simple) Fatal(msg ...interface{}) {
	stdlog.Printf("[FATAL] %s %s\n", fmt.Sprint(msg...), pretty(b.fields))
	os.Exit(1)
}

// Panic log message (and exit)
func (b *simple) Panic(msg ...interface{}) {
	stdlog.Printf("[PANIC] %s %s\n", fmt.Sprint(msg...), pretty(b.fields))
	panic(fmt.Sprint(msg...))
}

// Debugf log message with formatting
func (b *simple) Debugf(format string, args ...interface{}) {
	stdlog.Print(fmt.Sprintf("[DEBUG] "+format+"\n", args...), " ", pretty(b.fields))
}

// Infof log message with formatting
func (b *simple) Infof(format string, args ...interface{}) {
	stdlog.Print(fmt.Sprintf("[INFO] "+format+"\n", args...), " ", pretty(b.fields))
}

// Warnf log message with formatting
func (b *simple) Warnf(format string, args ...interface{}) {
	stdlog.Print(fmt.Sprintf("[WARN] "+format+"\n", args...), " ", pretty(b.fields))
}

// Errorf log message with formatting
func (b *simple) Errorf(format string, args ...interface{}) {
	stdlog.Print(fmt.Sprintf("[ERROR] "+format+"\n", args...), " ", pretty(b.fields))
}

// Fatalf log message with formatting
func (b *simple) Fatalf(format string, args ...interface{}) {
	stdlog.Print(fmt.Sprintf("[FATAL] "+format+"\n", args...), " ", pretty(b.fields))
	os.Exit(1)
}

// Panicf log message with formatting
func (b *simple) Panicf(format string, args ...interface{}) {
	stdlog.Print(fmt.Sprintf("[PANIC] "+format+"\n", args...), " ", pretty(b.fields))
	panic(fmt.Sprintf(format, args...))
}

// helper for pretty printing of fields
func pretty(m map[string]interface{}) string {
	if len(m) < 1 {
		return ""
	}

	s := ""
	for k, v := range m {
		s += fmt.Sprintf("%s=%v ", k, v)
	}

	return s[:len(s)-1]
}
