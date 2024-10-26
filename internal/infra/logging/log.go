package logging

import (
	"github.com/mauricioabreu/url-shortener/internal/config"
	"go.uber.org/zap"
)

func New(cfg *config.Config) (*zap.Logger, error) {
	logConfig := zap.NewProductionConfig()
	logConfig.Level = zap.NewAtomicLevelAt(cfg.Logging.Level)

	return logConfig.Build()
}
