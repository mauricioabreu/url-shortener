package handlers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/mauricioabreu/url-shortener/internal/api/handlers"
	"github.com/stretchr/testify/assert"
)

func TestShorten(t *testing.T) {
	gin.SetMode(gin.TestMode)

	tests := []struct {
		name       string
		input      handlers.ShortenRequest
		wantStatus int
	}{
		{
			name:       "valid request",
			input:      handlers.ShortenRequest{URL: "https://www.shortener.com"},
			wantStatus: http.StatusOK,
		},
		{
			name:       "invalid url",
			input:      handlers.ShortenRequest{URL: "not-a-valid-url"},
			wantStatus: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)

			jsonData, _ := json.Marshal(tt.input)
			c.Request = httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(jsonData))
			c.Request.Header.Set("Content-Type", "application/json")

			handlers.Shorten(c)

			assert.Equal(t, tt.wantStatus, w.Code)
		})
	}
}
