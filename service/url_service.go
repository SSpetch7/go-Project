package service

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"go-project/repository"
)

type urlService struct {
	urlRepo repository.URLRepository
}

func NewURLService(urlRepo repository.URLRepository) urlService {
	return urlService{urlRepo: urlRepo}
}

func (s urlService) HashURL(longURL string) (string, error) {
	hash := sha256.Sum256([]byte(longURL))
	encoded := base64.URLEncoding.EncodeToString(hash[:])

	shortHash, longHash := encoded[:7], encoded

	fmt.Println("shortHash", shortHash)
	fmt.Println("longHash", longHash)

	return shortHash, nil
}
