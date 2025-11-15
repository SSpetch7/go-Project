package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type userRepositoryDB struct {
	db *sqlx.DB
}

func NewUserRepositoryDB(db *sqlx.DB) userRepositoryDB {
	return userRepositoryDB{db: db}
}

func (r userRepositoryDB) RegisterUser(body *NewUserRequest) ([]User, error) {

	// user := User{}
	query := `INSERT INTO users (username, password, email) VALUE( ?, ?, ?)`

	// สร้าง token
	_, err := r.db.Exec(query, body.Username, body.Password, body.Email)

	if err != nil {
		return nil, err
	}

	user, err := r.GetUserByEmail(body.Email)

	if err != nil {
		return nil, err
	}

	fmt.Println("repo regis return row", user)

	return user, nil
}

func (r userRepositoryDB) GetAll() ([]User, error) {

	users := []User{}
	query := "select id, username, email, role_id, create_at, update_at from users"

	err := r.db.Select(&users, query)

	if err != nil {
		return nil, err
	}

	return users, nil
}

func (r userRepositoryDB) GetUserByEmail(email string) ([]User, error) {
	user := []User{}
	query := "SELECT id, username, email, role_id, create_at, update_at FROM users WHERE email = ?"
	err := r.db.Select(&user, query, email)

	if err != nil {
		return nil, err
	}

	return user, nil

}

// func (r userRepositoryDB) GetbyId(id int) (*User, error) {

// 	customer := User{}
// 	query := "select * from users where customer_id = ?"

// 	err := r.db.Get(&customer, query, id)

// 	if err != nil {
// 		return nil, err
// 	}
// 	return &customer, nil
// }
