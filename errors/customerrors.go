package errors

import (
	"encoding/json"
	"fmt"
	"os"
	"runtime"
	"strings"
	"time"
)

// LogLevel defines the level of logging
type LogLevel int

const (
	// DebugLogLevel represents the debug log level
	DebugLogLevel LogLevel = iota
	// InfoLogLevel represents the info log level
	InfoLogLevel
	// WarnLogLevel represents the warning log level
	WarnLogLevel
	// ErrorLogLevel represents the error log level
	ErrorLogLevel
	// PanicLogLevel represents the panic log level
	PanicLogLevel
	// FatalLogLevel represents the fatal log level
	FatalLogLevel
)

// Logger is a structure that implements logging
type Logger struct {
	logLevel LogLevel
}

// NewLogger creates a new instance of the Logger
func NewLogger(level string) *Logger {
	logLevel := getLogLevel(level)
	return &Logger{logLevel: logLevel}
}

func getLogLevel(level string) LogLevel {
	switch strings.ToLower(level) {
	case "debug":
		return DebugLogLevel
	case "info":
		return InfoLogLevel
	case "warn":
		return WarnLogLevel
	case "error":
		return ErrorLogLevel
	case "panic":
		return PanicLogLevel
	case "fatal":
		return FatalLogLevel
	default:
		return InfoLogLevel
	}
}

// Debug logs a debug message
func (l *Logger) Debug(message string, args ...map[string]interface{}) {
	if l.logLevel <= DebugLogLevel {
		msg := message

		for _, arg := range args {
			jsonStr, _ := json.Marshal(arg)
			msg = fmt.Sprintf("%s %s", msg, string(jsonStr))
		}

		l.log("[DEBUG]", msg)
	}
}

// Info logs an info message
func (l *Logger) Info(message string, args ...interface{}) {
	if l.logLevel <= InfoLogLevel {
		l.log("[INFO]", message, args...)
	}
}

// Warn logs a warning message
func (l *Logger) Warn(message string, args ...interface{}) {
	if l.logLevel <= WarnLogLevel {
		l.log("[WARNING]", message, args...)
	}
}

// Error logs an error message
func (l *Logger) Error(message string, err error, args ...interface{}) {
	if l.logLevel <= ErrorLogLevel {
		l.log("[ERROR]", fmt.Sprintf("%s: %s", message, err), args...)
	}
}

// Panic logs a panic message
func (l *Logger) Panic(message string, args ...interface{}) {
	if l.logLevel <= PanicLogLevel {
		l.log("[PANIC]", message, args...)
		panic(fmt.Sprintf(message, args...))
	}
}

// Fatal logs a fatal message
func (l *Logger) Fatal(message string, args ...interface{}) {
	if l.logLevel <= FatalLogLevel {
		l.log("[FATAL]", message, args...)
		os.Exit(1)
	}
}

func (l *Logger) log(level, message string, args ...interface{}) {
	_, file, line, ok := runtime.Caller(2)
	if ok {
		now := time.Now()
		timeStr := now.Format("2006-01-02 15:04:05")
		file = l.shortenFilePath(file)
		fmt.Printf("%s %s %s:%d %s\n", level, timeStr, file, line, fmt.Sprintf(message, args...))
	}
}

func (l *Logger) shortenFilePath(file string) string {
	parts := strings.Split(file, "/")
	return parts[len(parts)-1]
}
