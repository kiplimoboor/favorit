package main

import (
	"database/sql"
	"fmt"

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

func (s *SQLiteDB) Init() error {
	s.db.Exec("DROP TABLE users")
	return s.createUserTable()
}

func (s *SQLiteDB) AddUser(u *User) error {
	stmt := "INSERT INTO users(id,first_name,last_name,email,username,password,created_at) VALUES(?,?,?,?,?,?,?);"
	_, err := s.db.Exec(stmt, u.ID, u.FirstName, u.LastName, u.Email, u.UserName, u.Password, u.CreatedAt)
	return err
}

func (s *SQLiteDB) GetUserBy(key, value string) (*User, error) {
	stmt := fmt.Sprintf("SELECT * FROM users WHERE %s=?;", key)
	row := s.db.QueryRow(stmt, value)
	u := new(User)
	err := row.Scan(&u.ID, &u.FirstName, &u.LastName, &u.Email, &u.UserName, &u.Password, &u.CreatedAt)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (s *SQLiteDB) createUserTable() error {
	_, err := s.db.Exec(CreateUserTableQuery)
	return err
}
