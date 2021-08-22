package controllers

import (
	"encoding/json"
	"net/http"

	"../models/domain"
	"../models/service"
	"../utils"
	"github.com/gorilla/mux"
)

type VersioningController struct {
	versioningService *service.VersioningService
}

func SetupVersioningController(versioningService *service.VersioningService, router *mux.Router) {
	controller := &VersioningController{versioningService: versioningService}
	router.HandleFunc("/versioning/add", controller.AddVersion).Methods("POST")
	router.HandleFunc("/versioning/list", controller.ListByAuthor).Methods("GET")
	router.HandleFunc("/versioning/update_code", controller.UpdateCode).Methods("PATCH")
	router.HandleFunc("/versioning/delete", controller.Delete).Methods("DELETE")

	router.HandleFunc("/versioning/list_libraries", controller.ListLibraries).Methods("GET")
	router.HandleFunc("/versioning/add_library", controller.AddLibrary).Methods("PATCH")
	router.HandleFunc("/versioning/delete_library", controller.DeleteLibrary).Methods("DELETE")

	router.HandleFunc("/versioning/tree/bfs", controller.ListTreeBFS).Methods("GET")
	router.HandleFunc("/versioning/tree/parent", controller.FindParent).Methods("GET")
	router.HandleFunc("/versioning/tree/children", controller.ListChildren).Methods("GET")
	router.HandleFunc("/versioning/tree/add", controller.AddChildVersion).Methods("POST")

}

func (controller *VersioningController) ListByAuthor(w http.ResponseWriter, r *http.Request) {
	var resp map[string]interface{}
	dto := &domain.VersionDTO{}

	err := json.NewDecoder(r.Body).Decode(dto)
	if err != nil {
		resp = utils.Message(http.StatusBadRequest, "Invalid request")
	} else {
		resp = controller.versioningService.ListByAuthor(dto.Login)
	}

	utils.Respond(w, resp)
}

func (controller *VersioningController) AddVersion(w http.ResponseWriter, r *http.Request) {
	var resp map[string]interface{}
	dto := &domain.VersionDTO{}

	err := json.NewDecoder(r.Body).Decode(dto)
	if err != nil {
		resp = utils.Message(http.StatusBadRequest, "Invalid request")
	} else {
		resp = controller.versioningService.AddVersion(dto.Code, dto.Login, dto.Implementation)
	}

	utils.Respond(w, resp)
}

func (controller *VersioningController) AddChildVersion(w http.ResponseWriter, r *http.Request) {
	var resp map[string]interface{}
	dto := &domain.VersionDTO{}

	err := json.NewDecoder(r.Body).Decode(dto)
	if err != nil {
		resp = utils.Message(http.StatusBadRequest, "Invalid request")
	} else {
		resp = controller.versioningService.AddChildVersion(dto.Login, dto.Name, dto.Link)
	}

	utils.Respond(w, resp)
}

func (controller *VersioningController) FindParent(w http.ResponseWriter, r *http.Request) {
	var resp map[string]interface{}
	dto := &domain.VersionDTO{}

	err := json.NewDecoder(r.Body).Decode(dto)
	if err != nil {
		resp = utils.Message(http.StatusBadRequest, "Invalid request")
	} else {
		resp = controller.versioningService.FindParent(dto.Name)
	}

	utils.Respond(w, resp)
}

func (controller *VersioningController) ListChildren(w http.ResponseWriter, r *http.Request) {
	var resp map[string]interface{}
	dto := &domain.VersionDTO{}

	err := json.NewDecoder(r.Body).Decode(dto)
	if err != nil {
		resp = utils.Message(http.StatusBadRequest, "Invalid request")
	} else {
		resp = controller.versioningService.ListChildren(dto.Name)
	}

	utils.Respond(w, resp)
}

func (controller *VersioningController) ListTreeBFS(w http.ResponseWriter, r *http.Request) {
	var resp map[string]interface{}
	dto := &domain.VersionDTO{}

	err := json.NewDecoder(r.Body).Decode(dto)
	if err != nil {
		resp = utils.Message(http.StatusBadRequest, "Invalid request")
	} else {
		resp = controller.versioningService.ListTreeBFS(dto.Name)
	}

	utils.Respond(w, resp)
}

func (controller *VersioningController) UpdateCode(w http.ResponseWriter, r *http.Request) {
	var resp map[string]interface{}
	dto := &domain.VersionDTO{}

	err := json.NewDecoder(r.Body).Decode(dto)
	if err != nil {
		resp = utils.Message(http.StatusBadRequest, "Invalid request")
	} else {
		resp = controller.versioningService.UpdateCode(dto.Name, dto.Code)
	}

	utils.Respond(w, resp)
}

func (controller *VersioningController) ListLibraries(w http.ResponseWriter, r *http.Request) {
	var resp map[string]interface{}
	dto := &domain.VersionDTO{}

	err := json.NewDecoder(r.Body).Decode(dto)
	if err != nil {
		resp = utils.Message(http.StatusBadRequest, "Invalid request")
	} else {
		resp = controller.versioningService.ListLibraries(dto.Name)
	}

	utils.Respond(w, resp)
}

func (controller *VersioningController) AddLibrary(w http.ResponseWriter, r *http.Request) {
	var resp map[string]interface{}
	dto := &domain.VersionDTO{}

	err := json.NewDecoder(r.Body).Decode(dto)
	if err != nil {
		resp = utils.Message(http.StatusBadRequest, "Invalid request")
	} else {
		resp = controller.versioningService.AddLibrary(dto.Name, dto.LibraryName, dto.LibraryVersion)
	}

	utils.Respond(w, resp)
}

func (controller *VersioningController) DeleteLibrary(w http.ResponseWriter, r *http.Request) {
	var resp map[string]interface{}
	dto := &domain.VersionDTO{}

	err := json.NewDecoder(r.Body).Decode(dto)
	if err != nil {
		resp = utils.Message(http.StatusBadRequest, "Invalid request")
	} else {
		resp = controller.versioningService.DeleteLibrary(dto.Name, dto.LibraryName, dto.LibraryVersion)
	}

	utils.Respond(w, resp)
}

func (controller *VersioningController) Delete(w http.ResponseWriter, r *http.Request) {
	var resp map[string]interface{}
	dto := &domain.VersionDTO{}

	err := json.NewDecoder(r.Body).Decode(dto)
	if err != nil {
		resp = utils.Message(http.StatusBadRequest, "Invalid request")
	} else {
		resp = controller.versioningService.Delete(dto.Name)
	}

	utils.Respond(w, resp)
}
