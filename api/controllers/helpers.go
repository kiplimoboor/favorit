package controllers

import (
	"encoding/json"
	"net/http"
)

type Error struct {
	Error string `json:"error"`
}

type Success struct {
	Message string `json:"message"`
}

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}
