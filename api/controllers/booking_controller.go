package controllers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/kiplimoboor/favorit/database"
	"github.com/kiplimoboor/favorit/models"
)

type BookingController struct {
	db database.Database
}

func NewBookingController(db database.Database) *BookingController {
	return &BookingController{db: db}
}

func (bc *BookingController) HandleCreateBooking(w http.ResponseWriter, r *http.Request) error {
	bookingRequest := models.BookingRequest{}
	if err := json.NewDecoder(r.Body).Decode(&bookingRequest); err != nil {
		return err
	}
	if err := validateBooking(bookingRequest); err != nil {
		return err
	}
	if err := bc.db.CreateBooking(bookingRequest); err != nil {
		return err
	}
	return WriteJSON(w, http.StatusOK, Success{Message: "booking made successfully"})
}

func validateBooking(br models.BookingRequest) error {
	if br.GuestEmail == "" {
		return errors.New("guest email required")
	}
	if br.RoomNumber == "" {
		return errors.New("room number is required")
	}
	if br.CheckIn < 0 {
		return errors.New("invalid check in date")
	}
	if br.CheckOut < 0 {
		return errors.New("invalid check out date")
	}
	return nil
}
