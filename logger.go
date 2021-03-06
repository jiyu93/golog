package golog

import (
	"fmt"
	"io"
	"log"
	"os"
)

// Log levels
const (
	LevelTrace = iota
	LevelDebug
	LevelInfo
	LevelWarn
	LevelError
	LevelPanic
	LevelFatal
)

// Logger ...
type Logger struct {
	lg        *log.Logger
	lv        int
	calldepth int
}

// NewLogger 创建Logger，并指定log level
func NewLogger(w io.Writer, level int) *Logger {
	return &Logger{
		log.New(w, "", log.Ldate|log.Ltime|log.Lshortfile|log.Lmicroseconds),
		level,
		3,
	}
}

// SetOutput 修改输出目的地
func (l *Logger) SetOutput(w io.Writer) {
	l.lg.SetOutput(w)
}

// SetLevel 修改日志输出级别
func (l *Logger) SetLevel(level int) {
	l.lv = level
}

// SetCallDepth 修改默认CallDepth
func (l *Logger) SetCallDepth(calldepth int) {
	l.calldepth = calldepth
}

// SetFlags 使用与标准库"log"一致的flag对日志输出配置进行修改
func (l *Logger) SetFlags(flag int) {
	l.lg.SetFlags(flag)
}

// Trace 输出Trace级别的日志
func (l *Logger) Trace(v ...interface{}) {
	if l.lv <= LevelTrace {
		l.lg.Output(l.calldepth, fmt.Sprintf("[T] %s", fmt.Sprintln(v...)))
	}
}

// Debug 输出Debug级别的日志
func (l *Logger) Debug(v ...interface{}) {
	if l.lv <= LevelDebug {
		l.lg.Output(l.calldepth, fmt.Sprintf("[D] %s", fmt.Sprintln(v...)))
	}
}

// Info 输出Info级别的日志
func (l *Logger) Info(v ...interface{}) {
	if l.lv <= LevelInfo {
		l.lg.Output(l.calldepth, fmt.Sprintf("[I] %s", fmt.Sprintln(v...)))
	}
}

// Warn 输出Warn级别的日志
func (l *Logger) Warn(v ...interface{}) {
	if l.lv <= LevelWarn {
		l.lg.Output(l.calldepth, fmt.Sprintf("[W] %s", fmt.Sprintln(v...)))
	}
}

// Error 输出Error级别的日志
func (l *Logger) Error(v ...interface{}) {
	if l.lv <= LevelError {
		l.lg.Output(l.calldepth, fmt.Sprintf("[E] %s", fmt.Sprintln(v...)))
	}
}

// Panic 输出Panic级别的日志，然后调用panic()函数
func (l *Logger) Panic(v ...interface{}) {
	if l.lv <= LevelPanic {
		s := fmt.Sprintf("[P] %s", fmt.Sprintln(v...))
		l.lg.Output(l.calldepth, s)
		panic(s)
	}
}

// Fatal 输出Fatal级别的日志，然后直接退出程序
func (l *Logger) Fatal(v ...interface{}) {
	if l.lv <= LevelFatal {
		l.lg.Output(l.calldepth, fmt.Sprintf("[F] %s", fmt.Sprintln(v...)))
		os.Exit(1)
	}
}
