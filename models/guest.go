package models

import "time"

type Guest struct {
	Id        int    `json:"-"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Address   string `json:"address"`
	CreatedAt int64  `json:"createdAt"`
}

type GuestRequest struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Address   string `json:"address"`
}

func NewGuest(gr GuestRequest) Guest {
	return Guest{
		FirstName: gr.FirstName,
		LastName:  gr.LastName,
		Email:     gr.Email,
		Phone:     gr.Phone,
		Address:   gr.Address,
		CreatedAt: time.Now().Unix(),
	}
}

const GuestTableQuery string = `
CREATE TABLE IF NOT EXISTS guests
(
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    first_name VARCHAR(50),
    last_name VARCHAR(50),
    email VARCHAR(50) UNIQUE,
    phone VARCHAR(15),
	  address VARCHAR(256),
    created_at INTEGER
)`
