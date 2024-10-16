package database

import (
	"fmt"

	"github.com/kiplimoboor/favorit/models"
)

func (db *SQLiteDB) CreateUser(ur models.UserRequest) error {
	userExistStmt := "SELECT username, email FROM users where username=? OR email=?"
	row := db.db.QueryRow(userExistStmt, ur.UserName, ur.Email)
	var existingUsername, existingEmail string
	row.Scan(&existingUsername, &existingEmail)
	if existingUsername == ur.UserName {
		return fmt.Errorf("username %s is already in use", existingUsername)
	}
	if existingEmail == ur.Email {
		return fmt.Errorf("email %s is already in use", existingEmail)
	}
	u := models.NewUser(ur)
	stmt := "INSERT INTO users(first_name,last_name,username,email,role,password,created_at) VALUES(?,?,?,?,?,?,?);"
	_, err := db.db.Exec(stmt, u.FirstName, u.LastName, u.UserName, u.Email, u.Role, u.Password, u.CreatedAt)
	return err
}

func (db *SQLiteDB) GetUserBy(key, value string) (*models.User, error) {
	stmt := fmt.Sprintf("SELECT * FROM users WHERE %s=?;", key)
	row := db.db.QueryRow(stmt, value)
	u := new(models.User)
	err := row.Scan(&u.ID, &u.FirstName, &u.LastName, &u.UserName, &u.Email, &u.Role, &u.Password, &u.CreatedAt)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (db *SQLiteDB) UpdateUser(username, field, newValue string) error {
	if _, err := db.GetUserBy("username", username); err != nil {
		return err
	}
	stmt := fmt.Sprintf("UPDATE users SET %s=? WHERE username=?;", field)
	_, err := db.db.Exec(stmt, newValue, username)
	return err
}

func (db *SQLiteDB) DeleteUser(username string) error {
	if _, err := db.GetUserBy("username", username); err != nil {
		return err
	}
	stmt := "DELETE FROM users WHERE username=?;"
	_, err := db.db.Exec(stmt, username)
	return err
}
