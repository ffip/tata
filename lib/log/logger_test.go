package Logger_test

import (
	"testing"

	xlog "github.com/ffip/tata/lib/log"
)

func TestPrintf(t *testing.T) {
	logger := xlog.Logger{
		Level:    xlog.LevelAll,
		Module:   "LoggerTest",
		Function: "TestPrintf",
		Process:  "Testing"}

	logger.Info("Test:Info!")
	logger.Debug("Test:Debug!")
	logger.WithField("Service").WithField("Mysqld").Warn("Test:Warn!")
	logger.WithExit().WithField("Exit").Error("Test:Tip And Exit!")
}
