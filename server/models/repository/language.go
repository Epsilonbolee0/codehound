package repository

import (
	"../domain"
	"gorm.io/gorm"
)

type LanguageRepository struct {
	Conn *gorm.DB
}

func NewLanguageRepository(conn *gorm.DB) *LanguageRepository {
	return &LanguageRepository{Conn: conn}
}

func (repo *LanguageRepository) List() ([]string, error) {
	var names []string
	err := repo.Conn.Model(&domain.Language{}).Distinct().Pluck("name", &names).Error
	return names, err
}

func (repo *LanguageRepository) VersionsOfLanguage(name string) ([]string, error) {
	var versions []string
	err := repo.Conn.Model(&domain.Language{}).Where("name = ?", name).Pluck("version", &versions).Error
	return versions, err
}

func (repo *LanguageRepository) FindByNameAndVersion(name, version string) (domain.Language, error) {
	var language domain.Language
	err := repo.Conn.Where("name = ? AND version = ?", name, version).First(&language).Error
	return language, err
}

func (repo *LanguageRepository) Create(language domain.Language) error {
	return repo.Conn.Create(&language).Error
}

func (repo *LanguageRepository) DeleteByNameAndVersion(name, version string) error {
	return repo.Conn.Where("name = ? AND version = ?", name, version).Delete(domain.Language{}).Error
}
