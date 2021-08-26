package controllers

import (
	"encoding/json"
	"net/http"

	"../models/domain"
	"../models/service"
	"../utils"
	"github.com/gorilla/mux"
)

type ProfileController struct {
	profileService *service.ProfileService
}

func SetupProfileController(profileService *service.ProfileService, router *mux.Router) {
	controller := &ProfileController{profileService: profileService}

	router.HandleFunc("/profile/home", controller.Home).Methods("GET")
	router.HandleFunc("/profile/update/icon", controller.UpdateIcon).Methods("PATCH")
	router.HandleFunc("/profile/update/role", controller.UpdateRole).Methods("PATCH")

	router.HandleFunc("/profile/list/tests", controller.ListTests).Methods("GET")
	router.HandleFunc("/profile/list/versions", controller.ListVersions).Methods("GET")
	router.HandleFunc("/profile/list/descriptions", controller.ListDescriptions).Methods("GET")
	router.HandleFunc("/profile/level", controller.CountLevel).Methods("GET")
}

func (controller *ProfileController) Home(w http.ResponseWriter, r *http.Request) {
	var resp map[string]interface{}

	login, err := utils.LoginFromCookie(r)
	if err != nil {
		resp = utils.Message(http.StatusUnauthorized, "Cookie not found!")
	} else {
		resp = controller.profileService.Find(login)
	}

	utils.Respond(w, resp)
}

func (controller *ProfileController) UpdateIcon(w http.ResponseWriter, r *http.Request) {
	var resp map[string]interface{}
	dto := &domain.AccountDTO{}

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
		resp = controller.profileService.UpdateIcon(login, dto.Icon)
	}

	utils.Respond(w, resp)
}

func (controller *ProfileController) UpdateRole(w http.ResponseWriter, r *http.Request) {
	var resp map[string]interface{}
	dto := &domain.AccountDTO{}

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
		resp = controller.profileService.UpdateRole(login, dto.Role)
	}

	utils.Respond(w, resp)
}

func (controller *ProfileController) ListVersions(w http.ResponseWriter, r *http.Request) {
	var resp map[string]interface{}

	login, err := utils.LoginFromCookie(r)

	if err != nil {
		resp = utils.Message(http.StatusUnauthorized, "Cookie not found!")
	} else {
		resp = controller.profileService.ListVersions(login)
	}

	utils.Respond(w, resp)
}

func (controller *ProfileController) ListTests(w http.ResponseWriter, r *http.Request) {
	var resp map[string]interface{}

	login, err := utils.LoginFromCookie(r)

	if err != nil {
		resp = utils.Message(http.StatusUnauthorized, "Cookie not found!")
	} else {
		resp = controller.profileService.ListTests(login)
	}

	utils.Respond(w, resp)
}

func (controller *ProfileController) ListDescriptions(w http.ResponseWriter, r *http.Request) {
	var resp map[string]interface{}

	login, err := utils.LoginFromCookie(r)

	if err != nil {
		resp = utils.Message(http.StatusUnauthorized, "Cookie not found!")
	} else {
		resp = controller.profileService.ListDescriptions(login)
	}

	utils.Respond(w, resp)
}

func (controller *ProfileController) CountLevel(w http.ResponseWriter, r *http.Request) {
	var resp map[string]interface{}

	login, err := utils.LoginFromCookie(r)

	if err != nil {
		resp = utils.Message(http.StatusUnauthorized, "Cookie not found!")
	} else {
		resp = controller.profileService.CountLevel(login)
	}

	utils.Respond(w, resp)
}
