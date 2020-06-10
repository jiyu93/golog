package golog

import (
	"io"
	"os"
)

// 全局Logger，默认向标准输出流输出日志
var gLogger = NewLogger(os.Stdout, 0)

// SetDefaultOutput 修改默认的输出目的地
func SetDefaultOutput(w io.Writer) {
	gLogger.SetOutput(w)
}

// SetDefaultFlags 修改默认输出格式
func SetDefaultFlags(flags int) {
	gLogger.SetFlags(flags)
}

// SetDefaultLevel 修改默认输出级别
func SetDefaultLevel(level int) {
	gLogger.SetLevel(level)
}

// Trace 打印Trace级别的日志
func Trace(v ...interface{}) {
	gLogger.Trace(v...)
}

// Debug 打印Debug级别的日志
func Debug(v ...interface{}) {
	gLogger.Debug(v...)
}

// Info 打印Info级别的日志
func Info(v ...interface{}) {
	gLogger.Info(v...)
}

// Warn 打印Warn级别的日志
func Warn(v ...interface{}) {
	gLogger.Warn(v...)
}

// Error 打印Error级别的日志
func Error(v ...interface{}) {
	gLogger.Error(v...)
}

// Panic 打印Panic级别的日志，然后调用panic()函数
func Panic(v ...interface{}) {
	gLogger.Panic(v...)
}

// Fatal 打印Fatal级别的日志，然后直接退出程序
func Fatal(v ...interface{}) {
	gLogger.Fatal(v...)
}
