package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"

	"../token"

	"../models/domain"
	"../models/service"
	"../utils"
)

type AuthController struct {
	authService *service.AuthService
}

func SetupAuthController(authService *service.AuthService, router *mux.Router) {
	controller := &AuthController{authService: authService}
	router.HandleFunc("/auth/login", controller.Login).Methods("POST")
	router.HandleFunc("/auth/logout", controller.Logout).Methods("POST")
	router.HandleFunc("/auth/register", controller.Register).Methods("POST")
	router.HandleFunc("/auth/home", controller.Home).Methods("GET")
}

func (controller *AuthController) Login(w http.ResponseWriter, r *http.Request) {
	account := &domain.Account{}
	err := json.NewDecoder(r.Body).Decode(account)
	if err != nil {
		utils.Respond(w, utils.Message(http.StatusBadRequest, "Invalid request"))
	} else {
		resp := controller.authService.Login(account.Login, account.Password)
		claims, expTime := token.NewClaims(account.Login)
		token := claims.String()

		http.SetCookie(w, &http.Cookie{
			Name:    "token",
			Value:   token,
			Expires: expTime,
			Path:    "/",

			HttpOnly: true,
		})

		utils.Respond(w, resp)
	}
}

func (controller *AuthController) Logout(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   "",
		Expires: time.Unix(0, 0),
		Path:    "/",

		HttpOnly: true,
	})

	resp := utils.Message(http.StatusOK, "Logout was successful")
	utils.Respond(w, resp)
}

func (controller *AuthController) Register(w http.ResponseWriter, r *http.Request) {
	account := &domain.Account{}
	err := json.NewDecoder(r.Body).Decode(account)
	if err != nil {
		utils.Respond(w, utils.Message(http.StatusBadRequest, "Invalid request"))
	} else {
		resp := controller.authService.Register(account.Login, account.Email, account.Password)
		utils.Respond(w, resp)
	}
}

func (controller *AuthController) Home(w http.ResponseWriter, r *http.Request) {
	c, _ := r.Cookie("token")
	claims := &token.Claims{}
	tokenString := c.Value

	jwt.ParseWithClaims(tokenString, claims, func(tk *jwt.Token) (interface{}, error) {
		return os.Getenv("token_password"), nil
	})

	w.Write([]byte(fmt.Sprintf("Welcome, %s!", claims.Login)))
}
