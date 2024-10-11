package models

import "time"

type Room struct {
	Id          int
	Booked      bool
	Description string
	Number      string
	Size        string
	CreatedAt   time.Time
}

type CreateRoomRequest struct {
	Description string `json:"description"`
	Number      string `json:"number"`
	Size        string `json:"size"`
}

type GetRoomResponse struct {
	Booked      bool   `json:"booked"`
	Description string `json:"description"`
	Number      string `json:"number"`
	Size        string `json:"size"`
}

func NewRoom(r CreateRoomRequest) *Room {
	return &Room{
		Description: r.Description,
		Number:      r.Number,
		Size:        r.Size,
		CreatedAt:   time.Now().UTC(),
	}
}

const CreateRoomTableQuery string = `
CREATE TABLE IF NOT EXISTS rooms
(
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    number VARCHAR(10) UNIQUE,
    size VARCHAR(50),
    description TEXT,
    booked BOOLEAN DEFAULT 0,
    created_at TIMESTAMP
)`
