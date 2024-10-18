package database

import (
	"errors"
	"fmt"
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
	room, err := db.GetRoomBy("number", br.RoomNumber)
	if err != nil {
		return err
	}
	if room.Booked {
		return errors.New("room is unavailable")
	}
	if _, err := db.GetGuestBy("email", br.GuestEmail); err != nil {
		return err
	}
	b := models.NewBooking(br)
	stmt := "INSERT INTO bookings (guest_email,room_number,checkin,checkout,status,created_at) VALUES(?,?,?,?,?,?)"
	_, err = db.db.Exec(stmt, b.GuestEmail, b.RoomNumber, b.CheckIn, b.CheckOut, b.Status, b.CreatedAt)
	if err != nil {
		return err
	}
	stmt = "UPDATE rooms SET booked=1 WHERE number=?"
	_, err = db.db.Exec(stmt, b.RoomNumber)
	return err
}

func (db *SQLiteDB) GetAllBookings() (*[]models.Booking, error) {
	stmt := "SELECT * FROM bookings"
	rows, err := db.db.Query(stmt)
	if err != nil {
		return nil, err
	}
	var bookings []models.Booking

	for rows.Next() {
		b := models.Booking{}
		err := rows.Scan(&b.Id, &b.GuestEmail, &b.RoomNumber, &b.CheckIn, &b.CheckOut, &b.Status, &b.CreatedAt)
		if err != nil {
			return nil, err
		}
		bookings = append(bookings, b)
	}
	return &bookings, nil
}

func (db *SQLiteDB) GetBookingBy(field string, value any) (*models.Booking, error) {
	stmt := fmt.Sprintf("SELECT * FROM BOOKINGS WHERE %s=?", field)
	row := db.db.QueryRow(stmt, value)
	b := new(models.Booking)
	err := row.Scan(&b.Id, &b.GuestEmail, &b.RoomNumber, &b.CheckIn, &b.CheckOut, &b.Status, &b.CreatedAt)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func (db *SQLiteDB) UpdateBooking(id int, field, newVal any) error {
	if _, err := db.GetBookingBy("id", id); err != nil {
		return err
	}
	stmt := fmt.Sprintf("UPDATE bookings SET %s=? WHERE id=?", field)
	_, err := db.db.Exec(stmt, newVal, id)
	return err
}

// handles reasons such as chechout and cancellations
func (db *SQLiteDB) Vacate(id int, reason string) error {
	b, err := db.GetBookingBy("id", id)
	if err != nil {
		return err
	}
	stmt := "UPDATE rooms SET booked=0 WHERE number=?"
	_, err = db.db.Exec(stmt, b.RoomNumber)
	if err != nil {
		return err
	}
	stmt = "UPDATE bookings SET status=? WHERE id=?"
	_, err = db.db.Exec(stmt, reason, id)
	return err
}
