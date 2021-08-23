package app

import (
	"fmt"
	"net/http"
	"os"
	"time"

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
	setupToolsController(connections.GetConnection("admin"), router)
	setupVersioningController(connections.GetConnection("admin"), router)
	setupFunctionsController(connections.GetConnection("admin"), router)

	host := os.Getenv("server_host")
	port := os.Getenv("server_port")
	now := time.Now()

	fmt.Printf("[%s] Listening to %s:%s\n", now.Format("15:04:05"), host, port)

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

func setupToolsController(conn *gorm.DB, router *mux.Router) {
	libraryRepo := repository.NewLibraryRepository(conn)
	languageRepo := repository.NewLanguageRepository(conn)
	toolsService := service.NewToolsService(languageRepo, libraryRepo)
	controllers.SetupToolsController(toolsService, router)
}

func setupVersioningController(conn *gorm.DB, router *mux.Router) {
	accountRepo := repository.NewAccountRepository(conn)
	versionRepo := repository.NewVersionRepository(conn)
	libraryRepo := repository.NewLibraryRepository(conn)
	implRepo := repository.NewImplementationRepository(conn)
	treeRepo := repository.NewTreeRepository(conn)
	tagRepo := repository.NewTagRepository(conn)
	versioningService := service.NewVersioningService(accountRepo, versionRepo, implRepo, libraryRepo, treeRepo, tagRepo)
	controllers.SetupVersioningController(versioningService, router)
}

func setupFunctionsController(conn *gorm.DB, router *mux.Router) {
	languageRepo := repository.NewLanguageRepository(conn)
	implRepo := repository.NewImplementationRepository(conn)
	functionService := service.NewFunctionService(implRepo, languageRepo)
	controllers.SetupFunctionController(functionService, router)
}
