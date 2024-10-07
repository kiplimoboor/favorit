package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Server struct {
	listenAddr string
	db         *SQLiteDB
}

func NewServer(listenAddr string, db *SQLiteDB) *Server {
	return &Server{listenAddr: listenAddr, db: db}
}

func (s *Server) Run() {
	s.db.Init()
	router := mux.NewRouter()

	router.HandleFunc("/users", makeHandlerFunc(s.handleUser))

	log.Println("Server started and is running in port", s.listenAddr)
	http.ListenAndServe(s.listenAddr, router)
}

func (s *Server) handleUser(w http.ResponseWriter, r *http.Request) error {
	switch r.Method {
	case "POST":
		return s.handleAddUser(w, r)
	default:
		return fmt.Errorf("Method not allowed")
	}
}

func (s *Server) handleAddUser(w http.ResponseWriter, r *http.Request) error {
	newUserRequest := CreateUserRequest{}
	err := json.NewDecoder(r.Body).Decode(&newUserRequest)
	if err != nil {
		return WriteJSON(w, http.StatusBadRequest, APIError{Error: err.Error()})
	}

	createdUser, err := CreateNewUser(newUserRequest)
	if err != nil {
		return WriteJSON(w, http.StatusBadRequest, APIError{Error: err.Error()})
	}

	err = s.db.AddUser(createdUser)
	if err != nil {
		return WriteJSON(w, http.StatusBadRequest, APIError{Error: err.Error()})
	}

	return WriteJSON(w, http.StatusCreated, APISuccess{Message: "user successfully created"})
}

/* func (s *Server) handleGetUser(w http.ResponseWriter, r *http.Request) error {

	return nil
} */

type APIError struct {
	Error string `json:"error"`
}

type APISuccess struct {
	Message string `json:"message"`
}

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

type apiHandlerFunc func(http.ResponseWriter, *http.Request) error

func makeHandlerFunc(f apiHandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			WriteJSON(w, http.StatusBadRequest, APIError{Error: err.Error()})
		}
	}
}
