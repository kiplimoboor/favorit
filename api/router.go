package api

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kiplimoboor/favorit/api/controllers"
	"github.com/kiplimoboor/favorit/database"
)

func NewRouter(db database.Database) *mux.Router {

	uc := controllers.NewUserController(db)
	rc := controllers.NewRoomController(db)

	router := mux.NewRouter()

	// Users
	router.HandleFunc("/users", makeHandlerFunc(uc.HandleCreateUser)).Methods(http.MethodPost)
	router.HandleFunc("/users/{username}", makeHandlerFunc(uc.HandleUpdateUser)).Methods(http.MethodPatch)
	router.HandleFunc("/users/{username}", makeHandlerFunc(uc.HandleGetUser)).Methods(http.MethodGet)
	router.HandleFunc("/users/{username}", makeHandlerFunc(uc.HandleDeleteUser)).Methods(http.MethodDelete)

	// Rooms
	router.HandleFunc("/rooms", makeHandlerFunc(rc.HandleCreateRoom)).Methods(http.MethodPost)
	router.HandleFunc("/rooms/{number}", makeHandlerFunc(rc.HandleUpdateRoom)).Methods(http.MethodPatch)
	router.HandleFunc("/rooms/{number}", makeHandlerFunc(rc.HandleGetRoom)).Methods(http.MethodGet)
	router.HandleFunc("/rooms/{number}", makeHandlerFunc(rc.HandleDeleteRoom)).Methods(http.MethodDelete)

	return router
}

type apiHandlerFunc func(http.ResponseWriter, *http.Request) error

func makeHandlerFunc(f apiHandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := f(w, r)
		if err != nil {
			controllers.WriteJSON(w, http.StatusBadRequest, controllers.Error{Error: err.Error()})
		}
	}
}
