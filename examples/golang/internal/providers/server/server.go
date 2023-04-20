package server

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

func Server(lc fx.Lifecycle, engine *gin.Engine) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go engine.Run(":8000")
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return nil
		},
	})
}
