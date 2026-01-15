package repository

import (
	"database/sql"
	"fmt"
)

type URLRepository struct {
	db *sql.DB
}

func NewURLRepository(db *sql.DB) *URLRepository {
	return &URLRepository{db: db}
}

func (r *URLRepository) Create(originalURL string) (int64, string, error) {
	var id int64

	err := r.db.QueryRow(`SELECT nextval('urls_id_seq')`).Scan(&id)
	if err != nil {
		return 0, "", err
	}

	short := encodeBase62(uint64(10000 + id))

	query := `
		INSERT INTO urls (id, original_url, short_url)
		VALUES ($1, $2, $3)
	`
	_, err = r.db.Exec(query, id, originalURL, short)
	if err != nil {
		return 0, "", err
	}

	return id, short, nil
}

func (r *URLRepository) Find(shortCode string) (string, error) {
	var originalURL string

	err := r.db.QueryRow(fmt.Sprintf("SELECT original_url FROM urls WHERE short_url = '%s'", shortCode)).Scan(&originalURL)
	if err != nil {
		return "", err
	}

	return originalURL, nil
}
