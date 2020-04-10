package core

import (
	"fmt"
	"runtime"
	"time"

	"gologger/internal/printer"
	"gologger/pkg/models"
)

type (
	Core struct {
		repo models.Repository
		configuration
	}
	configuration struct {
		displayLogs bool
		logLevels   models.DisplayConfiguration
		saveLogs    bool
	}
)

const (
	infoColour    = "\033[1;36m%s\033[0m"
	errorColour   = "\033[1;31m%s\033[0m"
	warningColour = "\033[1;33m%s\033[0m"
	debugColour   = "\033[1;35m%s\033[0m"
)

// New creates a new configuration
func New(config models.Configuration) *Core {
	core := &Core{repo: config.Repository}
	configure(core, config)
	return core
}

func configure(c *Core, config models.Configuration) {
	c.configuration.displayLogs = config.DisplayLogs
	c.configuration.logLevels = config.LogLevels
	c.configuration.saveLogs = config.SaveLogs
}

func (c *Core) ReportWarning(msg string) {
	if c.configuration.logLevels.DisplayWarnings {
		formatedHead := ""
		if runtime.GOOS == "windows" {
			formatedHead = "WARNING"
		} else {
			formatedHead = fmt.Sprintf(warningColour, "WARNING")
		}
		c.processLog("WARNING", formatedHead, msg)
	}
}

func (c *Core) ReportError(msg string) {
	if c.configuration.logLevels.DisplayError {
		formatedHead := ""
		if runtime.GOOS == "windows" {
			formatedHead = "ERROR"
		} else {
			formatedHead = fmt.Sprintf(errorColour, "ERROR")
		}
		c.processLog("ERROR", formatedHead, msg)
	}
}

func (c *Core) ReportInfo(msg string) {
	if c.configuration.logLevels.DisplayInfo {
		formatedHead := ""
		if runtime.GOOS == "windows" {
			formatedHead = "INFO"
		} else {
			formatedHead = fmt.Sprintf(infoColour, "INFO")
		}
		c.processLog("INFO", formatedHead, msg)
	}
}

func (c *Core) ReportDebug(msg string) {
	if c.configuration.logLevels.DisplayDebug {
		formatedHead := ""
		if runtime.GOOS == "windows" {
			formatedHead = "DEBUG"
		} else {
			formatedHead = fmt.Sprintf(debugColour, "DEBUG")
		}
		c.processLog("DEBUG", formatedHead, msg)
	}
}

func (c *Core) processLog(head, formatedhead, msg string) {
	log := models.LogModel{FormatedHead: formatedhead, Head: head, Message: msg, Time: time.Now()}
	if c.configuration.saveLogs && c.repo != nil {
		c.repo.SaveLog(log)
	}
	if c.configuration.displayLogs {
		printer.PrintMessage(log)
	}
}
