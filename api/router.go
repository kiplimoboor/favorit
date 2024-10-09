package api

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kiplimoboor/favorit/api/controllers"
	"github.com/kiplimoboor/favorit/database"
)

func NewRouter(db *database.SQLiteDB) *mux.Router {

	userCtrl := controllers.NewUserController(db)

	router := mux.NewRouter()

	router.HandleFunc("/users", makeHandlerFunc(userCtrl.HandleCreateUser)).Methods("POST")
	router.HandleFunc("/users", makeHandlerFunc(userCtrl.HandleUpdateUser)).Methods("PATCH")
	router.HandleFunc("/users/{username}", makeHandlerFunc(userCtrl.HandleGetUser)).Methods("GET")
	router.HandleFunc("/users/", makeHandlerFunc(userCtrl.HandleDeleteUser)).Methods("DELETE")

	return router
}

type apiHandlerFunc func(http.ResponseWriter, *http.Request) error

func makeHandlerFunc(f apiHandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := f(w, r)
		if err != nil {
		}
	}
}
