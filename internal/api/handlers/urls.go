package handlers

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mauricioabreu/url-shortener/internal/api/serializers"
	"go.uber.org/zap"
)

type ShortenRequest struct {
	URL string `json:"url" binding:"required,url"`
}

type ShortenerService interface {
	Shorten(ctx context.Context, url string) (string, error)
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

	shortenedURL, err := s.shortenerService.Shorten(c.Request.Context(), req.URL)
	if err != nil {
		s.logger.Error("Failed to shorten URL", zap.Error(err))
		serializers.RespondWithError(c, http.StatusInternalServerError, "Internal server error")
		return
	}

	responseData := gin.H{
		"shortened_url": shortenedURL,
		"long_url":      req.URL,
		"created_at":    time.Now(),
	}
	serializers.RespondWithSuccess(c, http.StatusOK, "Successfully inserted URL", responseData)
}
