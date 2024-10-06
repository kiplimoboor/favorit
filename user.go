package main

import (
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        string
	FirstName string
	LastName  string
	Email     string
	Password  string
	CreatedAt time.Time
}

func NewUser(firstName, lastName, email, password string) *User {
	return &User{
		ID:        uuid.New().String(),
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		Password:  HashPassword(password),
		CreatedAt: time.Now(),
	}
}

func HashPassword(password string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), 4)
	return string(bytes)
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
