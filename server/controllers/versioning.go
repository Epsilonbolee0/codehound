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
	router.HandleFunc("/versioning/update_title", controller.UpdateTitle).Methods("PATCH")
	router.HandleFunc("/versioning/delete", controller.Delete).Methods("DELETE")
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
		resp = controller.versioningService.AddVersion(dto.Title, dto.Code, dto.Login)
	}

	utils.Respond(w, resp)
}

func (controller *VersioningController) UpdateTitle(w http.ResponseWriter, r *http.Request) {
	var resp map[string]interface{}
	dto := &domain.VersionDTO{}

	err := json.NewDecoder(r.Body).Decode(dto)
	if err != nil {
		resp = utils.Message(http.StatusBadRequest, "Invalid request")
	} else {
		resp = controller.versioningService.UpdateTitle(dto.Name, dto.Title)
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
