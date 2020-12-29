package logger

import (
	"github.com/thetogi/YReserve2/model"
)

type TestLogger struct {
}

func NewTestLogger(config *model.Config) *TestLogger {
	testLogger := &TestLogger{}
	return testLogger
}

func (zl *TestLogger) Debug(message string, args ...Argument) {
}

func (zl *TestLogger) Info(message string, args ...Argument) {
}

func (zl *TestLogger) Warn(message string, args ...Argument) {
}

func (zl *TestLogger) Error(message string, args ...Argument) {
}

func (zl *TestLogger) OnConfigChange(newConfig *model.Config) {
}
