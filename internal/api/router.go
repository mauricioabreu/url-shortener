package api

import (
	"github.com/mauricioabreu/url-shortener/internal/api/handlers"
	"github.com/mauricioabreu/url-shortener/internal/api/server"
)

func ExposeRoutes(srv *server.Server) {
	srv.GET("/healthcheck", handlers.DoHealthcheck)
	srv.POST("/shorten", handlers.Shorten)
}
