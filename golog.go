package golog

import (
	"fmt"
	"io"
	"os"
)

//--------------------
// LOG LEVEL
//--------------------

// Log levels to control the logging output.
const (
	LevelTrace = iota
	LevelDebug
	LevelInfo
	LevelWarning
	LevelError
	LevelCritical
)

// logLevel controls the global log level used by the logger.
var level = LevelTrace

// LogLevel returns the global log level and can be used in
// own implementations of the logger interface.
func Level() int {
	return level
}

// SetLogLevel sets the global log level used by the simple
// logger.
func SetLevel(l int) {
	level = l
}

// SetFlags sets the output flags for the logger.
func SetFlags(flag int) {
	GoLogger.SetFlags(flag)
}

// logger references the used application logger.
var GoLogger = New(os.Stdout, "", Ldate|Ltime|Lshortfile, 1)

func SetOutput(w io.Writer) {
	GoLogger = New(w, "", Ldate|Ltime|Lshortfile, 1)
}

// SetLogger sets a new logger.
func SetLogger(l *Logger) {
	GoLogger = l
}

// Trace logs a message at trace level.
func Trace(v ...interface{}) {
	if level <= LevelTrace {
		GoLogger.Printf("[T] %s", fmt.Sprintln(v...))
	}
}

// Debug logs a message at debug level.
func Debug(v ...interface{}) {
	if level <= LevelDebug {
		GoLogger.Printf("[D] %s", fmt.Sprintln(v...))
	}
}

// Info logs a message at info level.
func Info(v ...interface{}) {
	if level <= LevelInfo {
		GoLogger.Printf("[I] %s", fmt.Sprintln(v...))
	}
}

// Warning logs a message at warning level.
func Warn(v ...interface{}) {
	if level <= LevelWarning {
		GoLogger.Printf("[W] %s", fmt.Sprintln(v...))
	}
}

// Error logs a message at error level.
func Error(v ...interface{}) {
	if level <= LevelError {
		GoLogger.Printf("[E] %s", fmt.Sprintln(v...))
	}
}

// Critical logs a message at critical level.
func Critical(v ...interface{}) {
	if level <= LevelCritical {
		GoLogger.Printf("[C] %s", fmt.Sprintln(v...))
		os.Exit(1)
	}
}
