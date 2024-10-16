package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/kiplimoboor/favorit/auth"
	"github.com/kiplimoboor/favorit/database"
)

type LoginController struct {
	db database.Database
}

type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func NewLoginController(db database.Database) *LoginController {
	return &LoginController{db: db}
}

func (lc *LoginController) HandleLogin(w http.ResponseWriter, r *http.Request) error {
	credentials := Credentials{}
	if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
		return err
	}
	user, err := lc.db.GetUserBy("email", credentials.Email)
	if err != nil || !auth.CheckPasswordHash(credentials.Password, user.Password) {
		return WriteJSON(w, http.StatusUnauthorized, Error{Error: "unauthorized"})
	}
	claims := auth.Claims{Email: user.Email, Role: user.Role}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString(auth.SECRET_KEY)
	if err != nil {
		fmt.Println(err)
		return WriteJSON(w, http.StatusInternalServerError, Error{Error: "internal server error"})
	}

	http.SetCookie(w, &http.Cookie{Name: "token", Value: tokenStr})
	return WriteJSON(w, http.StatusOK, Success{Message: "login successful"})
}
