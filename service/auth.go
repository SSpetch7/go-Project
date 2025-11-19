package service

import (
	"context"
)

type PayloadToken struct {
	UserID   int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	RoleId   string `json:"role_id"`
	CreateAt string `json:"create_at"`
	UpdateAt string `json:"update_at"`
}

type AuthService interface {
	CreateToken(ctx context.Context, data PayloadToken) (string, error)
	VerifyToken(ctx context.Context, token string) (*PayloadToken, error)
}
