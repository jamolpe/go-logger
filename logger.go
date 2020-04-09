package logger

import (
	"gologger/internal/core"
	"gologger/pkg/models"
)

/*
	Logger interface base
*/
type LoggerI interface {
	INFO(message string)
	ERROR(message string)
	DEBUG(message string)
	WARNING(message string)
}

type Logger struct {
	core *core.Core
}

// New creates a new logger with the configuration
func New(config models.Configuration) LoggerI {
	core := core.New(config)
	logger := Logger{core: core}
	return &logger
}

// INFO logs an info message
func (l *Logger) INFO(message string) {
	l.core.ReportInfo(message)
}

// ERROR logs an error message
func (l *Logger) ERROR(message string) {
	l.core.ReportError(message)
}

// DEBUG logs a debug message
func (l *Logger) DEBUG(message string) {
	l.core.ReportDebug(message)
}

// WARNING logs a warning message
func (l *Logger) WARNING(message string) {
	l.core.ReportWarning(message)
}
