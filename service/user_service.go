package service

import (
	"context"
	"errors"
	"fmt"
	"go-project/repository"
	r "go-project/repository"
	"log"

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
		fmt.Println("err", err)
		return nil, err
	}

	if len(dupEmail) > 0 {
		return nil, errors.New("email already exists")
	}

	fmt.Println("not dupEmail", dupEmail)

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
