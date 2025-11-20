package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type urlRepositoryDB struct {
	db *sqlx.DB
}

func NewURLRepository(db *sqlx.DB) urlRepositoryDB {
	return urlRepositoryDB{db: db}
}

func (r urlRepositoryDB) InsertURL(body *OriginalURLInsert) error {
	fmt.Println("body", body)

	query := `INSERT INTO URL_store (origin_url, short_url, user_id) VALUE( ?, ?, ?)`

	_, err := r.db.Exec(query, body.OriginalURL, body.ShortURL, body.UserID)

	if err != nil {
		return err
	}

	return nil
}

func (r urlRepositoryDB) GetOriginURL(shortURL string) (*URLResponse, error) {

	urlRes := []URLResponse{}

	query := "SELECT id, origin_url, user_id from URL_store WHERE short_url = ?"

	err := r.db.Select(&urlRes, query, shortURL)

	if err != nil {
		return nil, err
	}

	return &urlRes[0], nil
}
