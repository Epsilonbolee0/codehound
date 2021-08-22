package repository

import (
	"../domain"
	"gorm.io/gorm"
)

type ImplementationRepository struct {
	Conn *gorm.DB
}

func NewImplementationRepository(conn *gorm.DB) *ImplementationRepository {
	return &ImplementationRepository{Conn: conn}
}

func (repo *ImplementationRepository) List() ([]domain.Implementation, error) {
	var impls []domain.Implementation
	err := repo.Conn.Find(&impls).Error
	return impls, err
}

func (repo *ImplementationRepository) Find(name string) (domain.Implementation, error) {
	var impl domain.Implementation
	err := repo.Conn.Where("name = ?", name).First(&impl).Error
	return impl, err
}

func (repo *ImplementationRepository) FindRoot(name string) (domain.Version, error) {
	var versions []domain.Version
	repo.Conn.Model(&domain.Implementation{Name: name}).Order("date ASC").Association("Versions").Find(&versions)
	return versions[0], nil
}

func (repo *ImplementationRepository) Create(impl domain.Implementation) error {
	return repo.Conn.Create(&impl).Error
}
