package controllers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
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
	if err := validateBookingReq(bookingRequest); err != nil {
		return err
	}
	if err := bc.db.CreateBooking(bookingRequest); err != nil {
		return err
	}
	return WriteJSON(w, http.StatusOK, Success{Message: "booking made successfully"})
}

func (bc *BookingController) HandleGetBooking(w http.ResponseWriter, r *http.Request) error {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	b, err := bc.db.GetBookingBy("id", id)
	if err != nil {
		return err
	}
	return WriteJSON(w, http.StatusOK, b)
}

func (bc *BookingController) HandleGetAllBookings(w http.ResponseWriter, r *http.Request) error {
	bookings, err := bc.db.GetAllBookings()
	if err != nil {
		return err
	}
	return WriteJSON(w, http.StatusOK, bookings)
}

func (bc *BookingController) HandleUpdateBooking(w http.ResponseWriter, r *http.Request) error {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	updateRequest := models.UpdateRequest{}
	if err := json.NewDecoder(r.Body).Decode(&updateRequest); err != nil {
		return err
	}
	err := bc.db.UpdateBooking(id, updateRequest.Field, updateRequest.NewValue)
	if err != nil {
		if err == sql.ErrNoRows {
			return WriteJSON(w, http.StatusNotFound, Error{Error: "not found"})
		}
		return err
	}
	return WriteJSON(w, http.StatusOK, Success{Message: "updated successfully"})
}

func (bc *BookingController) HandleCheckout(w http.ResponseWriter, r *http.Request) error {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	if err := bc.db.Vacate(id, "checkout"); err != nil {
		return WriteJSON(w, http.StatusBadRequest, Error{Error: err.Error()})
	}
	return WriteJSON(w, http.StatusOK, Success{Message: "booking checked out"})
}

func (bc *BookingController) HandleCancel(w http.ResponseWriter, r *http.Request) error {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	if err := bc.db.Vacate(id, "cancelled"); err != nil {
		return WriteJSON(w, http.StatusBadRequest, Error{Error: err.Error()})
	}
	return WriteJSON(w, http.StatusOK, Success{Message: "booking cancelled"})
}

func validateBookingReq(br models.BookingRequest) error {
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
