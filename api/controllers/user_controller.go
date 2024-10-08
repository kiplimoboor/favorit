package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/kiplimoboor/favorit/api/database/models"
)

func (c *Controller) HandleCreateUser(w http.ResponseWriter, r *http.Request) error {
	newUserReq := models.UserRequest{}
	if err := json.NewDecoder(r.Body).Decode(&newUserReq); err != nil {
		return WriteJSON(w, http.StatusBadRequest, Error{Error: err.Error()})
	}

	newUser, err := models.CreateNewUser(newUserReq)
	if err != nil {
		return WriteJSON(w, http.StatusBadRequest, Error{Error: err.Error()})
	}
	if err = validateUser(newUser); err != nil {
		return WriteJSON(w, http.StatusBadRequest, Error{Error: err.Error()})
	}
	if err = c.db.CreateUser(*newUser); err != nil {
		return WriteJSON(w, http.StatusBadRequest, Error{Error: err.Error()})
	}
	successMsg := fmt.Sprintf("user %s successfully created", newUser.UserName)
	return WriteJSON(w, http.StatusOK, Success{Message: successMsg})
}

func validateUser(user *models.User) error {
	if user.FirstName == "" {
		return errors.New("first name is required")
	}
	if user.LastName == "" {
		return errors.New("last name is required")
	}
	if user.Email == "" {
		return errors.New("email is required")
	}
	if user.Password == "" {
		return errors.New("password is required")
	}
	return nil
}
