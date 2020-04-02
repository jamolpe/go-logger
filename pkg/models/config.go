package models

import "time"

type Configuration struct {
	DisplayLogs bool
	LogLevels   DisplayConfiguration
	SaveLogs    bool
	Repository  Repository
}

type DisplayConfiguration struct {
	DisplayDebug    bool
	DisplayWarnings bool
	DisplayError    bool
	DisplayInfo     bool
}

type LogModel struct {
	Head    string
	Message string
	Time    time.Time
}

type Repository interface {
	SaveLog(log LogModel) error
}
