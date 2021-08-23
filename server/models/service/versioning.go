package service

import (
	"crypto/sha1"
	"encoding/hex"
	"net/http"
	"strconv"
	"time"

	"../../utils"
	"../builder"
	"../repository"
	"gorm.io/gorm"
)

type VersioningService struct {
	accountRepo *repository.AccountRepository
	versionRepo *repository.VersionRepository
	implRepo    *repository.ImplementationRepository
	libraryRepo *repository.LibraryRepository
	treeRepo    *repository.TreeRepository
	tagRepo     *repository.TagRepository
}

func NewVersioningService(
	accountRepo *repository.AccountRepository,
	versionRepo *repository.VersionRepository,
	implRepo *repository.ImplementationRepository,
	libraryRepo *repository.LibraryRepository,
	treeRepo *repository.TreeRepository,
	tagRepo *repository.TagRepository) *VersioningService {
	return &VersioningService{accountRepo, versionRepo, implRepo, libraryRepo, treeRepo, tagRepo}
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

func (versioning *VersioningService) AddVersion(code, login, implName string) map[string]interface{} {
	author, err := versioning.accountRepo.FindByLogin(login)
	switch err {
	case nil:
		break
	case gorm.ErrRecordNotFound:
		return utils.Message(http.StatusNotFound, "Version not found")
	default:
		return utils.Message(http.StatusInternalServerError, "Error occured while adding versions")
	}

	_, err = versioning.implRepo.Find(implName)
	switch err {
	case nil:
		break
	case gorm.ErrRecordNotFound:
		return utils.Message(http.StatusNotFound, "Implementation not found")
	default:
		return utils.Message(http.StatusInternalServerError, "Error occured while adding versions")
	}

	versionBuilder := builder.NewVersionBuilder()
	hash := versioning.generateName(author.Login)
	now := time.Now()

	version := versionBuilder.
		Date(now).
		Name(hash).
		Code(code).
		Author(author.ID).
		Implementation(implName).
		Build()

	err = versioning.versionRepo.Create(version)
	if err != nil {
		return utils.Message(http.StatusInternalServerError, "Error occured while adding versions!")
	}

	treeBuilder := builder.NewTreeBuilder()
	tree := treeBuilder.
		Name(hash).
		Build()

	err = versioning.treeRepo.Create(tree)
	if err != nil {
		return utils.Message(http.StatusInternalServerError, "Error occured while adding versions!")
	}

	return utils.Message(http.StatusOK, "Version added!")
}

func (versioning *VersioningService) AddChildVersion(login, nameOfOrigin, link string) map[string]interface{} {
	author, err := versioning.accountRepo.FindByLogin(login)
	switch err {
	case nil:
		break
	case gorm.ErrRecordNotFound:
		return utils.Message(http.StatusNotFound, "Version not found")
	default:
		return utils.Message(http.StatusInternalServerError, "Error occured while adding versions")
	}

	version, err := versioning.versionRepo.Find(nameOfOrigin)
	switch err {
	case nil:
		break
	case gorm.ErrRecordNotFound:
		return utils.Message(http.StatusNotFound, "Version not found")
	default:
		return utils.Message(http.StatusInternalServerError, "Error occured while adding versions")
	}

	versionBuilder := builder.NewVersionBuilder()
	hash := versioning.generateName(author.Login)
	now := time.Now()

	newVersion := versionBuilder.
		Date(now).
		Name(hash).
		Code(version.Code).
		Author(author.ID).
		Implementation(version.Implementation).
		Build()

	err = versioning.versionRepo.Create(newVersion)
	if err != nil {
		return utils.Message(http.StatusInternalServerError, "Error occured while adding versions!")
	}

	treeBuilder := builder.NewTreeBuilder()
	tree := treeBuilder.
		Name(hash).
		ParentName(version.Name).
		Link(link).
		Build()

	err = versioning.treeRepo.Create(tree)
	if err != nil {
		return utils.Message(http.StatusInternalServerError, "Error occured while adding versions!")
	}

	return utils.Message(http.StatusOK, "Version added!")
}

func (versioning *VersioningService) FindParent(version string) map[string]interface{} {
	parent, err := versioning.treeRepo.FindParent(version)
	if err != nil {
		return utils.Message(http.StatusInternalServerError, "Failure occured while finding parents(!")
	}

	response := utils.Message(http.StatusOK, "Parent found!")
	response["parent"] = parent
	return response
}

func (versioning *VersioningService) ListChildren(version string) map[string]interface{} {
	children, err := versioning.treeRepo.ListChildren(version)
	if err != nil {
		return utils.Message(http.StatusInternalServerError, "Failure occured while listing children!")
	}

	response := utils.Message(http.StatusOK, "Children found!")
	response["children"] = children
	return response
}

func (versioning *VersioningService) ListTreeBFS(version string) map[string]interface{} {
	tree, err := versioning.treeRepo.ListTreeBFS(version)
	if err != nil {
		return utils.Message(http.StatusInternalServerError, "Failure occured while listing tree!")
	}

	response := utils.Message(http.StatusOK, "Tree listed!")
	response["tree"] = tree
	return response
}

func (versioning *VersioningService) generateName(login string) string {
	stringToHash := login + strconv.FormatInt(time.Now().Unix(), 10)
	algorithm := sha1.New()
	algorithm.Write(([]byte(stringToHash)))
	return hex.EncodeToString(algorithm.Sum(nil))[:6]
}

func (versioning *VersioningService) UpdateCode(name, code string) map[string]interface{} {
	err := versioning.versionRepo.UpdateCode(name, code)
	if err != nil {
		return utils.Message(http.StatusInternalServerError, "Error occured while updating version's code!")
	}

	return utils.Message(http.StatusOK, "Version code successfully updated!")
}

func (versioning *VersioningService) Delete(name string) map[string]interface{} {
	err := versioning.versionRepo.ClearLibraries(name)
	if err != nil {
		return utils.Message(http.StatusInternalServerError, "Error occured while deleting versions!")
	}

	err = versioning.versionRepo.ClearTags(name)
	if err != nil {
		return utils.Message(http.StatusInternalServerError, "Error occured while deleting versions!")
	}

	err = versioning.versionRepo.Delete(name)
	if err != nil {
		return utils.Message(http.StatusInternalServerError, "Error occured while deleting versions!")
	}

	return utils.Message(http.StatusOK, "Version successfully deleted!")
}

func (versioning *VersioningService) ListLibraries(name string) map[string]interface{} {
	libraries, err := versioning.versionRepo.ListLibraries(name)
	if err != nil {
		return utils.Message(http.StatusInternalServerError, "Failure occured while listing libraries!")
	}

	response := utils.Message(http.StatusOK, "Libraries listed!")
	response["libraries"] = libraries
	return response
}

func (versioning *VersioningService) AddLibrary(name, libName, libVersion string) map[string]interface{} {
	library, err := versioning.libraryRepo.Find(libName, libVersion)
	switch err {
	case nil:
		break
	case gorm.ErrRecordNotFound:
		return utils.Message(http.StatusNotFound, "No library!")
	default:
		return utils.Message(http.StatusInternalServerError, "Failure occured while adding library!")
	}

	ok := versioning.versionRepo.LibraryIsValid(name, library)
	if !ok {
		return utils.Message(http.StatusInternalServerError, "Library does not match chosen language!")
	}

	err = versioning.versionRepo.AddLibrary(name, library)
	if err != nil {
		return utils.Message(http.StatusInternalServerError, "Failure occured while adding library!")
	}

	return utils.Message(http.StatusOK, "Library added!")
}

func (versioning *VersioningService) DeleteLibrary(name, libName, libVersion string) map[string]interface{} {
	library, err := versioning.libraryRepo.Find(libName, libVersion)
	switch err {
	case nil:
		break
	case gorm.ErrRecordNotFound:
		return utils.Message(http.StatusNotFound, "No library!")
	default:
		return utils.Message(http.StatusInternalServerError, "Failure occured while deleting library!")
	}

	err = versioning.versionRepo.DeleteLibrary(name, library)
	if err != nil {
		return utils.Message(http.StatusInternalServerError, "Failure occured while deleting library!")
	}

	return utils.Message(http.StatusOK, "Library deleted!")
}

func (versioning *VersioningService) ListTags(name string) map[string]interface{} {
	tags, err := versioning.versionRepo.ListTags(name)
	if err != nil {
		return utils.Message(http.StatusInternalServerError, "Failure occured while listing tags!")
	}

	response := utils.Message(http.StatusOK, "Tags listed!")
	response["tags"] = tags
	return response
}

func (versioning *VersioningService) AddTag(name, category, content string) map[string]interface{} {
	tag, err := versioning.tagRepo.FindOrCreate(category, content)
	if err != nil {
		return utils.Message(http.StatusInternalServerError, "Failure occured while adding tag!")
	}

	err = versioning.versionRepo.AddTag(name, tag)
	if err != nil {
		return utils.Message(http.StatusInternalServerError, "Failure occured while adding tag!")
	}

	return utils.Message(http.StatusOK, "Tag added!")
}

func (versioning *VersioningService) DeleteTag(name, category, content string) map[string]interface{} {
	tag, err := versioning.tagRepo.Find(category, content)
	switch err {
	case nil:
		break
	case gorm.ErrRecordNotFound:
		return utils.Message(http.StatusNotFound, "No tag!")
	default:
		return utils.Message(http.StatusInternalServerError, "Failure occured while deleting tag!")
	}

	err = versioning.versionRepo.DeleteTag(name, tag)
	if err != nil {
		return utils.Message(http.StatusInternalServerError, "Failure occured while deleting tag!")
	}

	return utils.Message(http.StatusOK, "Tag deleted!")
}
