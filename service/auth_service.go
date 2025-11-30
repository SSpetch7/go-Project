package service

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/spf13/viper"
)

type authService struct {
	ts AuthService
}

func NewAuthService() authService {
	return authService{}
}

func (auth authService) CreateToken(ctx context.Context, data PayloadToken) (string, error) {

	jwtSecretKey := viper.GetString("env.jwtSecretKey")

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = data.UserID
	claims["username"] = data.Username
	claims["email"] = data.Email
	claims["role"] = data.RoleId
	claims["createAt"] = data.CreateAt
	claims["updateAt"] = data.UpdateAt
	claims["exp"] = time.Now().Add(time.Hour * 12).Unix()

	accessToken, err := token.SignedString([]byte(jwtSecretKey))
	if err != nil {
		return "", err
	}

	return accessToken, nil
}

func (auth authService) VerifyToken(ctx context.Context, tokenString string) (*PayloadToken, error) {
	fmt.Println("====== auth_service / VerifyToken =======")

	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		// Validate signing method
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(viper.GetString("env.jwtSecretKey")), nil
	})

	if err != nil {
		if err.Error() == "Token is expired" {
			return nil, errors.New("token is expired")
		}
		return nil, errors.New("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok {
		return nil, errors.New("invalid token")
	}

	payloadToken := PayloadToken{
		UserID:   int(claims["id"].(float64)),
		Username: claims["username"].(string),
		Email:    claims["email"].(string),
		RoleId:   claims["role"].(string),
		CreateAt: claims["createAt"].(string),
		UpdateAt: claims["updateAt"].(string),
	}

	return &payloadToken, nil

}
