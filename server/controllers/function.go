package controllers

import (
	"encoding/json"
	"net/http"

	"../models/domain"
	"../models/service"
	"../utils"
	"github.com/gorilla/mux"
)

type FunctionController struct {
	functionService *service.FunctionService
}

func SetupFunctionController(functionService *service.FunctionService, router *mux.Router) {
	controller := &FunctionController{functionService: functionService}
	router.HandleFunc("/func/add", controller.AddImplementation).Methods("POST")
	router.HandleFunc("/func/list", controller.List).Methods("GET")
	router.HandleFunc("/func/find", controller.Find).Methods("GET")
	router.HandleFunc("/func/find_root", controller.FindRoot).Methods("GET")
}

func (controller *FunctionController) AddImplementation(w http.ResponseWriter, r *http.Request) {
	var resp map[string]interface{}
	dto := &domain.ImplementationDTO{}

	err := json.NewDecoder(r.Body).Decode(dto)

	if err != nil {
		resp = utils.Message(http.StatusBadRequest, "Invalid request")
	} else {
		resp = controller.functionService.Create(dto.Name, dto.LanguageName, dto.LanguageVersion)
	}

	utils.Respond(w, resp)
}

func (controller *FunctionController) List(w http.ResponseWriter, r *http.Request) {
	utils.Respond(w, controller.functionService.List())
}

func (controller *FunctionController) Find(w http.ResponseWriter, r *http.Request) {
	var resp map[string]interface{}
	dto := &domain.ImplementationDTO{}

	err := json.NewDecoder(r.Body).Decode(dto)

	if err != nil {
		resp = utils.Message(http.StatusBadRequest, "Invalid request")
	} else {
		resp = controller.functionService.Find(dto.Name)
	}

	utils.Respond(w, resp)
}

func (controller *FunctionController) FindRoot(w http.ResponseWriter, r *http.Request) {
	var resp map[string]interface{}
	dto := &domain.ImplementationDTO{}

	err := json.NewDecoder(r.Body).Decode(dto)

	if err != nil {
		resp = utils.Message(http.StatusBadRequest, "Invalid request")
	} else {
		resp = controller.functionService.FindRoot(dto.Name)
	}

	utils.Respond(w, resp)
}
