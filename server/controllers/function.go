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
	router.HandleFunc("/func/add", controller.Add).Methods("POST")
	router.HandleFunc("/func/list", controller.List).Methods("GET")
	router.HandleFunc("/func/find", controller.Find).Methods("GET")
	router.HandleFunc("/func/find_root", controller.FindRoot).Methods("GET")

	router.HandleFunc("/func/descr/add", controller.AddDescription).Methods("POST")
	router.HandleFunc("/func/descr/list", controller.ListDescriptions).Methods("GET")
	router.HandleFunc("/func/descr/update", controller.UpdateDescription).Methods("PATCH")
	router.HandleFunc("/func/descr/delete", controller.DeleteDescription).Methods("DELETE")
}

func (controller *FunctionController) Add(w http.ResponseWriter, r *http.Request) {
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

func (controller *FunctionController) ListDescriptions(w http.ResponseWriter, r *http.Request) {
	var resp map[string]interface{}
	dto := &domain.DescriptionDTO{}

	err := json.NewDecoder(r.Body).Decode(dto)

	if err != nil {
		resp = utils.Message(http.StatusBadRequest, "Invalid request")
	} else {
		resp = controller.functionService.ListDescriptions(dto.Implementation)
	}

	utils.Respond(w, resp)
}

func (controller *FunctionController) AddDescription(w http.ResponseWriter, r *http.Request) {
	var resp map[string]interface{}
	dto := &domain.DescriptionDTO{}

	err := json.NewDecoder(r.Body).Decode(dto)

	if err != nil {
		resp = utils.Message(http.StatusBadRequest, "Invalid request")
	} else {
		resp = controller.functionService.AddDescription(dto.Login, dto.Implementation, dto.Content)
	}

	utils.Respond(w, resp)
}

func (controller *FunctionController) UpdateDescription(w http.ResponseWriter, r *http.Request) {
	var resp map[string]interface{}
	dto := &domain.DescriptionDTO{}

	err := json.NewDecoder(r.Body).Decode(dto)

	if err != nil {
		resp = utils.Message(http.StatusBadRequest, "Invalid request")
	} else {
		resp = controller.functionService.UpdateDescription(dto.ID, dto.Content)
	}

	utils.Respond(w, resp)
}

func (controller *FunctionController) DeleteDescription(w http.ResponseWriter, r *http.Request) {
	var resp map[string]interface{}
	dto := &domain.DescriptionDTO{}

	err := json.NewDecoder(r.Body).Decode(dto)

	if err != nil {
		resp = utils.Message(http.StatusBadRequest, "Invalid request")
	} else {
		resp = controller.functionService.DeleteDescription(dto.ID)
	}

	utils.Respond(w, resp)
}
