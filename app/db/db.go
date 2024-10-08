package app

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/kiplimoboor/favorit/models"
	_ "github.com/mattn/go-sqlite3"
)

type SQLiteDB struct {
	db *sql.DB
}

func NewSQLiteDB() (*SQLiteDB, error) {
	db, err := sql.Open("sqlite3", "favorit.db")
	if err != nil {
		return nil, err
	}
	return &SQLiteDB{db: db}, nil
}

func (sqlite *SQLiteDB) Init() error {
	sqlite.db.Exec("DROP TABLE users")
	return sqlite.createUserTable()
}

func (sqlite *SQLiteDB) AddUser(u models.User) error {
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

// This function only updates via email
func (sqlite *SQLiteDB) UpdateUser(email, field, newValue string) error {
	_, err := sqlite.GetUserBy("email", email)
	if err != nil {
		return errors.New(fmt.Sprintf("user with email %s does not exist", email))
	}
	stmt := fmt.Sprintf("UPDATE users SET %s=? WHERE email=?;", field)
	_, err = sqlite.db.Exec(stmt, newValue, email)
	return err
}

func (sqlite *SQLiteDB) DeleteUser(email string) error {
	_, err := sqlite.GetUserBy("email", email)
	if err != nil {
		return errors.New(fmt.Sprintf("user with email %s does not exist", email))
	}

	stmt := "DELETE FROM users WHERE email=?;"
	_, err = sqlite.db.Exec(stmt, email)
	return err
}

func (sqlite *SQLiteDB) createUserTable() error {
	_, err := sqlite.db.Exec(models.CreateUserTableQuery)
	return err
}
