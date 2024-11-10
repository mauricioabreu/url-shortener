package url

import (
	"context"
	"crypto/sha256"
	"encoding/base64"
	"strings"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

const shortURLLength = 8

type URLService struct {
	db *pgxpool.Pool
}

func NewURLService(db *pgxpool.Pool) *URLService {
	return &URLService{db: db}
}

type ShortenData struct {
	LongURL string
	Tags    []string
}

type ShortenResponse struct {
	ID        int
	LongURL   string
	ShortURL  string
	Tags      []string
	CreatedAt string
}

func (s *URLService) Shorten(ctx context.Context, data *ShortenData) (*ShortenResponse, error) {
	hash := sha256.Sum256([]byte(data.LongURL))
	shortURL := base64.URLEncoding.EncodeToString(hash[:])
	shortURL = strings.TrimRight(shortURL, "=")[:shortURLLength]

	tx, err := s.db.Begin(ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback(ctx)

	var urlID int
	err = tx.QueryRow(ctx,
		`INSERT INTO urls (short_url, long_url, tags)
		VALUES ($1, $2, $3)`,
		shortURL, data.LongURL, data.Tags).
		Scan(&urlID)
	if err != nil {
		return nil, err
	}

	for _, tag := range data.Tags {
		var tagID int
		err = tx.QueryRow(ctx,
			`INSERT INTO tags (name)
		VALUES ($1)
		ON CONFLICT (name) DO UPDATE SET = EXCLUDED.name
		RETURNING id`,
			tag).Scan(&tagID)
		if err != nil {
			return nil, err
		}

		_, err = tx.Exec(ctx,
			`INSERT INTO url_tags (url_id, tag_id)
			VALUES ($1, $2) 
			ON CONFLICT DO NOTHING`,
			urlID, tagID)
		if err != nil {
			return nil, err
		}
	}

	if err = tx.Commit(ctx); err != nil {
		return nil, err
	}

	shortenResponse := &ShortenResponse{
		ID:        urlID,
		LongURL:   data.LongURL,
		ShortURL:  shortURL,
		Tags:      data.Tags,
		CreatedAt: time.Now().Format(time.RFC3339),
	}

	return shortenResponse, nil
}
