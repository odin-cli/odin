package server

import (
	"github.com/lbernardo/fx-webserver/internal/providers/server/handlers"
	"go.uber.org/fx"
)

var _handlers = []any{
	handlers.NewHandler,
}

func New() fx.Option {
	return fx.Module("server",
		fx.Provide(NewGinEngine),
		fx.Provide(NewRouter),
		fx.Provide(_handlers...),
		fx.Invoke(RegisterRoutes),
		fx.Invoke(Server),
	)
}
