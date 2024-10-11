package database

import (
	"database/sql"
	"os"

	"github.com/kiplimoboor/favorit/models"
	_ "github.com/mattn/go-sqlite3"
)

type Database interface {
	Init() error

	CreateUser(models.User) error
	GetUserBy(key, value string) (*models.User, error)
	UpdateUser(username, field, newValue string) error
	DeleteUser(username string) error

	CreateRoom(models.Room) error
	GetRoomBy(key, value string) (*models.Room, error)
	UpdateRoom(number, field string, newValue any) error
	DeleteRoom(number string) error
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
	tableQueries := []string{
		models.CreateUserTableQuery,
		models.CreateRoomTableQuery,
	}

	for _, v := range tableQueries {
		_, err := db.db.Exec(v)
		if err != nil {
			return err
		}
	}
	return nil
}
