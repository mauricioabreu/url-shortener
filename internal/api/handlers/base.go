package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mauricioabreu/url-shortener/internal/api/serializers"
)

func DoHealthcheck(c *gin.Context) {
	serializers.RespondWithSuccess(c, http.StatusOK, "WORKING", nil)
}
