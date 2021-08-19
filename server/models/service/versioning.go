package service

import (
	"crypto/sha1"
	"encoding/hex"
	"net/http"
	"strconv"
	"time"

	"../../utils"
	"../factory"
	"../repository"
	"gorm.io/gorm"
)

type VersioningService struct {
	accountRepo *repository.AccountRepository
	versionRepo *repository.VersionRepository
}

func NewVersioningService(accountRepo *repository.AccountRepository, versionRepo *repository.VersionRepository) *VersioningService {
	return &VersioningService{accountRepo, versionRepo}
}

func (versioning *VersioningService) ListByAuthor(login string) map[string]interface{} {
	author, err := versioning.accountRepo.FindByLogin(login)
	switch err {
	case nil:
		break
	case gorm.ErrRecordNotFound:
		return utils.Message(http.StatusNotFound, "No author!")
	default:
		return utils.Message(http.StatusInternalServerError, "Error occured while listing versions")
	}

	versions, err := versioning.versionRepo.ListByAuthor(author)
	switch err {
	case nil:
		break
	case gorm.ErrRecordNotFound:
		return utils.Message(http.StatusNotFound, "No version!")
	default:
		return utils.Message(http.StatusInternalServerError, "Error occured while listing versions")
	}

	response := utils.Message(http.StatusOK, "Versions listed!")
	response["versions"] = versions
	return response
}

func (versioning *VersioningService) AddVersion(title, code, login string) map[string]interface{} {
	author, err := versioning.accountRepo.FindByLogin(login)
	switch err {
	case nil:
		break
	case gorm.ErrRecordNotFound:
		return utils.Message(http.StatusNotFound, "Version not found")
	default:
		return utils.Message(http.StatusInternalServerError, "Error occured while adding versions")
	}

	builder := factory.NewVersionBuilder()
	hash := versioning.generateName(author.Login)
	now := time.Now()

	version := builder.
		Date(now).
		Name(hash).
		Code(code).
		Title(title).
		Author(author.ID).
		Build()

	err = versioning.versionRepo.Create(version)

	if err != nil {
		return utils.Message(http.StatusInternalServerError, "Error occured while adding versions!")
	}

	return utils.Message(http.StatusOK, "Version added!")
}

func (versioning *VersioningService) generateName(login string) string {
	stringToHash := login + strconv.FormatInt(time.Now().Unix(), 10)
	algorithm := sha1.New()
	algorithm.Write(([]byte(stringToHash)))
	return hex.EncodeToString(algorithm.Sum(nil))[0:6]
}

func (versioning *VersioningService) UpdateCode(name, code string) map[string]interface{} {
	err := versioning.versionRepo.UpdateCode(name, code)
	if err != nil {
		return utils.Message(http.StatusInternalServerError, "Error occured while updating version's code!")
	}

	return utils.Message(http.StatusOK, "Version code successfully updated!")
}

func (versioning *VersioningService) UpdateTitle(name, title string) map[string]interface{} {
	err := versioning.versionRepo.UpdateTitle(name, title)
	if err != nil {
		return utils.Message(http.StatusInternalServerError, "Error occured while updating version's title!")
	}

	return utils.Message(http.StatusOK, "Version title successfully updated!")
}

func (versioning *VersioningService) Delete(name string) map[string]interface{} {
	err := versioning.versionRepo.Delete(name)
	if err != nil {
		return utils.Message(http.StatusInternalServerError, "Error occured while deleting versions!")
	}

	return utils.Message(http.StatusOK, "Version successfully deleted!")
}
