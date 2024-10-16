package database

import (
	"errors"
	"time"

	"github.com/kiplimoboor/favorit/models"
)

func (db *SQLiteDB) CreateBooking(br models.BookingRequest) error {
	if br.CheckIn+20 < time.Now().Unix() {
		return errors.New("check in date cannot be in the past")
	}
	if br.CheckOut <= br.CheckIn {
		return errors.New("check out date must be greater than check in date")
	}
	if _, err := db.GetRoomBy("number", br.RoomNumber); err != nil {
		return err
	}
	if _, err := db.GetGuestBy("email", br.GuestEmail); err != nil {
		return err
	}
	b := models.NewBooking(br)
	stmt := "INSERT INTO bookings (guest_email,room_number,checkin,checkout,created_at) VALUES(?,?,?,?,?)"
	if _, err := db.db.Exec(stmt, b.GuestEmail, b.RoomNumber, b.CheckIn, b.CheckOut, b.CreatedAt); err != nil {
		return err
	}

	stmt = "UPDATE rooms SET booked=1 WHERE number=?"
	_, err := db.db.Exec(stmt, b.RoomNumber)
	return err
}
