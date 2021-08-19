package repository

import (
	"../domain"
	"gorm.io/gorm"
)

type LibraryRepository struct {
	Conn *gorm.DB
}

func NewLibraryRepository(conn *gorm.DB) *LibraryRepository {
	return &LibraryRepository{Conn: conn}
}

func (repo *LibraryRepository) FindByLanguage(name, version string) ([]string, error) {
	var libraries []string

	subQuery := repo.Conn.Model(&domain.Language{}).Where("name = ? AND version = ?", name, version).Select("id")
	err := repo.Conn.Model(&domain.Library{}).Where("language_id = (?)", subQuery).Distinct().Pluck("name", &libraries).Error
	return libraries, err
}

func (repo *LibraryRepository) VersionsOfLibrary(name string) ([]string, error) {
	var versions []string
	err := repo.Conn.Model(&domain.Library{}).Where("name = ?", name).Pluck("version", &versions).Error
	return versions, err
}

func (repo *LibraryRepository) Create(library domain.Library) error {
	return repo.Conn.Create(&library).Error
}

func (repo *LibraryRepository) DeleteByNameAndVersion(name, version string) error {
	return repo.Conn.Where("name = ? AND version = ?", name, version).Delete(domain.Library{}).Error
}
