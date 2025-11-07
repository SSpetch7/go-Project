package repository

import "time"

type User struct {
	UserID   int       `db:"id"`
	Username string    `db:"username"`
	Email    string    `db:"email"`
	RoleId   string    `db:"role_id"`
	CreateAt time.Time `db:"create_at"`
	UpdateAt time.Time `db:"update_at"`
}

type ResponNewUser struct {
	Massage string
	Status  string
}

type NewUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type UserRepository interface {
	RegisterUser(*NewUserRequest) (*User, error)
	GetAll() ([]User, error)
	GetUserByEmail(email string) ([]User, error)
	// GetbyId(int) (*User, error)
}
