package models

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        int    `json:"-"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	UserName  string `json:"username"`
	Role      string `json:"role"`
	Email     string `json:"email"`
	Password  string `json:"-"`
	CreatedAt int64  `json:"createdAt"`
}

type UserRequest struct {
	FirstName string `json:"firstName,omitempty"`
	LastName  string `json:"lastName,omitempty"`
	UserName  string `json:"username,omitempty"`
	Email     string `json:"email,omitempty"`
	Role      string `json:"role"`
	Password  string `json:"password,omitempty"`
}

func NewUser(u UserRequest) *User {
	user := User{
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Email:     u.Email,
		UserName:  u.UserName,
		Role:      u.Role,
		Password:  hashPassword(u.Password),
		CreatedAt: time.Now().Unix(),
	}
	return &user
}

func hashPassword(password string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), 4)
	return string(bytes)
}

const UserTableQuery string = `
CREATE TABLE IF NOT EXISTS users
(
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    first_name VARCHAR(50),
    last_name VARCHAR(50),
    username VARCHAR(50) UNIQUE,
    email VARCHAR(50) UNIQUE,
	  role VARCHAR(50),
    password VARCHAR(50),
    created_at INTEGER
)`
