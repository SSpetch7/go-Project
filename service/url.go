package service

type OriginalURLResponse struct {
	Id          int64
	OriginalURL string
	UserId      int64
	HashURL     string
}

type OriginalURLInsert struct {
	OriginalURL string
	ShortURL    string
	UserID      int
}

type HashURLResponse struct {
	longHash  string
	shortHash string
}

type URLService interface {
	CreateShortURL(longURL string, userId int) (string, error)
	HashURL(URL *HashURLResponse) (*HashURLResponse, error)
	GetOriginalURL(shortURL string) (*OriginalURLResponse, error)
}
