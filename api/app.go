package api

import (
	"log"
	"net/http"

	"github.com/kiplimoboor/favorit/database"
)

type Server struct {
	listenAddr string
	db         database.Database
}

func NewServer(listenAddr string, db database.Database) *Server {
	return &Server{listenAddr: listenAddr, db: db}
}

func (s *Server) Start() {

	router := NewRouter(s.db)

	log.Println("Server started in port", s.listenAddr)
	http.ListenAndServe(s.listenAddr, router)
}
