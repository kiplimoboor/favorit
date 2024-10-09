package api

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kiplimoboor/favorit/api/controllers"
	"github.com/kiplimoboor/favorit/database"
)

func NewRouter(db *database.SQLiteDB) *mux.Router {

	userRepo := database.NewUserRepository(db)
	userController := controllers.NewUserController(userRepo)

	router := mux.NewRouter()

	router.HandleFunc("/users", makeHandlerFunc(userController.HandleCreateUser)).Methods("POST")

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
