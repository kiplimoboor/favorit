package models

import "time"

type Booking struct {
	Id         int    `json:"id"`
	GuestEmail string `json:"guestEmail"`
	RoomNumber string `json:"roomNumber"`
	CheckIn    int64  `json:"checkIn"`
	CheckOut   int64  `json:"checkOut"`
	Status     string `json:"status"`
	CreatedAt  int64  `json:"createdAt"`
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
		Status:     "occupied",
		CreatedAt:  time.Now().Unix(),
	}
}

// booking status is one of occupied, cancelled, checkout
const BookingTableQuery string = `
CREATE TABLE IF NOT EXISTS bookings
(
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    guest_email VARCHAR(50) NOT NULL,
    room_number VARCHAR(10) NOT NULL,
    checkin INTEGER,
    checkout INTEGER,
	  status VARCHAR(15),
    created_at INTEGER,

    FOREIGN KEY(guest_email) REFERENCES guests(email),
	  FOREIGN KEY(room_number) REFERENCES rooms(number)
)`
