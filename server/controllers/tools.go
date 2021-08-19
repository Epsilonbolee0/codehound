package controllers

import (
	"encoding/json"
	"net/http"

	"../models/domain"
	"../models/service"
	"github.com/gorilla/mux"

	"../utils"
)

type ToolsController struct {
	toolsService *service.ToolsService
}

func SetupToolsController(toolsService *service.ToolsService, router *mux.Router) {
	controller := &ToolsController{toolsService: toolsService}
	router.HandleFunc("/tools/lang/add", controller.AddLanguage).Methods("POST")
	router.HandleFunc("/tools/lang/list", controller.ListLanguages).Methods("GET")
	router.HandleFunc("/tools/lang/versions", controller.ListLanguageVersions).Methods("GET")
	router.HandleFunc("/tools/lang/delete", controller.DeleteLanguageVersion).Methods("DELETE")

	router.HandleFunc("/tools/lib/add", controller.AddLibrary).Methods("POST")
	router.HandleFunc("/tools/lib/list", controller.ListLibraries).Methods("GET")
	router.HandleFunc("/tools/lib/versions", controller.ListLibraryVersions).Methods("GET")
	router.HandleFunc("/tools/lib/delete", controller.DeleteLibraryVersion).Methods("DELETE")
}

func (controller *ToolsController) AddLanguage(w http.ResponseWriter, r *http.Request) {
	var resp map[string]interface{}
	language := &domain.Language{}

	err := json.NewDecoder(r.Body).Decode(language)

	if err != nil {
		resp = utils.Message(http.StatusBadRequest, "Invalid request")
	} else {
		resp = controller.toolsService.AddLanguage(language.Name, language.Version)
	}

	utils.Respond(w, resp)
}

func (controller *ToolsController) ListLanguages(w http.ResponseWriter, r *http.Request) {
	utils.Respond(w, controller.toolsService.ListLanguages())
}

func (controller *ToolsController) ListLanguageVersions(w http.ResponseWriter, r *http.Request) {
	var resp map[string]interface{}
	language := &domain.Language{}
	err := json.NewDecoder(r.Body).Decode(language)

	if err != nil {
		resp = utils.Message(http.StatusBadRequest, "Invalid request")
	} else {
		resp = controller.toolsService.ListLanguageVersions(language.Name)
	}

	utils.Respond(w, resp)
}

func (controller *ToolsController) DeleteLanguageVersion(w http.ResponseWriter, r *http.Request) {
	var resp map[string]interface{}
	language := &domain.Language{}
	err := json.NewDecoder(r.Body).Decode(language)

	if err != nil {
		resp = utils.Message(http.StatusBadRequest, "Invalid request")
	} else {
		resp = controller.toolsService.DeleteLanguageVersion(language.Name, language.Version)
	}

	utils.Respond(w, resp)
}

func (controller *ToolsController) AddLibrary(w http.ResponseWriter, r *http.Request) {
	var resp map[string]interface{}
	library := &domain.LibraryDTO{}
	err := json.NewDecoder(r.Body).Decode(library)

	if err != nil {
		resp = utils.Message(http.StatusBadRequest, "Invalid request")
	} else {
		resp = controller.toolsService.AddLibrary(library.Name, library.Version, library.LanguageName, library.LanguageVersion)
	}

	utils.Respond(w, resp)
}

func (controller *ToolsController) ListLibraries(w http.ResponseWriter, r *http.Request) {
	var resp map[string]interface{}
	library := &domain.Library{}
	err := json.NewDecoder(r.Body).Decode(library)

	if err != nil {
		resp = utils.Message(http.StatusBadRequest, "Invalid request")
	} else {
		resp = controller.toolsService.ListLibraries(library.Name, library.Version)
	}

	utils.Respond(w, resp)
}

func (controller *ToolsController) ListLibraryVersions(w http.ResponseWriter, r *http.Request) {
	var resp map[string]interface{}
	library := &domain.Library{}
	err := json.NewDecoder(r.Body).Decode(library)

	if err != nil {
		resp = utils.Message(http.StatusBadRequest, "Invalid request")
	} else {
		resp = controller.toolsService.ListLibraryVersions(library.Name)
	}

	utils.Respond(w, resp)
}

func (controller *ToolsController) DeleteLibraryVersion(w http.ResponseWriter, r *http.Request) {
	var resp map[string]interface{}
	library := &domain.Library{}
	err := json.NewDecoder(r.Body).Decode(library)

	if err != nil {
		resp = utils.Message(http.StatusBadRequest, "Invalid request")
	} else {
		resp = controller.toolsService.DeleteLibraryVersion(library.Name, library.Version)
	}

	utils.Respond(w, resp)
}
