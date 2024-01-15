package logger

import (
	"fmt"
	"log/slog"
	"os"
)

type Logger interface {
	Info(msg string, args ...interface{})
	Infof(format string, args ...interface{})
	Error(msg string, args ...interface{})
	Errorf(format string, args ...interface{})
}

type log struct {
	log *slog.Logger
}

func NewLogger() Logger {
	return &log{
		slog.New(slog.NewTextHandler(os.Stdout, nil)),
	}
}

func (l log) Info(msg string, args ...interface{}) {
	l.log.Info(msg, args...)
}

func (l log) Infof(format string, args ...interface{}) {
	l.log.Info(fmt.Sprintf(format, args...))
}

func (l log) Error(msg string, args ...interface{}) {
	l.log.Error(msg, args...)
}

func (l log) Errorf(format string, args ...interface{}) {
	l.log.Error(fmt.Sprintf(format, args...))
}
