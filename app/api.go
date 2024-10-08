package app

// unexported struct (not accessible outside this file)
type apiStruct struct {
	// fields
}

// exported struct (accessible)
type ApiResponse struct {
	Message string `json:"message"`
}

/* import (
	"log"
	"net/http"
	"strings"

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

	router.HandleFunc("/users/", makeHandlerFunc(s.handleUser))
	router.HandleFunc("/users/{username}", makeHandlerFunc(s.handleUser))

	log.Println("Server started and is running in port", s.listenAddr)
	http.ListenAndServe(s.listenAddr, router)
}

type APIError struct {
	Error string `json:"error"`
}

type APISuccess struct {
	Message string `json:"message"`
}

type apiHandlerFunc func(http.ResponseWriter, *http.Request) error

func makeHandlerFunc(f apiHandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		url := *r.URL
		url.Path = strings.TrimSuffix(r.URL.Path, "/")
		r.URL = &url
		err := f(w, r)
		if err != nil {
			WriteJSON(w, http.StatusBadRequest, APIError{Error: err.Error()})
		}
	}
} */
