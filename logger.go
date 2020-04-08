package logger

import (
	"gologger/internal/core"
	"gologger/pkg/models"
)

type LoggerI interface {
	INFO(message string)
	ERROR(message string)
	DEBUG(message string)
	WARNING(message string)
}

type Logger struct {
	core *core.Core
}

func New(config models.Configuration) LoggerI {
	core := core.New(config)
	logger := Logger{core: core}
	return &logger
}

func (l *Logger) INFO(message string) {
	l.core.ReportInfo(message)
}

func (l *Logger) ERROR(message string) {
	l.core.ReportError(message)
}

func (l *Logger) DEBUG(message string) {
	l.core.ReportDebug(message)
}

func (l *Logger) WARNING(message string) {
	l.core.ReportWarning(message)
}
