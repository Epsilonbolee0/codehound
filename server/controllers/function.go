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

	router.HandleFunc("/func/in/list", controller.ListInArguments).Methods("GET")
	router.HandleFunc("/func/in/add", controller.AddInArgument).Methods("PATCH")
	router.HandleFunc("/func/in/update", controller.UpdateInArgument).Methods("PATCH")
	router.HandleFunc("/func/in/delete", controller.DeleteInArgument).Methods("PATCH")

	router.HandleFunc("/func/out/list", controller.ListOutArguments).Methods("GET")
	router.HandleFunc("/func/out/add", controller.AddOutArgument).Methods("PATCH")
	router.HandleFunc("/func/out/update", controller.UpdateOutArgument).Methods("PATCH")
	router.HandleFunc("/func/out/delete", controller.DeleteOutArgument).Methods("PATCH")

	router.HandleFunc("/func/test/list", controller.ListTests).Methods("GET")
	router.HandleFunc("/func/test/add", controller.AddTest).Methods("POST")
	router.HandleFunc("/func/test/update", controller.UpdateDescriptionForTest).Methods("PATCH")
	router.HandleFunc("/func/test/in/update", controller.UpdateInArgumentForTest).Methods("PATCH")
	router.HandleFunc("/func/test/out/update", controller.UpdateOutArgumentForTest).Methods("PATCH")
	router.HandleFunc("/func/test/delete", controller.DeleteTest).Methods("DELETE")
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
		resp = controller.functionService.AddDescription(login, dto.Implementation, dto.Content)
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

func (controller *FunctionController) ListInArguments(w http.ResponseWriter, r *http.Request) {
	var resp map[string]interface{}
	dto := &domain.ImplementationDTO{}

	err := json.NewDecoder(r.Body).Decode(dto)

	if err != nil {
		resp = utils.Message(http.StatusBadRequest, "Invalid request")
	} else {
		resp = controller.functionService.ListInArguments(dto.Name)
	}

	utils.Respond(w, resp)
}

func (controller *FunctionController) ListOutArguments(w http.ResponseWriter, r *http.Request) {
	var resp map[string]interface{}
	dto := &domain.ImplementationDTO{}

	err := json.NewDecoder(r.Body).Decode(dto)

	if err != nil {
		resp = utils.Message(http.StatusBadRequest, "Invalid request")
	} else {
		resp = controller.functionService.ListOutArguments(dto.Name)
	}

	utils.Respond(w, resp)
}

func (controller *FunctionController) AddInArgument(w http.ResponseWriter, r *http.Request) {
	var resp map[string]interface{}
	dto := &domain.ImplementationDTO{}

	err := json.NewDecoder(r.Body).Decode(dto)

	if err != nil {
		resp = utils.Message(http.StatusBadRequest, "Invalid request")
	} else {
		resp = controller.functionService.AddInArgument(dto.Name, dto.Value)
	}

	utils.Respond(w, resp)
}

func (controller *FunctionController) AddOutArgument(w http.ResponseWriter, r *http.Request) {
	var resp map[string]interface{}
	dto := &domain.ImplementationDTO{}

	err := json.NewDecoder(r.Body).Decode(dto)

	if err != nil {
		resp = utils.Message(http.StatusBadRequest, "Invalid request")
	} else {
		resp = controller.functionService.AddOutArgument(dto.Name, dto.Value)
	}

	utils.Respond(w, resp)
}

func (controller *FunctionController) UpdateInArgument(w http.ResponseWriter, r *http.Request) {
	var resp map[string]interface{}
	dto := &domain.ImplementationDTO{}

	err := json.NewDecoder(r.Body).Decode(dto)

	if err != nil {
		resp = utils.Message(http.StatusBadRequest, "Invalid request")
	} else {
		resp = controller.functionService.UpdateInArgument(dto.Name, dto.Index, dto.Value)
	}

	utils.Respond(w, resp)
}

func (controller *FunctionController) UpdateOutArgument(w http.ResponseWriter, r *http.Request) {
	var resp map[string]interface{}
	dto := &domain.ImplementationDTO{}

	err := json.NewDecoder(r.Body).Decode(dto)

	if err != nil {
		resp = utils.Message(http.StatusBadRequest, "Invalid request")
	} else {
		resp = controller.functionService.UpdateOutArgument(dto.Name, dto.Index, dto.Value)
	}

	utils.Respond(w, resp)
}

func (controller *FunctionController) DeleteInArgument(w http.ResponseWriter, r *http.Request) {
	var resp map[string]interface{}
	dto := &domain.ImplementationDTO{}

	err := json.NewDecoder(r.Body).Decode(dto)

	if err != nil {
		resp = utils.Message(http.StatusBadRequest, "Invalid request")
	} else {
		resp = controller.functionService.DeleteInArgument(dto.Name, dto.Index)
	}

	utils.Respond(w, resp)
}

func (controller *FunctionController) DeleteOutArgument(w http.ResponseWriter, r *http.Request) {
	var resp map[string]interface{}
	dto := &domain.ImplementationDTO{}

	err := json.NewDecoder(r.Body).Decode(dto)

	if err != nil {
		resp = utils.Message(http.StatusBadRequest, "Invalid request")
	} else {
		resp = controller.functionService.DeleteOutArgument(dto.Name, dto.Index)
	}

	utils.Respond(w, resp)
}

func (controller *FunctionController) ListTests(w http.ResponseWriter, r *http.Request) {
	var resp map[string]interface{}
	dto := &domain.TestDTO{}

	err := json.NewDecoder(r.Body).Decode(dto)

	if err != nil {
		resp = utils.Message(http.StatusBadRequest, "Invalid request")
	} else {
		resp = controller.functionService.ListTests(dto.Implementation)
	}

	utils.Respond(w, resp)
}

func (controller *FunctionController) AddTest(w http.ResponseWriter, r *http.Request) {
	var resp map[string]interface{}
	dto := &domain.TestDTO{}

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
		resp = controller.functionService.AddTest(login, dto.Implementation, dto.Description)
	}

	utils.Respond(w, resp)
}

func (controller *FunctionController) UpdateDescriptionForTest(w http.ResponseWriter, r *http.Request) {
	var resp map[string]interface{}
	dto := &domain.TestDTO{}

	err := json.NewDecoder(r.Body).Decode(dto)

	if err != nil {
		resp = utils.Message(http.StatusBadRequest, "Invalid request")
	} else {
		resp = controller.functionService.UpdateDescriptionForTest(dto.ID, dto.Description)
	}

	utils.Respond(w, resp)
}

func (controller *FunctionController) UpdateInArgumentForTest(w http.ResponseWriter, r *http.Request) {
	var resp map[string]interface{}
	dto := &domain.TestDTO{}

	err := json.NewDecoder(r.Body).Decode(dto)

	if err != nil {
		resp = utils.Message(http.StatusBadRequest, "Invalid request")
	} else {
		resp = controller.functionService.UpdateInArgumentForTest(dto.ID, dto.Index, dto.Value)
	}

	utils.Respond(w, resp)
}

func (controller *FunctionController) UpdateOutArgumentForTest(w http.ResponseWriter, r *http.Request) {
	var resp map[string]interface{}
	dto := &domain.TestDTO{}

	err := json.NewDecoder(r.Body).Decode(dto)

	if err != nil {
		resp = utils.Message(http.StatusBadRequest, "Invalid request")
	} else {
		resp = controller.functionService.UpdateOutArgumentForTest(dto.ID, dto.Index, dto.Value)
	}

	utils.Respond(w, resp)
}

func (controller *FunctionController) DeleteTest(w http.ResponseWriter, r *http.Request) {
	var resp map[string]interface{}
	dto := &domain.TestDTO{}

	err := json.NewDecoder(r.Body).Decode(dto)

	if err != nil {
		resp = utils.Message(http.StatusBadRequest, "Invalid request")
	} else {
		resp = controller.functionService.DeleteTest(dto.ID)
	}

	utils.Respond(w, resp)
}
