package repository

type NewURLRequest struct {
	LongURL string `json:"long_url"`
}

type URLRepository interface {
	InsertURL(*NewURLRequest) error
}
