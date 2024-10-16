package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kiplimoboor/favorit/database"
	"github.com/kiplimoboor/favorit/models"
)

type GuestController struct {
	db database.Database
}

func NewGuestController(db database.Database) *GuestController {
	return &GuestController{db: db}
}

func (gc *GuestController) HandleCreateGuest(w http.ResponseWriter, r *http.Request) error {
	gr := models.GuestRequest{}
	if err := json.NewDecoder(r.Body).Decode(&gr); err != nil {
		return err
	}
	if err := gc.db.CreateGuest(gr); err != nil {
		return err
	}
	return WriteJSON(w, http.StatusOK, Success{Message: "guest created"})
}

func (gc *GuestController) HandleGetGuest(w http.ResponseWriter, r *http.Request) error {
	email := mux.Vars(r)["email"]

	guest, err := gc.db.GetGuestBy("email", email)
	if err != nil {
		return WriteJSON(w, http.StatusNotFound, Error{Error: "guest not found"})
	}
	return WriteJSON(w, http.StatusOK, guest)
}
