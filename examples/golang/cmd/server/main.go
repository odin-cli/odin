package main

import (
	"github.com/lbernardo/fx-webserver/internal/providers"
	"github.com/lbernardo/fx-webserver/internal/providers/server"
	"go.uber.org/fx"
)

func main() {
	newServerRegisters().Run()
}

func newServerRegisters() *fx.App {
	return fx.New(
		fx.Provide(providers.DefaultProviders...),
		server.New(),
	)
}
