package logger

import (
	"fmt"
	"time"
)

type LogLevel int

const (
	DEBUG LogLevel = iota
	INFO
	WARN
	ERROR
	None
)

type logger struct {
	LogLevel
}

type Logger interface {
	Debug(f string, v ...any)
	Info(f string, v ...any)
	Warn(f string, v ...any)
	Error(f string, v ...any)
	Panic(f string, v ...any)
}

func NewLogger(l LogLevel) Logger {
	return logger{LogLevel: l}
}

// LogLevelがDebug以下の時ログを標準出力する
func (l logger) Debug(f string, v ...any) {
	if l.LogLevel > DEBUG {
		return
	}
	fmt.Printf("%s [DEBUG] mes=%s", l.TimeFormat(), fmt.Sprintf(f, v...))
}

// LogLevelがInfo以下の時ログを標準出力する
func (l logger) Info(f string, v ...any) {
	if l.LogLevel > INFO {
		return
	}
	fmt.Printf("%s [INFO] mes=%s", l.TimeFormat(), fmt.Sprintf(f, v...))
}

// LogLevelがWarn以下の時デバッグログを標準出力する
func (l logger) Warn(f string, v ...any) {
	if l.LogLevel > WARN {
		return
	}
	fmt.Printf("%s [WARN] mes=%s", l.TimeFormat(), fmt.Sprintf(f, v...))
}

// LogLevelがError以下の時デバッグログを標準出力する
func (l logger) Error(f string, v ...any) {
	if l.LogLevel > ERROR {
		return
	}
	fmt.Printf("%s [ERROR] mes=%s", l.TimeFormat(), fmt.Sprintf(f, v...))
}

// エラーログを出力し,Panicにする
func (l logger) Panic(f string, v ...any) {
	l.Error(f, v...)
	panic(fmt.Sprintf(f, v...))
}

func (l logger) TimeFormat() string {
	return time.Now().Format("2006-01-02 15:04:05.000000 +0900 UTC")
}
