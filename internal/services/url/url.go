package url

import (
	"context"
	"crypto/sha256"
	"encoding/base64"
	"strings"

	"github.com/jackc/pgx/v5/pgxpool"
)

const shortURLLength = 8

type URLService struct {
	db *pgxpool.Pool
}

func NewURLService(db *pgxpool.Pool) *URLService {
	return &URLService{db: db}
}

func (s *URLService) Shorten(ctx context.Context, longURL string) (string, error) {
	hash := sha256.Sum256([]byte(longURL))
	shortURL := base64.URLEncoding.EncodeToString(hash[:])
	shortURL = strings.TrimRight(shortURL, "=")[:shortURLLength]

	_, err := s.db.Exec(ctx,
		"INSERT INTO urls (long_url, short_url) VALUES ($1, $2)",
		longURL, shortURL)
	if err != nil {
		return "", err
	}

	return shortURL, nil
}
