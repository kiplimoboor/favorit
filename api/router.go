package api

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kiplimoboor/favorit/api/controllers"
)

func NewRouter(c controllers.Controller) *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/users", makeHandlerFunc(c.HandleCreateUser)).Methods("POST")
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
