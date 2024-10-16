package models

import "time"

type Room struct {
	Id          int    `json:"id"`
	Booked      bool   `json:"booked"`
	Description string `json:"desc"`
	Number      string `json:"number"`
	Size        string `json:"size"`
	CreatedAt   int64  `json:"createdAt"`
}

type RoomRequest struct {
	Description string `json:"description"`
	Number      string `json:"number"`
	Size        string `json:"size"`
}

func NewRoom(r RoomRequest) *Room {
	return &Room{
		Description: r.Description,
		Number:      r.Number,
		Size:        r.Size,
		CreatedAt:   time.Now().Unix(),
	}
}

const RoomTableQuery string = `
CREATE TABLE IF NOT EXISTS rooms
(
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    number VARCHAR(10) UNIQUE,
    size VARCHAR(50),
    description TEXT,
    booked BOOLEAN DEFAULT 0,
    created_at INTEGER
)`
