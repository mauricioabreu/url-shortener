package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mauricioabreu/url-shortener/internal/api/serializers"
)

type ShortenRequest struct {
	URL string `json:"url" binding:"required,url"`
}

type ShortenerService interface {
	Shorten(url string) (string, error)
}

type ShortenerHandler struct {
	shortenerService ShortenerService
}

func NewShortenerHandler(shortenerService ShortenerService) *ShortenerHandler {
	return &ShortenerHandler{
		shortenerService: shortenerService,
	}
}

func (s *ShortenerHandler) Shorten(c *gin.Context) {
	var req ShortenRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		serializers.RespondWithError(c, http.StatusBadRequest, err.Error())
		return
	}

	shortenedURL, err := s.shortenerService.Shorten(req.URL)
	if err != nil {
		serializers.RespondWithError(c, http.StatusInternalServerError, err.Error())
		return
	}

	serializers.RespondWithSuccess(c, http.StatusOK, "Successfully inserted URL", gin.H{"shortened_url": shortenedURL})
}
