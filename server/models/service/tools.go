package service

import (
	"net/http"

	"../../utils"
	"../factory"
	"../repository"
	"gorm.io/gorm"
)

type ToolsService struct {
	languageRepo *repository.LanguageRepository
	libraryRepo  *repository.LibraryRepository
}

func NewToolsService(languageRepo *repository.LanguageRepository, libraryRepo *repository.LibraryRepository) *ToolsService {
	return &ToolsService{languageRepo, libraryRepo}
}

func (tools *ToolsService) AddLanguage(name, version string) map[string]interface{} {
	builder := factory.NewLanguageBuilder()
	language := builder.
		Name(name).
		Version(version).
		Build()

	err := tools.languageRepo.Create(language)
	if err != nil {
		return utils.Message(http.StatusInternalServerError, "Failure occured while creating language!")
	}

	response := utils.Message(http.StatusOK, "Language has been created!")
	response["language"] = language
	return response
}

func (tools *ToolsService) AddLibrary(name, version, langName, langVersion string) map[string]interface{} {
	language, err := tools.languageRepo.FindByNameAndVersion(langName, langVersion)
	switch err {
	case nil:
		break
	case gorm.ErrUnsupportedRelation:
		return utils.Message(http.StatusNotFound, "Language not found!")
	default:
		return utils.Message(http.StatusInternalServerError, "Error occured while creating library!")
	}

	builder := factory.NewLibraryBuilder()
	library := builder.
		Name(name).
		Version(version).
		Language(language.ID).
		Build()

	err = tools.libraryRepo.Create(library)
	if err != nil {
		return utils.Message(http.StatusInternalServerError, "Error occured while creating library!")
	}

	response := utils.Message(http.StatusOK, "Library has been created!")
	response["library"] = library
	return response
}

func (tools *ToolsService) ListLanguages() map[string]interface{} {
	languages, err := tools.languageRepo.List()
	switch err {
	case nil:
		break
	case gorm.ErrRecordNotFound:
		return utils.Message(http.StatusNotFound, "No languages!")
	default:
		return utils.Message(http.StatusInternalServerError, "Failure occured while listing languages!")
	}

	response := utils.Message(http.StatusOK, "Language names listed!")
	response["languages"] = languages
	return response
}

func (tools *ToolsService) ListLibraries(name, version string) map[string]interface{} {
	libraries, err := tools.libraryRepo.FindByLanguage(name, version)
	switch err {
	case nil:
		break
	case gorm.ErrRecordNotFound:
		return utils.Message(http.StatusNotFound, "No libraries!")
	default:
		return utils.Message(http.StatusInternalServerError, "Failure occured while listing libraries!")
	}

	response := utils.Message(http.StatusOK, "Library names listed!")
	response["libraries"] = libraries
	return response
}

func (tools *ToolsService) ListLanguageVersions(name string) map[string]interface{} {
	versions, err := tools.languageRepo.VersionsOfLanguage(name)
	if err != nil {
		return utils.Message(http.StatusInternalServerError, "Failure occured while listing language versions!")
	}

	response := utils.Message(http.StatusOK, "Languages listed!")
	response["versions"] = versions
	return response
}

func (tools *ToolsService) ListLibraryVersions(name string) map[string]interface{} {
	versions, err := tools.libraryRepo.VersionsOfLibrary(name)
	if err != nil {
		return utils.Message(http.StatusInternalServerError, "Failure occured while listing language versions!")
	}

	response := utils.Message(http.StatusOK, "Libraries listed!")
	response["versions"] = versions
	return response
}

func (tools *ToolsService) DeleteLanguageVersion(name, version string) map[string]interface{} {
	err := tools.languageRepo.DeleteByNameAndVersion(name, version)
	if err != nil {
		return utils.Message(http.StatusInternalServerError, "Failure occured while deleting language version!")
	}

	return utils.Message(http.StatusOK, "Language version deleted!")
}

func (tools *ToolsService) DeleteLibraryVersion(name, version string) map[string]interface{} {
	err := tools.libraryRepo.DeleteByNameAndVersion(name, version)
	if err != nil {
		return utils.Message(http.StatusInternalServerError, "Failure occured while deleting library version!")
	}

	return utils.Message(http.StatusOK, "Library version deleted!")
}
