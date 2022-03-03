package main

import (
	"go.uber.org/zap"
)

// NewZapLogger returns a new zap logger.
// TODO: Only new dev, think about passing a configuration and/or a switch
// between prod/dev.
func NewZapLogger() *zap.Logger {
	logger, _ := zap.NewDevelopment()

	return logger
}
