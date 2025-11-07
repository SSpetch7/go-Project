// เป็น port ของฝั่ง service
package service

import (
	"context"
	r "go-project/repository"
	"time"
)

// retrun เป็น json format
type UserResponse struct {
	UserID   int       `json:"id"`
	Username string    `json:"username"`
	Email    string    `json:"email"`
	RoleId   string    `json:"role_id"`
	CreateAt time.Time `json:"create_at"`
	UpdateAt time.Time `json:"update_at"`
}

type NewUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type UserService interface {
	RegisterUser(ctx context.Context, newUser *r.NewUserRequest) (*UserResponse, error)
	GetUsers() ([]UserResponse, error)
	// GetUser(int) (*UserResponse, error)
}
