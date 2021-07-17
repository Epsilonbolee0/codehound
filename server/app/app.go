package app

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"gorm.io/gorm"

	"../connections"
	"../controllers"
	"../models/repository"
	"../models/service"
	"../token"
)

func Run() {
	router := mux.NewRouter()
	router.Use(token.JwtAuthenithication)
	setupAuthController(connections.GetConnection("admin"), router)

	host := os.Getenv("server_host")
	port := os.Getenv("server_port")

	fmt.Printf(" Listening to %s:%s\n", host, port)

	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		panic(err)
	}
}

func setupAuthController(conn *gorm.DB, router *mux.Router) {
	authRepo := repository.NewAccountRepository(conn)
	authService := service.NewAuthService(authRepo)
	controllers.SetupAuthController(authService, router)
}
