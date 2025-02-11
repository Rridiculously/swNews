package logger

import (
	"backend/pkg/conf"
	"testing"
)

func TestInit(t *testing.T) {
	conf.Init("../../config.yaml")
	Init()
	globalLogger.Info("info")
	globalLogger.Debug("debug")
	globalLogger.Warn("warn")
	globalLogger.Error("error")
}
