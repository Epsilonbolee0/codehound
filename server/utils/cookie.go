package utils

import (
	"net/http"
	"os"

	"../token"
	"github.com/dgrijalva/jwt-go"
)

func LoginFromCookie(r *http.Request) (string, error) {
	c, err := r.Cookie("token")
	if err != nil {
		return "", err
	}

	claims := &token.Claims{}
	tokenString := c.Value

	jwt.ParseWithClaims(tokenString, claims, func(tk *jwt.Token) (interface{}, error) {
		return os.Getenv("token_password"), nil
	})

	return claims.Login, nil
}
