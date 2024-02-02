package scenexec

import (
	"fmt"

	logger "github.com/subrahamanyam341/andes-logger-123"
)

// SetLoggingForTests configures the logger package with *:TRACE and enabled logger names
func SetLoggingForTests() {
	SetLoggingForTestsWithLogger("*")
}

// SetLoggingForTestsWithLogger configures the logger package with a certain logger
func SetLoggingForTestsWithLogger(loggerName string) {
	_ = logger.SetLogLevel(fmt.Sprintf("*:NONE,%s:TRACE", loggerName))
	logger.ToggleCorrelation(false)
	logger.ToggleLoggerName(true)
}

// DisableLoggingForTests sets log level to *:NONE
func DisableLoggingForTests() {
	_ = logger.SetLogLevel("*:NONE")
}
