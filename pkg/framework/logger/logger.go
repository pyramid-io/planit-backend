package logger

import (
	"errors"
	"fmt"
	"os"
	"sync"
	"time"
)

type LoggerInterface interface {
	Info(message string)
	Notice(message string)
	Warning(message string)
	Error(message string)
	Critical(message string)

	Close()
}

type LogLevel string

const (
	DEBUG    LogLevel = "debug"
	INFO     LogLevel = "info"
	NOTICE   LogLevel = "notice"
	WARNING  LogLevel = "warning"
	ERROR    LogLevel = "error"
	CRITICAL LogLevel = "critical"
)

type Logger struct {
	mu      sync.Mutex
	logFile *os.File
}

func New(logPath string) (*Logger, error) {
	if logPath == "" {
		return nil, errors.New("provided logPath must not be empty")
	}

	logPathFile := fmt.Sprintf("%s/logs.log", logPath)

	logFile, err := os.OpenFile(logPathFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return nil, err
	}

	return &Logger{
		logFile: logFile,
	}, nil
}

func (l *Logger) Close() {
	l.logFile.Close()
}

func (l *Logger) log(level LogLevel, message string) {
	l.mu.Lock()
	defer l.mu.Unlock()

	timestamp := time.Now().Format(time.RFC3339)

	message = fmt.Sprintf("[%s] [%s] %s\n", timestamp, level, message)

	l.logFile.WriteString(message)
}

func (l *Logger) Debug(message string) {
	l.log(DEBUG, message)
}

func (l *Logger) Info(message string) {
	l.log(INFO, message)
}

func (l *Logger) Notice(message string) {
	l.log(NOTICE, message)
}

func (l *Logger) Warning(message string) {
	l.log(WARNING, message)
}

func (l *Logger) Error(message string) {
	l.log(ERROR, message)
}

func (l *Logger) Critical(message string) {
	l.log(CRITICAL, message)
}
