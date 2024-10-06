package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Server struct {
	listenAddr string
}

func NewServer(listenAddr string) *Server {
	return &Server{
		listenAddr: listenAddr,
	}
}

func (s *Server) Run() {
	router := mux.NewRouter()

	http.ListenAndServe(s.listenAddr, router)
}
