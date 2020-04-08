package models

import "time"

// Configuration config information that needs to be provided
type Configuration struct {
	DisplayLogs bool
	LogLevels   DisplayConfiguration
	SaveLogs    bool
	Repository  Repository
}

// DisplayConfiguration DisplayDebug DisplayWarnings DisplayError and DisplayInfo available options
type DisplayConfiguration struct {
	DisplayDebug    bool
	DisplayWarnings bool
	DisplayError    bool
	DisplayInfo     bool
}

// LogModel working log model for the Database formated head is ignored in database -bson-
type LogModel struct {
	FormatedHead string    `bson:"-"`
	Head         string    `bson:"Head"`
	Message      string    `bson:"Message"`
	Time         time.Time `bson:"Time"`
}

// Repository interface to implement in case you need to save the logs(recomended)
type Repository interface {
	SaveLog(log LogModel) error
}
