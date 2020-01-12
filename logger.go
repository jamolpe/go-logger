package logger

import (
	"fmt"
	"time"
)

type mType string

const (
	error   = mType("ERROR")
	info    = mType("INFO")
	warning = mType("WARNING")
	debug   = mType("DEBUG")
)

type (
	loggerMessage struct {
		time    time.Time
		message string
	}
)

func createBaseForLog(message string) *loggerMessage {
	date := time.Now()
	return &loggerMessage{time: date, message: message}
}

func printMessage(typeM mType, loggerMessage *loggerMessage) {
	fmt.Println(typeM, " at: ", loggerMessage.time.Format(time.RFC3339), " message: ", loggerMessage.message)
}

// INFO : print information message
func INFO(message string) {
	loggerMessage := createBaseForLog(message)
	printMessage(info, loggerMessage)
}

// ERROR : print error information message
func ERROR(message string) {
	loggerMessage := createBaseForLog(message)
	printMessage(error, loggerMessage)
}

// WARNING : print warning message
func WARNING(message string) {
	loggerMessage := createBaseForLog(message)
	printMessage(warning, loggerMessage)
}

// DEBUG : print message at debug lvl
func DEBUG(message string) {
	loggerMessage := createBaseForLog(message)
	printMessage(warning, loggerMessage)
}
