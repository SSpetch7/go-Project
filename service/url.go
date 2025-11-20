package service

type OriginalURLResponse struct {
	Id          int64  `json:"id"`
	OriginalURL string `json:"original_url"`
	UserId      int64  `json:"user_id"`
}

type OriginalURLInsert struct {
	OriginalURL string
	ShortURL    string
	UserID      int
}

type URLService interface {
	CreateShortURL(longURL string, userId int) (string, error)
	HashURL(longURL string) (string, error)
	GetOriginalURL(shortURL string) (*OriginalURLResponse, error)
}
