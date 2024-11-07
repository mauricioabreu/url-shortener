package server

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/requestid"
	ginzap "github.com/gin-contrib/zap"
	"github.com/mauricioabreu/url-shortener/internal/config"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type Server struct {
	*gin.Engine
	server *http.Server
}

// New create a new server using Gin
// Configure a few middlewares
func New(cfg *config.Config, logger *zap.Logger) *Server {
	gin.SetMode(gin.ReleaseMode)

	r := gin.New()

	r.Use(cors.New(configureCORS()))
	r.Use(gin.Recovery())
	r.Use(requestid.New())
	r.Use(timeoutMiddleware(cfg.Server.Timeout))
	r.Use(ginzap.Ginzap(logger, time.RFC3339, true))
	r.Use(ginzap.RecoveryWithZap(logger, true))

	return &Server{
		Engine: r,
		server: &http.Server{
			Addr:         ":" + cfg.Server.Port,
			Handler:      r,
			ReadTimeout:  cfg.Server.Timeout,
			WriteTimeout: cfg.Server.Timeout,
		},
	}
}

func RegisterHooks(lc fx.Lifecycle, srv *Server, cfg *config.Config) {
	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			go func() {
				if err := srv.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
					panic(err)
				}
			}()

			return nil
		},
		OnStop: func(ctx context.Context) error {
			return srv.server.Shutdown(ctx)
		},
	})
}
