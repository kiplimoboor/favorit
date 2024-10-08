package database

import (
	"errors"
	"fmt"
	"github.com/kiplimoboor/favorit/api/database/models"
)

func (sqlite *SQLiteDB) CreateUser(u models.User) error {
	stmt := "INSERT INTO users(id,first_name,last_name,email,username,password,created_at) VALUES(?,?,?,?,?,?,?);"
	_, err := sqlite.db.Exec(stmt, u.ID, u.FirstName, u.LastName, u.Email, u.UserName, u.Password, u.CreatedAt)
	return err
}

func (sqlite *SQLiteDB) GetUserBy(key, value string) (*models.User, error) {
	stmt := fmt.Sprintf("SELECT * FROM users WHERE %s=?;", key)
	row := sqlite.db.QueryRow(stmt, value)
	u := new(models.User)
	err := row.Scan(&u.ID, &u.FirstName, &u.LastName, &u.Email, &u.UserName, &u.Password, &u.CreatedAt)
	if err != nil {
		return nil, err
	}
	return u, nil
}

// Only updates based on email
func (sqlite *SQLiteDB) UpdateUser(email, field, newValue string) error {
	_, err := sqlite.GetUserBy("email", email)
	if err != nil {
		return errors.New(fmt.Sprintf("user with email %s does not exist", email))
	}
	stmt := fmt.Sprintf("UPDATE users SET %s=? WHERE email=?;", field)
	_, err = sqlite.db.Exec(stmt, newValue, email)
	return err
}

// Only deletes based on email
func (sqlite *SQLiteDB) DeleteUser(email string) error {
	_, err := sqlite.GetUserBy("email", email)
	if err != nil {
		return errors.New(fmt.Sprintf("user with email %s does not exist", email))
	}

	stmt := "DELETE FROM users WHERE email=?;"
	_, err = sqlite.db.Exec(stmt, email)
	return err
}
