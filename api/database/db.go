package database

import (
	"database/sql"

	"github.com/kiplimoboor/favorit/api/database/models"
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
	tableQueries := []string{models.CreateUserTableQuery}

	for _, v := range tableQueries {
		_, err := sqlite.db.Exec(v)
		if err != nil {
			return err
		}
	}
	return nil
}
