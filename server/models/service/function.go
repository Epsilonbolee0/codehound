package service

import (
	"net/http"

	"../../utils"
	"../builder"
	"../repository"
	"gorm.io/gorm"
)

type FunctionService struct {
	accountRepo        *repository.AccountRepository
	implementationRepo *repository.ImplementationRepository
	languageRepo       *repository.LanguageRepository
	descriptionRepo    *repository.DescriptionRepository
}

func NewFunctionService(accountRepo *repository.AccountRepository,
	implementationRepo *repository.ImplementationRepository,
	langRepo *repository.LanguageRepository,
	descriptionRepo *repository.DescriptionRepository) *FunctionService {
	return &FunctionService{accountRepo, implementationRepo, langRepo, descriptionRepo}
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
	switch err {
	case nil:
		break
	case gorm.ErrRecordNotFound:
		return utils.Message(http.StatusNotFound, "Root is not found!")
	default:
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

	implBuilder := builder.NewImplementationBuilder()
	implementation := implBuilder.
		Name(name).
		Language(language.ID).
		Build()

	err = function.implementationRepo.Create(implementation)
	if err != nil {
		return utils.Message(http.StatusInternalServerError, "Error occured while creating implementation!")
	}

	return utils.Message(http.StatusOK, "Implementation created!")
}

func (function *FunctionService) ListDescriptions(impl string) map[string]interface{} {
	descriptions, err := function.descriptionRepo.ListByImplementation(impl)
	if err != nil {
		return utils.Message(http.StatusInternalServerError, "Failure occured while listing descriptions!")
	}

	response := utils.Message(http.StatusOK, "Descriptions listed!")
	response["descriptions"] = descriptions
	return response
}

func (function *FunctionService) AddDescription(login, impl, content string) map[string]interface{} {
	author, err := function.accountRepo.FindByLogin(login)
	switch err {
	case nil:
		break
	case gorm.ErrRecordNotFound:
		return utils.Message(http.StatusNotFound, "Author not found")
	default:
		return utils.Message(http.StatusInternalServerError, "Error occured while adding description")
	}

	descriptionBuilder := builder.NewDescriptionBuilder()
	description := descriptionBuilder.
		Author(author.ID).
		Implementation(impl).
		Content(content).
		Build()

	err = function.descriptionRepo.Create(description)
	if err != nil {
		return utils.Message(http.StatusInternalServerError, "Failure occured while adding tag!")
	}

	return utils.Message(http.StatusOK, "Tag added!")
}

func (function *FunctionService) UpdateDescription(id uint, content string) map[string]interface{} {
	err := function.descriptionRepo.Update(id, content)
	if err != nil {
		return utils.Message(http.StatusInternalServerError, "Failure occured while deleting description!")
	}

	return utils.Message(http.StatusOK, "Description deleted!")
}

func (function *FunctionService) DeleteDescription(id uint) map[string]interface{} {
	err := function.descriptionRepo.Delete(id)
	if err != nil {
		return utils.Message(http.StatusInternalServerError, "Failure occured while deleting description!")
	}

	return utils.Message(http.StatusOK, "Description deleted!")
}
