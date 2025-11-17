package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/golang-jwt/jwt/v4"
	"github.com/spf13/viper"
)

type authService struct {
	ts AuthService
}

func NewAuthService() authService {
	return authService{}
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
		return nil, err
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
