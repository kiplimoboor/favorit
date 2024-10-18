package database

import (
	"errors"
	"fmt"

	"github.com/kiplimoboor/favorit/models"
)

func (db *SQLiteDB) CreateRoom(rq models.RoomRequest) error {
	checkExists := "SELECT number FROM rooms WHERE number=?"
	row := db.db.QueryRow(checkExists, rq.Number)
	var result string
	if err := row.Scan(&result); err == nil {
		return fmt.Errorf("room %s already exists", rq.Number)
	}
	rm := models.NewRoom(rq)

	stmt := "INSERT INTO rooms(number,size,description,booked,created_at) VALUES(?,?,?,?,?)"
	_, err := db.db.Exec(stmt, rm.Number, rm.Size, rm.Description, rm.Booked, rm.CreatedAt)
	return err
}

func (db *SQLiteDB) GetRoomBy(key, value string) (*models.Room, error) {
	rm := new(models.Room)
	stmt := fmt.Sprintf("SELECT * FROM rooms WHERE %s=?", key)
	row := db.db.QueryRow(stmt, value)
	err := row.Scan(&rm.Id, &rm.Number, &rm.Size, &rm.Description, &rm.Booked, &rm.CreatedAt)
	if err != nil {
		return nil, errors.New("room not found")
	}
	return rm, nil
}

func (db *SQLiteDB) GetAllRooms() (*[]models.Room, error) {
	var rooms []models.Room
	stmt := "SELECT * FROM rooms;"
	rows, err := db.db.Query(stmt)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		rm := models.Room{}
		err := rows.Scan(&rm.Id, &rm.Number, &rm.Size, &rm.Description, &rm.Booked, &rm.CreatedAt)
		if err != nil {
			return nil, err
		}
		rooms = append(rooms, rm)
	}
	return &rooms, nil
}

func (db *SQLiteDB) UpdateRoom(number, field string, newValue any) error {
	if _, err := db.GetRoomBy("number", number); err != nil {
		return err
	}
	stmt := fmt.Sprintf("UPDATE rooms SET %s=? WHERE number=?", field)
	_, err := db.db.Exec(stmt, newValue, number)
	return err
}

func (db *SQLiteDB) DeleteRoom(number string) error {
	if _, err := db.GetRoomBy("number", number); err != nil {
		return err
	}
	stmt := "DELETE FROM rooms WHERE number=?"
	_, err := db.db.Exec(stmt, number)
	return err
}
