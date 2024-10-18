package api

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kiplimoboor/favorit/api/controllers"
	"github.com/kiplimoboor/favorit/auth"
	"github.com/kiplimoboor/favorit/database"
)

func NewRouter(db database.Database) *mux.Router {
	router := mux.NewRouter()

	// Login Logout
	lc := controllers.NewLoginController(db)
	router.HandleFunc("/login", makeHandler(lc.HandleLogin)).Methods(http.MethodPost)
	router.HandleFunc("/logout", makeHandler(controllers.HandleLogout)).Methods(http.MethodPost)

	// Users
	uc := controllers.NewUserController(db)
	router.HandleFunc("/users", auth.AuthAdmin(makeHandler(uc.HandleCreateUser))).Methods(http.MethodPost)
	router.HandleFunc("/users", auth.AuthAdmin(makeHandler(uc.HandleGetAllUsers))).Methods(http.MethodGet)
	router.HandleFunc("/users/{username}", auth.AuthAdmin(makeHandler(uc.HandleUpdateUser))).Methods(http.MethodPatch)
	router.HandleFunc("/users/{username}", makeHandler(uc.HandleGetUser)).Methods(http.MethodGet)
	router.HandleFunc("/users/{username}", auth.AuthAdmin(makeHandler(uc.HandleDeleteUser))).Methods(http.MethodDelete)

	// Rooms
	rc := controllers.NewRoomController(db)
	router.HandleFunc("/rooms", auth.AuthAdmin(makeHandler(rc.HandleCreateRoom))).Methods(http.MethodPost)
	router.HandleFunc("/rooms", auth.AuthAdmin(makeHandler(rc.HandleGetAllRooms))).Methods(http.MethodGet)
	router.HandleFunc("/rooms/{number}", auth.AuthAdmin(makeHandler(rc.HandleUpdateRoom))).Methods(http.MethodPatch)
	router.HandleFunc("/rooms/{number}", makeHandler(rc.HandleGetRoom)).Methods(http.MethodGet)
	router.HandleFunc("/rooms/{number}", auth.AuthAdmin(makeHandler(rc.HandleDeleteRoom))).Methods(http.MethodDelete)

	// Guests
	gc := controllers.NewGuestController(db)
	router.HandleFunc("/guests", makeHandler(gc.HandleCreateGuest)).Methods(http.MethodPost)
	router.HandleFunc("/guests", makeHandler(gc.HandleGetAllGuests)).Methods(http.MethodGet)
	router.HandleFunc("/guests/{email}", makeHandler(gc.HandleGetGuest)).Methods(http.MethodGet)

	// Bookings
	bc := controllers.NewBookingController(db)
	router.HandleFunc("/bookings", makeHandler(bc.HandleCreateBooking)).Methods(http.MethodPost)
	router.HandleFunc("/bookings", makeHandler(bc.HandleGetAllBookings)).Methods(http.MethodGet)
	router.HandleFunc("/bookings/{id}", makeHandler(bc.HandleGetBooking)).Methods(http.MethodGet)
	router.HandleFunc("/bookings/{id}", makeHandler(bc.HandleUpdateBooking)).Methods(http.MethodPatch)
	router.HandleFunc("/bookings/{id}", makeHandler(bc.HandleCheckout)).Methods(http.MethodPost)
	router.HandleFunc("/bookings/cancel/{id}", makeHandler(bc.HandleCancel)).Methods(http.MethodPost)

	return router
}

type apiHandlerFunc func(http.ResponseWriter, *http.Request) error

func makeHandler(f apiHandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := f(w, r)
		if err != nil {
			controllers.WriteJSON(w, http.StatusBadRequest, controllers.Error{Error: err.Error()})
		}
	}
}
