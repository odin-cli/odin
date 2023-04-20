package providers

import (
	"github.com/lbernardo/fx-webserver/internal/providers/config"
	"github.com/lbernardo/fx-webserver/internal/providers/logger"
)

var DefaultProviders = []any{
	config.NewConfig,
	logger.NewLogger,
}
