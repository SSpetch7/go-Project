package service

type URLService interface {
	HashURL(longURL string) (string, error)
}
