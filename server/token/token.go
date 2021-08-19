package token

import (
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

type Claims struct {
	Login string `json:"login"`
	jwt.StandardClaims
}

func NewClaims(login string) (*Claims, time.Time) {
	expirationTime := time.Now().Add(30 * time.Minute)
	return &Claims{
			Login: login,
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: expirationTime.Unix(),
			},
		},
		expirationTime
}

func (tk *Claims) String() string {
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, err := token.SignedString([]byte(os.Getenv("token_password")))
	if err != nil {
		panic(err)
	}

	return tokenString
}
