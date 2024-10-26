package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mauricioabreu/url-shortener/internal/api/serializers"
)

type ShortenRequest struct {
	URL string `json:"url" binding:"required,url"`
}

func Shorten(c *gin.Context) {
	var req ShortenRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		serializers.RespondWithError(c, http.StatusBadRequest, err.Error())
		return
	}

	serializers.RespondWithSuccess(c, http.StatusOK, "OK", nil)
}
