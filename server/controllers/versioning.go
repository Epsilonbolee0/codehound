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
	router.HandleFunc("/versioning/update_code", controller.UpdateCode).Methods("PATCH")
	router.HandleFunc("/versioning/delete", controller.Delete).Methods("DELETE")

	router.HandleFunc("/versioning/lib/list", controller.ListLibraries).Methods("GET")
	router.HandleFunc("/versioning/lib/add", controller.AddLibrary).Methods("PATCH")
	router.HandleFunc("/versioning/lib/delete", controller.DeleteLibrary).Methods("DELETE")

	router.HandleFunc("/versioning/tag/list", controller.ListTags).Methods("GET")
	router.HandleFunc("/versioning/tag/add", controller.AddTag).Methods("POST")
	router.HandleFunc("/versioning/tag/delete", controller.DeleteTag).Methods("DELETE")

	router.HandleFunc("/versioning/tree/bfs", controller.ListTreeBFS).Methods("GET")
	router.HandleFunc("/versioning/tree/parent", controller.FindParent).Methods("GET")
	router.HandleFunc("/versioning/tree/children", controller.ListChildren).Methods("GET")
	router.HandleFunc("/versioning/tree/add", controller.AddChildVersion).Methods("POST")
}

func (controller *VersioningController) AddVersion(w http.ResponseWriter, r *http.Request) {
	var resp map[string]interface{}
	dto := &domain.VersionDTO{}

	login, err := utils.LoginFromCookie(r)
	if err != nil {
		resp = utils.Message(http.StatusUnauthorized, "Cookie not found!")
		utils.Respond(w, resp)
		return
	}

	err = json.NewDecoder(r.Body).Decode(dto)
	if err != nil {
		resp = utils.Message(http.StatusBadRequest, "Invalid request")
	} else {
		resp = controller.versioningService.AddVersion(dto.Code, login, dto.Implementation)
	}

	utils.Respond(w, resp)
}

func (controller *VersioningController) AddChildVersion(w http.ResponseWriter, r *http.Request) {
	var resp map[string]interface{}
	dto := &domain.VersionDTO{}

	login, err := utils.LoginFromCookie(r)
	if err != nil {
		resp = utils.Message(http.StatusUnauthorized, "Cookie not found!")
		utils.Respond(w, resp)
		return
	}

	err = json.NewDecoder(r.Body).Decode(dto)
	if err != nil {
		resp = utils.Message(http.StatusBadRequest, "Invalid request")
	} else {
		resp = controller.versioningService.AddChildVersion(login, dto.Name, dto.Link)
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

func (controller *VersioningController) ListTags(w http.ResponseWriter, r *http.Request) {
	var resp map[string]interface{}
	dto := &domain.VersionDTO{}

	err := json.NewDecoder(r.Body).Decode(dto)
	if err != nil {
		resp = utils.Message(http.StatusBadRequest, "Invalid request")
	} else {
		resp = controller.versioningService.ListTags(dto.Name)
	}

	utils.Respond(w, resp)
}

func (controller *VersioningController) AddTag(w http.ResponseWriter, r *http.Request) {
	var resp map[string]interface{}
	dto := &domain.VersionDTO{}

	err := json.NewDecoder(r.Body).Decode(dto)
	if err != nil {
		resp = utils.Message(http.StatusBadRequest, "Invalid request")
	} else {
		resp = controller.versioningService.AddTag(dto.Name, dto.Category, dto.Content)
	}

	utils.Respond(w, resp)
}

func (controller *VersioningController) DeleteTag(w http.ResponseWriter, r *http.Request) {
	var resp map[string]interface{}
	dto := &domain.VersionDTO{}

	err := json.NewDecoder(r.Body).Decode(dto)
	if err != nil {
		resp = utils.Message(http.StatusBadRequest, "Invalid request")
	} else {
		resp = controller.versioningService.DeleteTag(dto.Name, dto.Category, dto.Content)
	}

	utils.Respond(w, resp)
}
