package log

import (
	"errors"
)

// Logger is a generic interface used in the app.
type Logger interface {
	Debug(args ...interface{})
	Debugf(format string, args ...interface{})
	Info(args ...interface{})
	Infof(format string, args ...interface{})
	Warn(args ...interface{})
	Warnf(format string, args ...interface{})
	Error(args ...interface{})
	Errorf(format string, args ...interface{})
	Fatal(args ...interface{})
	Fatalf(format string, args ...interface{})
	Panic(args ...interface{})
	Panicf(format string, args ...interface{})
}

var logger Logger = &NilLogger{}

// Setup sets up the logger.
func Setup(l Logger) error {
	if l == nil {
		return errors.New("nil logger")
	}

	logger = l
	return nil
}

// Panic logging.
func Panic(args ...interface{}) {
	logger.Panic(args...)
}

// Panicf logging.
func Panicf(format string, args ...interface{}) {
	logger.Panicf(format, args...)
}

// Fatal logging.
func Fatal(args ...interface{}) {
	logger.Fatal(args...)
}

// Fatalf logging.
func Fatalf(format string, args ...interface{}) {
	logger.Fatalf(format, args...)
}

// Error logging.
func Error(args ...interface{}) {
	logger.Error(args...)
}

// Errorf logging.
func Errorf(format string, args ...interface{}) {
	logger.Errorf(format, args...)
}

// Warn logging.
func Warn(args ...interface{}) {
	logger.Warn(args...)
}

// Warnf logging.
func Warnf(format string, args ...interface{}) {
	logger.Warnf(format, args...)
}

// Info logging.
func Info(args ...interface{}) {
	logger.Info(args...)
}

// Infof logging.
func Infof(format string, args ...interface{}) {
	logger.Infof(format, args...)
}

// Debug logging.
func Debug(args ...interface{}) {
	logger.Debug(args...)
}

// Debugf logging.
func Debugf(format string, args ...interface{}) {
	logger.Debugf(format, args...)
}

// NilLogger is a default logger that logs nothing.
type NilLogger struct{}

// Panic logging.
func (nl *NilLogger) Panic(args ...interface{}) {}

// Panicf logging.
func (nl *NilLogger) Panicf(format string, args ...interface{}) {}

// Error logging.
func (nl *NilLogger) Error(args ...interface{}) {}

// Errorf logging.
func (nl *NilLogger) Errorf(format string, args ...interface{}) {}

// Debug logging.
func (nl *NilLogger) Debug(args ...interface{}) {}

// Debugf logging.
func (nl *NilLogger) Debugf(format string, args ...interface{}) {}

// Info logging.
func (nl *NilLogger) Info(args ...interface{}) {}

// Infof logging.
func (nl *NilLogger) Infof(format string, args ...interface{}) {}

// Warn logging.
func (nl *NilLogger) Warn(args ...interface{}) {}

// Warnf logging.
func (nl *NilLogger) Warnf(format string, args ...interface{}) {}

// Fatal logging.
func (nl *NilLogger) Fatal(args ...interface{}) {}

// Fatalf logging.
func (nl *NilLogger) Fatalf(format string, args ...interface{}) {}
