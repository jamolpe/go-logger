package models

type Configuration struct {
	DisplayLogs bool
	LogLevel    string
	SaveLogs    string
	Repository  Repository
}

type LogModel struct {
	TypeCode   int
	Message    string
	ColourCode string
}

type Repository interface {
	SaveLog() error
}
