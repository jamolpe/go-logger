package main

import (
	logger "gologger"
	"gologger/pkg/models"
)

func main() {
	config := models.Configuration{DisplayLogs: true, SaveLogs: false,
		LogLevels: models.DisplayConfiguration{
			DisplayDebug:    true,
			DisplayWarnings: true,
			DisplayError:    true,
			DisplayInfo:     true,
		},
	}
	logger := logger.New(config)
	logger.DEBUG("I'm debug")
	logger.ERROR("I'm error")
	logger.INFO("I'm info")
	logger.WARNING("I'm waring")
}
