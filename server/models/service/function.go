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
	testRepo           *repository.TestRepository
}

func NewFunctionService(accountRepo *repository.AccountRepository,
	implementationRepo *repository.ImplementationRepository,
	langRepo *repository.LanguageRepository,
	descriptionRepo *repository.DescriptionRepository,
	testRepo *repository.TestRepository) *FunctionService {
	return &FunctionService{accountRepo, implementationRepo, langRepo, descriptionRepo, testRepo}
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
	language, err := function.languageRepo.Find(langName, langVersion)
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
		return utils.Message(http.StatusInternalServerError, "Failure occured while adding description!")
	}

	return utils.Message(http.StatusOK, "Description added!")
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

func (function *FunctionService) ListInArguments(name string) map[string]interface{} {
	inArgs, err := function.implementationRepo.GetArguments("in_args", name)
	if err != nil {
		return utils.Message(http.StatusInternalServerError, "Failure occured while listing in arguments!")
	}

	response := utils.Message(http.StatusOK, "In arguments listed!")
	response["in_args"] = inArgs
	return response
}

func (function *FunctionService) ListOutArguments(name string) map[string]interface{} {
	outArgs, err := function.implementationRepo.GetArguments("out_args", name)
	if err != nil {
		return utils.Message(http.StatusInternalServerError, "Failure occured while listing out arguments!")
	}

	response := utils.Message(http.StatusOK, "Out arguments listed!")
	response["out_args"] = outArgs
	return response
}

func (function *FunctionService) AddInArgument(name, value string) map[string]interface{} {
	err := function.implementationRepo.AddArgument("in_args", name, value)
	if err != nil {
		return utils.Message(http.StatusInternalServerError, "Failure occured while adding in argument!")
	}

	return utils.Message(http.StatusOK, "In argument added!")
}

func (function *FunctionService) AddOutArgument(name, value string) map[string]interface{} {
	err := function.implementationRepo.AddArgument("out_args", name, value)
	if err != nil {
		return utils.Message(http.StatusInternalServerError, "Failure occured while adding out argument!")
	}

	return utils.Message(http.StatusOK, "Out argument added!")
}

func (function *FunctionService) UpdateInArgument(name string, index uint, value string) map[string]interface{} {
	err := function.implementationRepo.UpdateArgument("in_args", name, index, value)
	if err != nil {
		return utils.Message(http.StatusInternalServerError, "Failure occured while updating in argument!")
	}

	return utils.Message(http.StatusOK, "In argument updated!")
}

func (function *FunctionService) UpdateOutArgument(name string, index uint, value string) map[string]interface{} {
	err := function.implementationRepo.UpdateArgument("out_args", name, index, value)
	if err != nil {
		return utils.Message(http.StatusInternalServerError, "Failure occured while updating out argument!")
	}

	return utils.Message(http.StatusOK, "Out argument updated!")
}

func (function *FunctionService) DeleteInArgument(name string, index uint) map[string]interface{} {
	err := function.implementationRepo.RemoveArgument("in_args", name, index)
	if err != nil {
		return utils.Message(http.StatusInternalServerError, "Failure occured while deleting in argument!")
	}

	return utils.Message(http.StatusOK, "In argument deleted!")
}

func (function *FunctionService) DeleteOutArgument(name string, index uint) map[string]interface{} {
	err := function.implementationRepo.RemoveArgument("out_args", name, index)
	if err != nil {
		return utils.Message(http.StatusInternalServerError, "Failure occured while deleting out argument!")
	}

	return utils.Message(http.StatusOK, "Out argument deleted!")
}

func (function *FunctionService) ListTests(impl string) map[string]interface{} {
	tests, err := function.testRepo.List(impl)
	if err != nil {
		return utils.Message(http.StatusInternalServerError, "Failure occured while listing tests!")
	}

	response := utils.Message(http.StatusOK, "Tests listed!")
	response["tests"] = tests
	return response
}

func (function *FunctionService) AddTest(login, impl, description string) map[string]interface{} {
	author, err := function.accountRepo.FindByLogin(login)
	switch err {
	case nil:
		break
	case gorm.ErrRecordNotFound:
		return utils.Message(http.StatusNotFound, "Author not found")
	default:
		return utils.Message(http.StatusInternalServerError, "Error occured while adding test")
	}

	inCount, err := function.implementationRepo.CountArguments("in_args", impl)
	if err != nil {
		return utils.Message(http.StatusInternalServerError, "Failure occured while adding test!")
	}

	outCount, err := function.implementationRepo.CountArguments("out_args", impl)
	if err != nil {
		return utils.Message(http.StatusInternalServerError, "Failure occured while adding test!")
	}

	testBuilder := builder.NewTestBuilder()
	test := testBuilder.
		Author(author.ID).
		Implementation(impl).
		InArgs(inCount).
		OutArgs(outCount).
		Description(description).
		Build()

	err = function.testRepo.Create(test)
	if err != nil {
		return utils.Message(http.StatusInternalServerError, "Failure occured while adding test!")
	}

	return utils.Message(http.StatusOK, "Test added!")
}

func (function *FunctionService) UpdateDescriptionForTest(id uint, description string) map[string]interface{} {
	err := function.testRepo.UpdateDescription(id, description)
	if err != nil {
		return utils.Message(http.StatusInternalServerError, "Failure occured while updating test description!")
	}

	return utils.Message(http.StatusOK, "Test description updated")
}

func (function *FunctionService) UpdateInArgumentForTest(id, index uint, value string) map[string]interface{} {
	err := function.testRepo.UpdateArgument("in_args", id, index, value)
	if err != nil {
		return utils.Message(http.StatusInternalServerError, "Failure occured while updating test in argument!")
	}

	return utils.Message(http.StatusOK, "Test in argument updated")
}

func (function *FunctionService) UpdateOutArgumentForTest(id, index uint, value string) map[string]interface{} {
	err := function.testRepo.UpdateArgument("out_args", id, index, value)
	if err != nil {
		return utils.Message(http.StatusInternalServerError, "Failure occured while updating test out argument!")
	}

	return utils.Message(http.StatusOK, "Test out argument updated")
}

func (function *FunctionService) DeleteTest(id uint) map[string]interface{} {
	err := function.testRepo.Delete(id)
	if err != nil {
		return utils.Message(http.StatusInternalServerError, "Failure occured while deleting test!")
	}

	return utils.Message(http.StatusOK, "Test deleted")
}
