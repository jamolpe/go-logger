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
	logger.DEBUG("soy debug")
	logger.ERROR("soy error")
	logger.INFO("soy info")
	logger.WARNING("soy waring")
}
