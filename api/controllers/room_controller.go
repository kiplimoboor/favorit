package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kiplimoboor/favorit/database"
	"github.com/kiplimoboor/favorit/models"
)

type RoomController struct {
	db database.Database
}

func NewRoomController(db database.Database) *RoomController {
	return &RoomController{db: db}
}

func (rc *RoomController) HandleCreateRoom(w http.ResponseWriter, r *http.Request) error {
	roomReq := models.CreateRoomRequest{}
	if err := json.NewDecoder(r.Body).Decode(&roomReq); err != nil {
		return err
	}
	newRoom := models.NewRoom(roomReq)
	if err := rc.db.CreateRoom(*newRoom); err != nil {
		return err
	}
	successMsg := fmt.Sprintf("room %s created", newRoom.Number)
	return WriteJSON(w, http.StatusOK, Success{Message: successMsg})
}

func (rc *RoomController) HandleGetRoom(w http.ResponseWriter, r *http.Request) error {
	number := mux.Vars(r)["number"]
	room, err := rc.db.GetRoomBy("number", number)
	if err != nil {
		return WriteJSON(w, http.StatusNotFound, Error{Error: fmt.Sprintf("room %s not found", number)})
	}
	res := models.GetRoomResponse{
		Booked:      room.Booked,
		Description: room.Description,
		Number:      room.Number,
		Size:        room.Size,
	}
	return WriteJSON(w, http.StatusOK, res)
}

func (rc *RoomController) HandleUpdateRoom(w http.ResponseWriter, r *http.Request) error {
	number := mux.Vars(r)["number"]
	updateRequest := models.UpdateRequest{}
	if err := json.NewDecoder(r.Body).Decode(&updateRequest); err != nil {
		return err
	}
	if err := rc.db.UpdateRoom(number, updateRequest.Field, updateRequest.NewValue); err != nil {
		if err == sql.ErrNoRows {
			return WriteJSON(w, http.StatusNotFound, Error{Error: fmt.Sprintf("room %s not found", number)})
		}
		return err
	}
	return WriteJSON(w, http.StatusOK, Success{Message: fmt.Sprintf("room %s updated successfully", number)})
}

func (rc *RoomController) HandleDeleteRoom(w http.ResponseWriter, r *http.Request) error {
	number := mux.Vars(r)["number"]
	if err := rc.db.DeleteRoom(number); err != nil {
		if err == sql.ErrNoRows {
			return WriteJSON(w, http.StatusNotFound, Error{Error: fmt.Sprintf("room %s not found", number)})
		}
		return err
	}
	return WriteJSON(w, http.StatusOK, Success{Message: fmt.Sprintf("room %s deleted", number)})
}
