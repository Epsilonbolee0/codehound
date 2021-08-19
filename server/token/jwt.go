package token

import (
	"net/http"
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

var JwtAuthenithication = func(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		notAuth := []string{"/auth/register", "/auth/login"}
		requestPath := r.URL.Path
		for _, value := range notAuth {
			if value == requestPath {
				next.ServeHTTP(w, r)
				return
			}
		}

		c, err := r.Cookie("token")
		if err != nil {
			if err == http.ErrNoCookie {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			w.WriteHeader(http.StatusBadRequest)
			return
		}

		tokenString := c.Value
		claims := &Claims{}

		token, err := jwt.ParseWithClaims(tokenString, claims, func(tk *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("token_password")), nil
		})
		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if !token.Valid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		if time.Now().Unix()-claims.ExpiresAt > 30 {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		claims, expTime := NewClaims(claims.Login)
		tokenStr := claims.String()

		http.SetCookie(w, &http.Cookie{
			Name:    "token",
			Value:   tokenStr,
			Expires: expTime,
			Path:    "/",

			HttpOnly: true,
		})

		next.ServeHTTP(w, r)
	})
}
