package core

import "gologger/pkg/models"

type (
	Core struct {
		repo *models.Repository
		configuration
	}
	configuration struct {
		displayLogs bool
		logLevel    string
		saveLogs    string
	}
)

func New(config models.Configuration) *Core {
	core := &Core{repo: &config.Repository}
	configure(core, config)
	return core
}

func configure(c *Core, config models.Configuration) {
	c.configuration.displayLogs = config.DisplayLogs
	c.configuration.logLevel = config.LogLevel
	c.configuration.saveLogs = config.SaveLogs
}

func (c *Core) ReportError(msg string) {
	// print it
	// save it
}
func (c *Core) ReportInfo(msg string) {

}

func (c *Core) ReportDebug(msg string) {

}

func (c *Core) printMessage() {
	//test
}
