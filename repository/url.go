package repository

type URLResponse struct {
	Id      int64  `db:"id"`
	LongURL string `db:"origin_url"`
	UserID  int64  `db:"user_id"`
}

type OriginalURLInsert struct {
	OriginalURL string `json:"long_url"`
	ShortURL    string
	UserID      int
}

type URLRepository interface {
	InsertURL(body *OriginalURLInsert) error
	GetOriginURL(shortURL string) (*URLResponse, error)
}
