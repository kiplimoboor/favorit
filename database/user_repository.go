package database

import (
	"errors"
	"fmt"
	"github.com/kiplimoboor/favorit/database/models"
)

type UserRepository struct {
	store *SQLiteDB
}

func NewUserRepository(db *SQLiteDB) *UserRepository {
	return &UserRepository{store: db}
}

func (r *UserRepository) CreateUser(u models.User) error {
	stmt := "INSERT INTO users(id,first_name,last_name,email,username,password,created_at) VALUES(?,?,?,?,?,?,?);"
	_, err := r.store.db.Exec(stmt, u.ID, u.FirstName, u.LastName, u.Email, u.UserName, u.Password, u.CreatedAt)
	return err
}

func (r *UserRepository) GetUserBy(key, value string) (*models.User, error) {
	stmt := fmt.Sprintf("SELECT * FROM users WHERE %s=?;", key)
	row := r.store.db.QueryRow(stmt, value)
	u := new(models.User)
	err := row.Scan(&u.ID, &u.FirstName, &u.LastName, &u.Email, &u.UserName, &u.Password, &u.CreatedAt)
	if err != nil {
		return nil, err
	}
	return u, nil
}

// Only updates based on email
func (r *UserRepository) UpdateUser(email, field, newValue string) error {
	_, err := r.GetUserBy("email", email)
	if err != nil {
		return errors.New(fmt.Sprintf("user with email %s does not exist", email))
	}
	stmt := fmt.Sprintf("UPDATE users SET %s=? WHERE email=?;", field)
	_, err = r.store.db.Exec(stmt, newValue, email)
	return err
}

// Only deletes based on email
func (r *UserRepository) DeleteUser(email string) error {
	_, err := r.GetUserBy("email", email)
	if err != nil {
		return errors.New(fmt.Sprintf("user with email %s does not exist", email))
	}

	stmt := "DELETE FROM users WHERE email=?;"
	_, err = r.store.db.Exec(stmt, email)
	return err
}
