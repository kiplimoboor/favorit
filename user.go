package main

import (
	"errors"
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

type CreateUserRequest struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	UserName  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"password"`
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

func CreateNewUser(u CreateUserRequest) (*User, error) {
	err := ValidateUserCreateRequest(u)
	if err != nil {
		return nil, err
	}
	user := User{
		ID:        uuid.New().String(),
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Email:     u.Email,
		UserName:  u.UserName,
		Password:  HashPassword(u.Password),
		CreatedAt: time.Now().UTC(),
	}
	return &user, nil
}

func ValidateUserCreateRequest(u CreateUserRequest) error {
	if u.FirstName == "" {
		return errors.New("check firstName field")
	}
	if u.LastName == "" {
		return errors.New("check lastName field")
	}
	if u.Email == "" {
		return errors.New("check email field")
	}
	if u.Password == "" {
		return errors.New("check password field")
	}
	return nil
}

func HashPassword(password string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), 4)
	return string(bytes)
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
