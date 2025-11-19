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

func (r urlRepositoryDB) InsertURL(body *NewURLRequest) error {
	fmt.Println("body", body)
	return nil
}
