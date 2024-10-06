package main

import (
	"database/sql"

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
	stmt := "INSERT INTO users(id,first_name,last_name,email,password,created_at) VALUES(?,?,?,?,?,?)"

	_, err := s.db.Exec(stmt, u.ID, u.FirstName, u.LastName, u.Email, u.Password, u.CreatedAt)

	return err
}

func (s *SQLiteDB) GetUserByEmail(email string) (*User, error) {
	stmt := "SELECT * FROM users WHERE email=?"
	row := s.db.QueryRow(stmt, email)

	user := new(User)
	err := row.Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *SQLiteDB) createUserTable() error {
	query := `CREATE TABLE IF NOT EXISTS users(
	id VARCHAR PRIMARY KEY,
	first_name VARCHAR(50),
	last_name VARCHAR(50),
	email VARCHAR(50),
	password VARCHAR(50),
	created_at TIMESTAMP)`
	_, err := s.db.Exec(query)

	return err
}
