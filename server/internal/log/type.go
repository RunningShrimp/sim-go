package server

import "go.uber.org/zap"

type logLevel int

const (
	Fatal logLevel = iota
	Panic
	Error
	Warning
	Info
	Debug
)

const (
	errKey = "error"
)

type Logger interface {
}

type InnerStdLogger struct {
}

type StdLogger struct {
	*zap.Logger
	message  string   // 打印的消息
	printVal []string // 即将要打印的值
	level    logLevel
	err      error
}
