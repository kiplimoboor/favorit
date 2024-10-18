package database

import (
	"database/sql"
	// "os"

	"github.com/kiplimoboor/favorit/models"
	_ "github.com/mattn/go-sqlite3"
)

type Database interface {
	Init() error

	// User Management
	CreateUser(models.UserRequest) error
	GetUserBy(key, value string) (*models.User, error)
	GetAllUsers() (*[]models.User, error)
	UpdateUser(username, field, newValue string) error
	DeleteUser(username string) error

	// Room Management
	CreateRoom(models.RoomRequest) error
	GetRoomBy(key, value string) (*models.Room, error)
	GetAllRooms() (*[]models.Room, error)
	UpdateRoom(number, field string, newValue any) error
	DeleteRoom(number string) error

	// Guest Management
	CreateGuest(models.GuestRequest) error
	GetGuestBy(key, value string) (*models.Guest, error)
	GetAllGuests() (*[]models.Guest, error)
	UpdateGuest(email, field, newVal string) error

	// Booking Management
	CreateBooking(br models.BookingRequest) error
	GetBookingBy(field string, value any) (*models.Booking, error)
	GetAllBookings() (*[]models.Booking, error)
	UpdateBooking(id int, field, newVal any) error
	Vacate(id int, reason string) error
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
	// os.Remove("database/favorit.db")
	tableQueries := []string{
		"PRAGMA foreign_keys = ON",
		models.UserTableQuery,
		models.RoomTableQuery,
		models.GuestTableQuery,
		models.BookingTableQuery,
	}

	for _, v := range tableQueries {
		_, err := db.db.Exec(v)
		if err != nil {
			panic(err)
		}
	}
	return nil
}
