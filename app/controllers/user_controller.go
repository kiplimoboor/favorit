package app

func SomeFunction() {
	// This will not work if apiStruct is unexported
	// var a apiStruct  // This will cause an error

	// This will work
	var response ApiResponse
	response.Message = "Success"
}

// import (
// 	"fmt"
// 	"net/http"
// )
//
// func (s *Server) HandleUser(w http.ResponseWriter, r *http.Request) error {
// 	switch r.Method {
// 	case http.MethodPost:
// 		return s.handleAddUser(w, r)
// 	case http.MethodGet:
// 		return s.handleGetUser(w, r)
// 	case http.MethodPatch:
// 		return s.handleUpdateUser(w, r)
// 	case http.MethodDelete:
// 		return s.handleDeleteUser(w, r)
// 	default:
// 		return fmt.Errorf("Method not allowed")
// 	}
// }
//
// func (s *Server) handleAddUser(w http.ResponseWriter, r *http.Request) error {
	newUserRequest := models.UserRequest{}
	err := json.NewDecoder(r.Body).Decode(&newUserRequest)
	if err != nil {
		return WriteJSON(w, http.StatusBadRequest, APIError{Error: err.Error()})
	}

	createdUser, err := models.CreateNewUser(newUserRequest)
	if err != nil {
		return WriteJSON(w, http.StatusBadRequest, APIError{Error: err.Error()})
	}

	err = s.db.AddUser(*createdUser)
	if err != nil {
		var errMsg string
		if strings.Contains(err.Error(), "email") {
			errMsg = fmt.Sprintf("user with email %s already exists", createdUser.Email)
		}
		if strings.Contains(err.Error(), "username") {
			errMsg = fmt.Sprintf("username %s is taken", createdUser.UserName)
		}
		return WriteJSON(w, http.StatusBadRequest, APIError{Error: errMsg})
	}

	return WriteJSON(w, http.StatusCreated, APISuccess{Message: "user successfully created"})
// }
//
// // Gets user based on username variable
// func (s *Server) handleGetUser(w http.ResponseWriter, r *http.Request) error {
// 	username := mux.Vars(r)["username"]
// 	user, err := s.db.GetUserBy("username", username)
// 	if err != nil {
// 		return WriteJSON(w, http.StatusNotFound, APIError{Error: "user not found"})
// 	}
// 	userRes := models.GetUserResponse{
// 		FirstName: user.FirstName,
// 		LastName:  user.LastName,
// 		Email:     user.Email,
// 		UserName:  user.UserName,
// 		CreatedAt: user.CreatedAt,
// 	}
// 	return WriteJSON(w, http.StatusOK, userRes)
// }
//
// func (s *Server) handleUpdateUser(w http.ResponseWriter, r *http.Request) error {
// 	updateRequest := models.UpdateUserRequest{}
// 	json.NewDecoder(r.Body).Decode(&updateRequest)
// 	err := s.db.UpdateUser(updateRequest.Email, updateRequest.Field, updateRequest.NewValue)
// 	if err != nil {
// 		return WriteJSON(w, http.StatusBadRequest, APIError{Error: err.Error()})
// 	}
// 	successMsg := fmt.Sprintf("user %s updated successfully", updateRequest.Email)
// 	return WriteJSON(w, http.StatusOK, APISuccess{Message: successMsg})
// }
//
// func (s *Server) handleDeleteUser(w http.ResponseWriter, r *http.Request) error {
// 	email := r.URL.Query().Get("email")
// 	err := s.db.DeleteUser(email)
// 	if err != nil {
// 		return WriteJSON(w, http.StatusBadRequest, APIError{Error: err.Error()})
// 	}
// 	return WriteJSON(w, http.StatusOK, APISuccess{Message: fmt.Sprintf("user %s deleted", email)})
// }
