package controllers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kiplimoboor/favorit/database"
	"github.com/kiplimoboor/favorit/models"
)

type UserController struct {
	db database.Database
}

func NewUserController(db database.Database) *UserController {
	return &UserController{db: db}
}

func (ctrl *UserController) HandleCreateUser(w http.ResponseWriter, r *http.Request) error {
	userReq := models.UserRequest{}
	if err := json.NewDecoder(r.Body).Decode(&userReq); err != nil {
		return err
	}
	if err := validateUserReq(userReq); err != nil {
		return err
	}
	if err := ctrl.db.CreateUser(userReq); err != nil {
		return err
	}
	successMsg := fmt.Sprintf("user %s successfully created", userReq.UserName)
	return WriteJSON(w, http.StatusCreated, Success{Message: successMsg})
}

// Only gets a user based on username
func (ctrl *UserController) HandleGetUser(w http.ResponseWriter, r *http.Request) error {
	username := mux.Vars(r)["username"]
	user, err := ctrl.db.GetUserBy("username", username)
	if err != nil {
		return WriteJSON(w, http.StatusNotFound, Error{Error: err.Error()})
	}
	return WriteJSON(w, http.StatusOK, user)
}

func (ctrl *UserController) HandleGetAllUsers(w http.ResponseWriter, r *http.Request) error {
	users, err := ctrl.db.GetAllUsers()
	if err != nil {
		return err
	}
	return WriteJSON(w, http.StatusOK, users)
}

func (ctrl *UserController) HandleUpdateUser(w http.ResponseWriter, r *http.Request) error {
	username := mux.Vars(r)["username"]
	updateRequest := models.UpdateRequest{}
	if err := json.NewDecoder(r.Body).Decode(&updateRequest); err != nil {
		return err
	}
	err := ctrl.db.UpdateUser(username, updateRequest.Field, updateRequest.NewValue)
	if err != nil {
		if err == sql.ErrNoRows {
			return WriteJSON(w, http.StatusNotFound, Error{Error: fmt.Sprintf("user %s not found", username)})
		}
		return err
	}
	return WriteJSON(w, http.StatusOK, Success{Message: fmt.Sprintf("user %s updated successfully", username)})
}

func (ctrl *UserController) HandleDeleteUser(w http.ResponseWriter, r *http.Request) error {
	username := mux.Vars(r)["username"]
	err := ctrl.db.DeleteUser(username)
	if err != nil {
		if err == sql.ErrNoRows {
			return WriteJSON(w, http.StatusNotFound, Error{Error: fmt.Sprintf("user %s not found", username)})
		}
		return err
	}
	return WriteJSON(w, http.StatusOK, Success{Message: fmt.Sprintf("user %s deleted", username)})
}

func validateUserReq(ur models.UserRequest) error {
	if ur.FirstName == "" {
		return errors.New("first name is required")
	}
	if ur.LastName == "" {
		return errors.New("last name is required")
	}
	if ur.Email == "" {
		return errors.New("email is required")
	}
	if ur.Password == "" {
		return errors.New("password is required")
	}
	return nil
}
