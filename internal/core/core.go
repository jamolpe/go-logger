package core

import (
	"fmt"
	"gologger/internal/printer"
	"gologger/pkg/models"
	"time"
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
		head := fmt.Sprintf(warningColour, "WARNING")
		c.workWithLog(head, msg)
	}
}

func (c *Core) ReportError(msg string) {
	if c.configuration.logLevels.DisplayError {
		head := fmt.Sprintf(errorColour, "ERROR")
		c.workWithLog(head, msg)
	}
}

func (c *Core) ReportInfo(msg string) {
	if c.configuration.logLevels.DisplayInfo {
		head := fmt.Sprintf(infoColour, "INFO")
		c.workWithLog(head, msg)
	}
}

func (c *Core) ReportDebug(msg string) {
	if c.configuration.logLevels.DisplayDebug {
		head := fmt.Sprintf(debugColour, "DEBUG")
		c.workWithLog(head, msg)
	}
}

func (c *Core) workWithLog(head, msg string) {
	log := models.LogModel{Head: head, Message: msg, Time: time.Now()}
	if c.configuration.saveLogs && c.repo != nil {
		c.repo.SaveLog(log)
	}
	if c.configuration.displayLogs {
		printer.PrintMessage(log)
	}
}