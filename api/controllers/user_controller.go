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
	newUserReq := models.CreateUserRequest{}
	if err := json.NewDecoder(r.Body).Decode(&newUserReq); err != nil {
		return err
	}
	newUser := models.CreateNewUser(newUserReq)
	if err := validateUser(newUser); err != nil {
		return err
	}
	if err := ctrl.db.CreateUser(*newUser); err != nil {
		return err
	}
	successMsg := fmt.Sprintf("user %s successfully created", newUser.UserName)
	return WriteJSON(w, http.StatusCreated, Success{Message: successMsg})
}

// Only gets a user based on username
func (ctrl *UserController) HandleGetUser(w http.ResponseWriter, r *http.Request) error {
	username := mux.Vars(r)["username"]
	user, err := ctrl.db.GetUserBy("username", username)
	if err != nil {
		return WriteJSON(w, http.StatusNotFound, Error{Error: fmt.Sprintf("user %s not found", username)})
	}
	userRes := models.GetUserResponse{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		UserName:  user.UserName,
		CreatedAt: user.CreatedAt,
	}
	return WriteJSON(w, http.StatusOK, userRes)
}

func (ctrl *UserController) HandleUpdateUser(w http.ResponseWriter, r *http.Request) error {

	username := mux.Vars(r)["username"]
	updateRequest := models.UpdateUserRequest{}
	json.NewDecoder(r.Body).Decode(&updateRequest)
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
