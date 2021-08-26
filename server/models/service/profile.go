package service

import (
	"net/http"

	"../../utils"
	"../repository"
	"gorm.io/gorm"
)

type ProfileService struct {
	accountRepo     *repository.AccountRepository
	testRepo        *repository.TestRepository
	versionRepo     *repository.VersionRepository
	descriptionRepo *repository.DescriptionRepository
}

func NewProfileService(accountRepo *repository.AccountRepository,
	testRepo *repository.TestRepository,
	versionRepo *repository.VersionRepository,
	descriptionRepo *repository.DescriptionRepository) *ProfileService {
	return &ProfileService{accountRepo, testRepo, versionRepo, descriptionRepo}
}

func (profile *ProfileService) Find(login string) map[string]interface{} {
	account, err := profile.accountRepo.FindByLogin(login)
	if err != nil {
		return utils.Message(http.StatusInternalServerError, "Failure occured while finding account!")
	}

	account.Password = ""
	response := utils.Message(http.StatusOK, "Account found!")
	response["account"] = account
	return response
}

func (profile *ProfileService) UpdateIcon(login, icon string) map[string]interface{} {
	err := profile.accountRepo.UpdateIcon(login, icon)
	if err != nil {
		return utils.Message(http.StatusInternalServerError, "Failure occured while updating icon!")
	}

	return utils.Message(http.StatusOK, "Icon updated!")
}

func (profile *ProfileService) UpdateRole(login, role string) map[string]interface{} {
	author, err := profile.accountRepo.FindByLogin(login)
	if err != nil {
		return utils.Message(http.StatusInternalServerError, "Failure occured while updating role!")
	}

	if author.Role != "admin" {
		return utils.Message(http.StatusForbidden, "Role update is forbidden!")
	}

	err = profile.accountRepo.UpdateRole(login, role)
	if err != nil {
		return utils.Message(http.StatusInternalServerError, "Failure occured while updating role!")
	}

	return utils.Message(http.StatusOK, "Role updated!")
}

func (profile *ProfileService) ListVersions(login string) map[string]interface{} {
	author, err := profile.accountRepo.FindByLogin(login)
	switch err {
	case nil:
		break
	case gorm.ErrRecordNotFound:
		return utils.Message(http.StatusNotFound, "Author not found")
	default:
		return utils.Message(http.StatusInternalServerError, "Error occured while listing versions!")
	}

	versions, err := profile.versionRepo.ListByAuthor(author.ID)
	switch err {
	case nil:
		break
	case gorm.ErrEmptySlice:
		return utils.Message(http.StatusNotFound, "No versions found!")
	default:
		return utils.Message(http.StatusInternalServerError, "Error occured while listing versions!")
	}

	response := utils.Message(http.StatusOK, "Versions listed!")
	response["versions"] = versions
	return response
}

func (profile *ProfileService) ListTests(login string) map[string]interface{} {
	author, err := profile.accountRepo.FindByLogin(login)
	switch err {
	case nil:
		break
	case gorm.ErrRecordNotFound:
		return utils.Message(http.StatusNotFound, "Author not found")
	default:
		return utils.Message(http.StatusInternalServerError, "Error occured while listing tests!")
	}

	versions, err := profile.testRepo.ListByAuthor(author.ID)
	switch err {
	case nil:
		break
	case gorm.ErrEmptySlice:
		return utils.Message(http.StatusNotFound, "No versions found!")
	default:
		return utils.Message(http.StatusInternalServerError, "Error occured while listing tests!")
	}

	response := utils.Message(http.StatusOK, "Tests listed!")
	response["tests"] = versions
	return response
}

func (profile *ProfileService) ListDescriptions(login string) map[string]interface{} {
	author, err := profile.accountRepo.FindByLogin(login)
	switch err {
	case nil:
		break
	case gorm.ErrRecordNotFound:
		return utils.Message(http.StatusNotFound, "Author not found")
	default:
		return utils.Message(http.StatusInternalServerError, "Error occured while listing descriptions!")
	}

	descriptions, err := profile.descriptionRepo.ListByAuthor(author.ID)
	switch err {
	case nil:
		break
	case gorm.ErrEmptySlice:
		return utils.Message(http.StatusNotFound, "No descriptions found!")
	default:
		return utils.Message(http.StatusInternalServerError, "Error occured while listing descriptions!")
	}

	response := utils.Message(http.StatusOK, "Descriptions listed!")
	response["descriptions"] = descriptions
	return response
}

func (profile *ProfileService) CountLevel(login string) map[string]interface{} {
	author, err := profile.accountRepo.FindByLogin(login)
	switch err {
	case nil:
		break
	case gorm.ErrRecordNotFound:
		return utils.Message(http.StatusNotFound, "Author not found")
	default:
		return utils.Message(http.StatusInternalServerError, "Error occured while listing descriptions!")
	}

	versions, err := profile.versionRepo.ListByAuthor(author.ID)
	if err != nil {
		return utils.Message(http.StatusInternalServerError, "Error occured while counting level!")
	}

	tests, err := profile.testRepo.ListByAuthor(author.ID)
	if err != nil {
		return utils.Message(http.StatusInternalServerError, "Error occured while counting level!")
	}

	descriptions, err := profile.descriptionRepo.ListByAuthor(author.ID)
	if err != nil {
		return utils.Message(http.StatusInternalServerError, "Error occured while counting level!")
	}

	totalPoins := 5*len(versions) + 2*len(tests) + len(descriptions)

	response := utils.Message(http.StatusOK, "Level counted!")
	response["level"] = totalPoins / 10
	response["tenth"] = totalPoins % 10
	return response
}
