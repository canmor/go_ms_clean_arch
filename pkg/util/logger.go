package util

import (
	"fmt"
	"log"
)

type LogLevel int

const (
	LogError LogLevel = iota + 1
	LogInfo
	LogWarn
	LogDebug
	LogFatal
)

type Logger interface {
	Log(level LogLevel, content string)
}

type LoggerWrapper struct {
	logger Logger
}

func (l LoggerWrapper) Debug(args ...interface{}) {
	l.logger.Log(LogDebug, fmt.Sprint(args...))
}

func (l LoggerWrapper) Debugf(format string, args ...interface{}) {
	l.logger.Log(LogDebug, fmt.Sprintf(format, args...))
}

func (l LoggerWrapper) Info(args ...interface{}) {
	l.logger.Log(LogInfo, fmt.Sprint(args...))
}

func (l LoggerWrapper) Infof(format string, args ...interface{}) {
	l.logger.Log(LogInfo, fmt.Sprintf(format, args...))
}

func (l LoggerWrapper) Warn(args ...interface{}) {
	l.logger.Log(LogWarn, fmt.Sprint(args...))
}

func (l LoggerWrapper) Warnf(format string, args ...interface{}) {
	l.logger.Log(LogWarn, fmt.Sprintf(format, args...))
}

func (l LoggerWrapper) Error(args ...interface{}) {
	l.logger.Log(LogError, fmt.Sprint(args...))
}

func (l LoggerWrapper) Errorf(format string, args ...interface{}) {
	l.logger.Log(LogError, fmt.Sprintf(format, args...))
}

func (l LoggerWrapper) Fatal(args ...interface{}) {
	l.logger.Log(LogFatal, fmt.Sprint(args...))
}

func (l LoggerWrapper) Fatalf(format string, args ...interface{}) {
	l.logger.Log(LogFatal, fmt.Sprintf(format, args...))
}

type LoggerImpl struct {
}

func (l LoggerImpl) Log(_ LogLevel, content string) {
	log.Printf(content)
}
