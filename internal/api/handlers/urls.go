package handlers

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mauricioabreu/url-shortener/internal/api/serializers"
	"github.com/mauricioabreu/url-shortener/internal/services/url"
	"go.uber.org/zap"
)

type ShortenRequest struct {
	URL  string   `json:"url" binding:"required,url"`
	Tags []string `json:"tags" binding:"dive,lowercase,alphanum"`
}

type ShortenerService interface {
	Shorten(ctx context.Context, data *url.ShortenData) (*url.ShortenResponse, error)
}

type ShortenerHandler struct {
	shortenerService ShortenerService
	logger           *zap.Logger
}

func NewShortenerHandler(shortenerService ShortenerService, logger *zap.Logger) *ShortenerHandler {
	return &ShortenerHandler{
		shortenerService: shortenerService,
		logger:           logger,
	}
}

func (s *ShortenerHandler) Shorten(c *gin.Context) {
	var req ShortenRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		serializers.RespondWithError(c, http.StatusBadRequest, err.Error())
		return
	}

	shortenData := &url.ShortenData{
		LongURL: req.URL,
		Tags:    req.Tags,
	}
	shortenedURL, err := s.shortenerService.Shorten(c.Request.Context(), shortenData)
	if err != nil {
		s.logger.Error("Failed to shorten URL", zap.Error(err))
		serializers.RespondWithError(c, http.StatusInternalServerError, "Internal server error")
		return
	}

	serializers.RespondWithSuccess(c, http.StatusOK, "Successfully inserted URL", shortenedURL)
}
