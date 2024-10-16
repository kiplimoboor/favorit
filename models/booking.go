package models

import "time"

type Booking struct {
	Id         int
	GuestEmail string
	RoomNumber string
	CheckIn    int64
	CheckOut   int64
	CreatedAt  int64
}

type BookingRequest struct {
	GuestEmail string `json:"guestEmail"`
	RoomNumber string `json:"roomNumber"`
	CheckIn    int64  `json:"checkIn"`
	CheckOut   int64  `json:"checkOut"`
}

func NewBooking(b BookingRequest) Booking {
	return Booking{
		GuestEmail: b.GuestEmail,
		RoomNumber: b.RoomNumber,
		CheckIn:    b.CheckIn,
		CheckOut:   b.CheckOut,
		CreatedAt:  time.Now().Unix(),
	}
}

const BookingTableQuery string = `
CREATE TABLE IF NOT EXISTS bookings
(
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    guest_email VARCHAR(50) NOT NULL,
    room_number VARCHAR(10) NOT NULL,
    checkin INTEGER,
    checkout INTEGER,
    created_at INTEGER,

    FOREIGN KEY(guest_email) REFERENCES guests(email),
	  FOREIGN KEY(room_number) REFERENCES rooms(number)
)`
