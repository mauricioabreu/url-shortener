package handlers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/mauricioabreu/url-shortener/internal/api/handlers"
	"github.com/mauricioabreu/url-shortener/internal/services/url"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"go.uber.org/zap"
)

func TestShorten(t *testing.T) {
	gin.SetMode(gin.TestMode)

	router := gin.New()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := handlers.NewMockShortenerService(ctrl)
	logger, _ := zap.NewDevelopment()

	h := handlers.NewShortenerHandler(mockService, logger)
	router.POST("/", h.Shorten)

	t.Run("valid request", func(t *testing.T) {
		w := httptest.NewRecorder()
		resp := &url.ShortenResponse{
			ShortURL: "https://shortener.com/abc123xy",
		}
		mockService.EXPECT().Shorten(gomock.Any(), gomock.Any()).Return(resp, nil)
		input := handlers.ShortenRequest{URL: "https://www.google.com"}
		jsonData, _ := json.Marshal(input)

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(jsonData))
		req.Header.Set("Content-Type", "application/json")
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("valid request with tags", func(t *testing.T) {
		w := httptest.NewRecorder()
		resp := &url.ShortenResponse{
			ShortURL: "https://shortener.com/abc123xy",
		}
		mockService.EXPECT().Shorten(gomock.Any(), gomock.Any()).Return(resp, nil)
		input := handlers.ShortenRequest{URL: "https://www.google.com", Tags: []string{"tag1", "tag2"}}
		jsonData, _ := json.Marshal(input)

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(jsonData))
		req.Header.Set("Content-Type", "application/json")
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("invalid tags", func(t *testing.T) {
		w := httptest.NewRecorder()
		input := handlers.ShortenRequest{URL: "https://www.google.com", Tags: []string{"invalid-tag!"}}
		jsonData, _ := json.Marshal(input)

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(jsonData))
		req.Header.Set("Content-Type", "application/json")
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("invalid url", func(t *testing.T) {
		w := httptest.NewRecorder()
		input := handlers.ShortenRequest{URL: "not-a-valid-url"}
		jsonData, _ := json.Marshal(input)

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(jsonData))
		req.Header.Set("Content-Type", "application/json")
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("database error", func(t *testing.T) {
		w := httptest.NewRecorder()
		mockService.EXPECT().Shorten(gomock.Any(), gomock.Any()).Return(nil, assert.AnError)
		input := handlers.ShortenRequest{URL: "https://www.google.com"}
		jsonData, _ := json.Marshal(input)

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(jsonData))
		req.Header.Set("Content-Type", "application/json")
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})
}
