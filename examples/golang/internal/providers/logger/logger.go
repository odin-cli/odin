package logger

import (
	"github.com/lbernardo/fx-webserver/pkg/model/config"
	"go.uber.org/zap"
)

func NewLogger(config *config.Config) *zap.Logger {
	logger, _ := zap.NewProduction()
	if config.Log.Mode == "development" {
		logger, _ = zap.NewDevelopment()
		logger.Info("Log start in development mode! Change in config.yaml -> log.mode='production'")
	}
	return logger
}
