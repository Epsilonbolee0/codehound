package repository

import (
	"../domain"
	"gorm.io/gorm"
)

type VersionRepository struct {
	Conn *gorm.DB
}

func NewVersionRepository(conn *gorm.DB) *VersionRepository {
	return &VersionRepository{Conn: conn}
}

func (repo *VersionRepository) ListByAuthor(author domain.Account) ([]domain.Version, error) {
	var versions []domain.Version
	err := repo.Conn.Where("author = ?", author.ID).Find(&versions).Error
	return versions, err
}

func (repo *VersionRepository) Find(name string) (domain.Version, error) {
	var version domain.Version
	err := repo.Conn.Where("name = ?", name).First(&version).Error
	return version, err
}

func (repo *VersionRepository) Create(version domain.Version) error {
	return repo.Conn.Create(&version).Error
}

func (repo *VersionRepository) UpdateCode(name, code string) error {
	return repo.Conn.Model(&domain.Version{}).Where("name = ?", name).Update("code", code).Error
}

func (repo *VersionRepository) UpdateTitle(name, title string) error {
	return repo.Conn.Model(&domain.Version{}).Where("name = ?", name).Update("title", title).Error
}

func (repo *VersionRepository) ListLibraries(name string) ([]domain.Library, error) {
	var libraries []domain.Library
	repo.Conn.Model(&domain.Version{Name: name}).Association("Libraries").Find(&libraries)
	return libraries, nil
}

func (repo *VersionRepository) AddLibrary(name string, library domain.Library) error {
	repo.Conn.Model(&domain.Version{Name: name}).Association("Libraries").Append(&library)
	return nil
}

func (repo *VersionRepository) DeleteLibrary(name string, library domain.Library) error {
	repo.Conn.Model(&domain.Version{Name: name}).Association("Libraries").Delete(&library)
	return nil
}

func (repo *VersionRepository) Delete(name string) error {
	return repo.Conn.Where("name = ?", name).Delete(domain.Version{}).Error
}

func (repo *VersionRepository) LibraryIsValid(name string, library domain.Library) bool {
	var version domain.Version
	err := repo.Conn.Where("name = ? AND language_id = ?", name, library.LanguageID).First(&version).Error
	return err == nil && version.Name != ""
}
