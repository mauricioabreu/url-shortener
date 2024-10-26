package main

import (
	"github.com/mauricioabreu/url-shortener/internal/api"
	"github.com/mauricioabreu/url-shortener/internal/api/server"
	"github.com/mauricioabreu/url-shortener/internal/config"
	"github.com/mauricioabreu/url-shortener/internal/infra/db"
	"github.com/mauricioabreu/url-shortener/internal/infra/logging"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		fx.Provide(
			config.New,
			logging.New,
			db.New,
			server.New,
		),
		fx.Invoke(server.RegisterHooks, api.ExposeRoutes),
	).Run()
}
