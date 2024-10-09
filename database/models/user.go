package models

import (
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        string
	FirstName string
	LastName  string
	UserName  string
	Email     string
	Password  string
	CreatedAt time.Time
}

type UserRequest struct {
	FirstName string `json:"firstName,omitempty"`
	LastName  string `json:"lastName,omitempty"`
	UserName  string `json:"username,omitempty"`
	Email     string `json:"email,omitempty"`
	Password  string `json:"password,omitempty"`
}

type GetUserResponse struct {
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	UserName  string    `json:"username"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"createdAt"`
}

type UpdateUser struct {
	Email    string `json:"email"`
	Field    string `json:"field"`
	NewValue string `json:"newVal"`
}

func CreateNewUser(u UserRequest) (*User, error) {
	user := User{
		ID:        uuid.New().String(),
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Email:     u.Email,
		UserName:  u.UserName,
		Password:  hashPassword(u.Password),
		CreatedAt: time.Now().UTC(),
	}
	return &user, nil
}

func hashPassword(password string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), 4)
	return string(bytes)
}

const CreateUserTableQuery string = `
CREATE TABLE IF NOT EXISTS users
(
    id VARCHAR PRIMARY KEY,
    first_name VARCHAR(50),
    last_name VARCHAR(50),
    username VARCHAR(50) UNIQUE,
    email VARCHAR(50) UNIQUE,
    password VARCHAR(50),
    created_at TIMESTAMP
)`
