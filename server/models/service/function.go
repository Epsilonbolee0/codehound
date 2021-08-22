package service

import (
	"net/http"

	"../../utils"
	"../factory"
	"../repository"
	"gorm.io/gorm"
)

type FunctionService struct {
	implementationRepo *repository.ImplementationRepository
	languageRepo       *repository.LanguageRepository
}

func NewFunctionService(implRepo *repository.ImplementationRepository, langRepo *repository.LanguageRepository) *FunctionService {
	return &FunctionService{implRepo, langRepo}
}

func (function *FunctionService) List() map[string]interface{} {
	impls, err := function.implementationRepo.List()
	if err != nil {
		return utils.Message(http.StatusInternalServerError, "Failure occured while listing functions!")
	}

	response := utils.Message(http.StatusOK, "Functions listed!")
	response["functions"] = impls
	return response
}

func (function *FunctionService) Find(name string) map[string]interface{} {
	impls, err := function.implementationRepo.Find(name)
	if err != nil {
		return utils.Message(http.StatusInternalServerError, "Failure occured while finding function!")
	}

	response := utils.Message(http.StatusOK, "Function found!")
	response["function"] = impls
	return response
}

func (function *FunctionService) FindRoot(name string) map[string]interface{} {
	root, err := function.implementationRepo.FindRoot(name)
	if err != nil {
		return utils.Message(http.StatusInternalServerError, "Failure occured while finding root!")
	}

	response := utils.Message(http.StatusOK, "Root found!")
	response["root"] = root
	return response
}

func (function *FunctionService) Create(name, langName, langVersion string) map[string]interface{} {
	language, err := function.languageRepo.FindByNameAndVersion(langName, langVersion)
	switch err {
	case nil:
		break
	case gorm.ErrRecordNotFound:
		return utils.Message(http.StatusNotFound, "Language not found!")
	default:
		return utils.Message(http.StatusInternalServerError, "Error occured while creating implementation!")
	}

	builder := factory.NewImplementationBuilder()
	implementation := builder.
		Name(name).
		Language(language.ID).
		Build()

	err = function.implementationRepo.Create(implementation)
	if err != nil {
		return utils.Message(http.StatusInternalServerError, "Error occured while creating implementation!")
	}

	return utils.Message(http.StatusOK, "Implementation created!")
}
