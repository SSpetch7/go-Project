package service

import (
	"context"
	"errors"
	"fmt"
	"go-project/repository"
	r "go-project/repository"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) userService {
	return userService{userRepo: userRepo}
}

func (s userService) GetUsers() ([]UserResponse, error) {
	users, err := s.userRepo.GetAll()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	userResponses := []UserResponse{}

	for _, user := range users {
		userResponse := UserResponse{
			UserID:   user.UserID,
			Username: user.Username,
			Email:    user.Email,
			RoleId:   user.RoleId,
			CreateAt: user.CreateAt,
			UpdateAt: user.UpdateAt,
		}
		userResponses = append(userResponses, userResponse)
	}
	return userResponses, nil

}

func (s userService) RegisterUser(ctx context.Context, newUser *r.NewUserRequest) (*UserResponse, error) {

	dupEmail, err := s.userRepo.GetUserByEmail(newUser.Email)

	if err != nil {
		return nil, err
	}

	if len(dupEmail) > 0 {
		return nil, errors.New("email already exists")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)

	if err != nil {
		return nil, err
	}

	newUser.Password = string(hashedPassword)

	user, err := s.userRepo.RegisterUser(newUser)

	if err != nil {
		return nil, err
	}

	userResponses := UserResponse{}

	for _, user := range user {
		userResponse := UserResponse{
			UserID:   user.UserID,
			Username: user.Username,
			Email:    user.Email,
			RoleId:   user.RoleId,
			CreateAt: user.CreateAt,
			UpdateAt: user.UpdateAt,
		}
		userResponses = userResponse
	}

	return &userResponses, nil
}

func (s userService) Login(ctx context.Context, email, password string) (string, error) {
	fmt.Println("====== user_service / login =======")
	user, err := s.userRepo.GetUserByEmail(email)

	if err != nil {
		return "", err
	}

	if len(user) < 1 {
		return "", errors.New("login is fail")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user[0].Password), []byte(password))

	if err != nil {
		return "", errors.New("login is fail")
	}

	// generate token

	jwtSecretKey := "TESTsecret"
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = user[0].UserID
	claims["username"] = user[0].Username
	claims["email"] = user[0].Email
	claims["role"] = user[0].RoleId
	claims["createAt"] = user[0].CreateAt
	claims["updateAt"] = user[0].UpdateAt
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	accessToken, err := token.SignedString([]byte(jwtSecretKey))
	if err != nil {
		return "", err
	}

	fmt.Println("====== user_service / login =======")
	return accessToken, nil

}
