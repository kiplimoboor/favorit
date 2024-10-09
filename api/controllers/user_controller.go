package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kiplimoboor/favorit/database"
	"github.com/kiplimoboor/favorit/database/models"
)

type UserController struct {
	db *database.SQLiteDB
}

func NewUserController(db *database.SQLiteDB) *UserController {
	return &UserController{db: db}
}

func (ctrl *UserController) HandleCreateUser(w http.ResponseWriter, r *http.Request) error {
	newUserReq := models.CreateUserRequest{}
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

	if err = ctrl.db.CreateUser(*newUser); err != nil {
		return WriteJSON(w, http.StatusBadRequest, Error{Error: err.Error()})
	}
	successMsg := fmt.Sprintf("user %s successfully created", newUser.UserName)
	return WriteJSON(w, http.StatusCreated, Success{Message: successMsg})
}

// Only gets a user based on username
func (ctrl *UserController) HandleGetUser(w http.ResponseWriter, r *http.Request) error {
	username := mux.Vars(r)["username"]
	user, err := ctrl.db.GetUserBy("username", username)
	if err != nil {
		return WriteJSON(w, http.StatusNotFound, Error{Error: "user not found"})
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
	updateRequest := models.UpdateUserRequest{}
	json.NewDecoder(r.Body).Decode(&updateRequest)
	err := ctrl.db.UpdateUser(updateRequest.Email, updateRequest.Field, updateRequest.NewValue)
	if err != nil {
		return WriteJSON(w, http.StatusBadRequest, Error{Error: err.Error()})
	}
	successMsg := fmt.Sprintf("user %s updated successfully", updateRequest.Email)
	return WriteJSON(w, http.StatusOK, Success{Message: successMsg})
}

func (ctrl *UserController) HandleDeleteUser(w http.ResponseWriter, r *http.Request) error {
	email := r.URL.Query().Get("email")
	err := ctrl.db.DeleteUser(email)
	if err != nil {
		return WriteJSON(w, http.StatusBadRequest, Error{Error: err.Error()})
	}
	return WriteJSON(w, http.StatusOK, Success{Message: fmt.Sprintf("user %s deleted", email)})
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
