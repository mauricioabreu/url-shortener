package main

import (
	"github.com/mauricioabreu/url-shortener/internal/api"
	"github.com/mauricioabreu/url-shortener/internal/api/handlers"
	"github.com/mauricioabreu/url-shortener/internal/api/server"
	"github.com/mauricioabreu/url-shortener/internal/config"
	"github.com/mauricioabreu/url-shortener/internal/infra/db"
	"github.com/mauricioabreu/url-shortener/internal/infra/logging"
	"github.com/mauricioabreu/url-shortener/internal/services/url"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		fx.Provide(
			config.New,
			logging.New,
			db.New,
			server.New,
			fx.Annotate(
				url.NewURLService,
				fx.As(new(handlers.ShortenerService)),
			),
			handlers.NewShortenerHandler,
		),
		fx.Invoke(server.RegisterHooks, api.ExposeRoutes),
	).Run()
}
