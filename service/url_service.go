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

func (s urlService) CreateShortURL(longURL string, userId int) (string, error) {
	isNotExist := true

	shortURL := longURL
	var err error

	for isNotExist {
		shortURL, err = s.HashURL(shortURL)
		result, errOrignal := s.urlRepo.GetOriginURL(shortURL)

		if errOrignal != nil {
			return "", errOrignal
		}

		fmt.Println("result", result)

		isNotExist = result != nil

	}
	if err != nil {
		fmt.Println("isNotExist err: ", err)
		return "", err
	}

	// shortURL, err := s.HashURL(longURL)

	// if err != nil {
	// 	return "", err
	// }

	// return shortURL, nil
	// item := OriginalURLInsert{}

	item := repository.OriginalURLInsert{
		OriginalURL: longURL,
		ShortURL:    shortURL,
		UserID:      userId,
	}

	err = s.urlRepo.InsertURL(&item)

	if err != nil {
		fmt.Println("InsertURL err: ", err)
		return "", err
	}

	return "", nil
}

func (s urlService) HashURL(longURL string) (string, error) {
	hash := sha256.Sum256([]byte(longURL))
	encoded := base64.URLEncoding.EncodeToString(hash[:])

	shortHash, longHash := encoded[:7], encoded

	fmt.Println("shortHash", shortHash)
	fmt.Println("longHash", longHash)

	return shortHash, nil
}

func (s urlService) GetOriginalURL(shortURL string) (*OriginalURLResponse, error) {
	return nil, nil
}
