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
	urlDup := ""

	itemInsert := repository.OriginalURLInsert{}

	hashURL := &HashURLResponse{
		longHash:  longURL,
		shortHash: longURL,
	}
	var err error
	for isNotExist {
		hashURL, err = s.HashURL(hashURL)

		if err != nil {
			return "", err
		}

		result, errOrignal := s.urlRepo.GetOriginURL(hashURL.shortHash)

		if errOrignal != nil && errOrignal.Error() != "not found data" {
			return "", errOrignal
		}

		if result == nil {
			isNotExist = false
		} else if result.LongURL == longURL {
			isNotExist = false
			urlDup = result.LongURL
		}

	}

	itemInsert = repository.OriginalURLInsert{
		OriginalURL: longURL,
		ShortURL:    hashURL.shortHash,
		UserID:      userId,
	}

	if itemInsert.OriginalURL == urlDup {
		return itemInsert.ShortURL, nil
	}

	err = s.urlRepo.InsertURL(&itemInsert)

	if err != nil {
		fmt.Println("InsertURL err: ", err)
		return "", err
	}

	return itemInsert.ShortURL, nil
}

func (s urlService) HashURL(URL *HashURLResponse) (*HashURLResponse, error) {
	hash := sha256.Sum256([]byte(URL.longHash))
	encoded := base64.URLEncoding.EncodeToString(hash[:])

	shortHash, longHash := encoded[:7], encoded

	resHash := HashURLResponse{
		longHash:  longHash,
		shortHash: shortHash,
	}

	return &resHash, nil
}

func (s urlService) GetOriginalURL(shortURL string) (*OriginalURLResponse, error) {

	originRes := &repository.URLResponse{}

	originRes, err := s.urlRepo.GetOriginURL(shortURL)

	if err != nil {
		return nil, err
	}

	response := &OriginalURLResponse{
		Id:          originRes.Id,
		UserId:      originRes.UserID,
		OriginalURL: originRes.LongURL,
		HashURL:     originRes.HashURL,
	}

	return response, nil
}
