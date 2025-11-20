package service

type OriginalURLResponse struct {
	OriginalURL string `json:"original_url"`
	UserName    string `json:"username"`
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
