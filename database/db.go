package database

import (
	"database/sql"
	"os"

	"github.com/kiplimoboor/favorit/database/models"
	_ "github.com/mattn/go-sqlite3"
)

type Database interface {
	Init() error

	CreateUser(models.User) error
	GetUserBy(key, value string) (*models.User, error)
	UpdateUser(email, field, newValue string) error
	DeleteUser(email string) error
}

type SQLiteDB struct {
	db *sql.DB
}

func NewSQLiteDB() (*SQLiteDB, error) {
	db, err := sql.Open("sqlite3", "database/favorit.db")
	if err != nil {
		return nil, err
	}
	return &SQLiteDB{db: db}, nil
}

func (db *SQLiteDB) Init() error {
	os.Remove("database/favorit.db")
	tableQueries := []string{models.CreateUserTableQuery}

	for _, v := range tableQueries {
		_, err := db.db.Exec(v)
		if err != nil {
			return err
		}
	}
	return nil
}
