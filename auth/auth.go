package auth

import (
	"errors"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var SECRET_KEY = []byte("my secret")

type Claims struct {
	Email string `json:"email"`
	Role  string `json:"role"`
	jwt.RegisteredClaims
}

func AuthAdmin(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		claims, err := CheckJWT(r)
		if err != nil || claims.Role != "admin" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	}
}

func CheckJWT(r *http.Request) (*Claims, error) {
	cookie, err := r.Cookie("token")
	if err != nil {
		return nil, err
	}
	tokenString := cookie.Value
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (any, error) {
		if token.Method != jwt.SigningMethodHS256 {
			return nil, errors.New("invalid signing method")
		}
		return SECRET_KEY, nil
	})
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, errors.New("invalid token")
	}
	claims, ok := token.Claims.(*Claims)
	if !ok {
		return nil, errors.New("unkown claims")
	}
	return claims, nil
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
