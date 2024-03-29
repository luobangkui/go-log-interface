package log

import (
	"fmt"
	"github.com/go-log/log"
	golog "github.com/go-log/log/log"
	"os"
)

// level is a log level
type Level int

const (
	LevelFatal Level = iota
	LevelError
	LevelInfo
	LevelWarn
	LevelDebug
)

var (
	// the local logger
	logger log.Logger = golog.New()

	// default log level is info
	level = LevelInfo
)

func init() {
	switch os.Getenv("CTRPRE_LOG_LEVEL") {
	case "debug":
		level = LevelDebug
	case "info":
		level = LevelInfo
	case "error":
		level = LevelError
	case "fatal":
		level = LevelFatal
	case "warn":
		level = LevelWarn
	}
}

func (l *Level) String() string {
	switch *l {
	case LevelFatal:
		return "FATAL"
	case LevelInfo:
		return "INFO"
	case LevelWarn:
		return "WARN"
	case LevelDebug:
		return "DEBUG"
	case LevelError:
		return "ERROR"
	}
	return ""
}

// Log makes use of github.com/go-log/log.Log
func Log(v ...interface{}) {
	logger.Log(v...)
}

// Logf makes use of github.com/go-log/log.Logf
func Logf(format string, v ...interface{}) {
	logger.Logf(format, v...)
}

// WithLevel logs with the level specified
func WithLevel(l Level, v ...interface{}) {
	if l > level {
		return
	}
	Log(fmt.Sprintf(l.String()+": %v", v...))
}

// WithLevel logs with the level specified
func WithLevelf(l Level, format string, v ...interface{}) {
	if l > level {
		return
	}
	Logf(l.String()+": "+format, v...)
}

// Trace provides trace level logging
func Error(v ...interface{}) {
	WithLevel(LevelError, v...)
}

// Tracef provides trace level logging
func Errorf(format string, v ...interface{}) {
	WithLevelf(LevelError, format, v...)
}

// Debug provides debug level logging
func Debug(v ...interface{}) {
	WithLevel(LevelDebug, v...)
}

// Debugf provides debug level logging
func Debugf(format string, v ...interface{}) {
	WithLevelf(LevelDebug, format, v...)
}

// Info provides info level logging
func Info(v ...interface{}) {
	WithLevel(LevelInfo, v...)
}

// Infof provides info level logging
func Infof(format string, v ...interface{}) {
	WithLevelf(LevelInfo, format, v...)
}

// Warn provides warn level logging
func Warn(v ...interface{}) {
	WithLevel(LevelWarn, v...)
}

// Warnf provides warn level logging
func Warnf(format string, v ...interface{}) {
	WithLevelf(LevelWarn, format, v...)
}

// Fatal logs with Log and then exits with os.Exit(1)
func Fatal(v ...interface{}) {
	WithLevel(LevelFatal, v...)
	os.Exit(1)
}

// Fatalf logs with Logf and then exits with os.Exit(1)
func Fatalf(format string, v ...interface{}) {
	WithLevelf(LevelFatal, format, v...)
	os.Exit(1)
}

// SetLogger sets the local logger
func SetLogger(l log.Logger) {
	logger = l
}

// GetLogger returns the local logger
func GetLogger() log.Logger {
	return logger
}

// SetLevel sets the log level
func SetLevel(l Level) {
	level = l
}

// GetLevel returns the current level
func GetLevel() Level {
	return level
}
