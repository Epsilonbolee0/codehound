package repository

import (
	"../domain"
	"gorm.io/gorm"
)

type DescriptionRepository struct {
	Conn *gorm.DB
}

func NewDescriptionRepository(conn *gorm.DB) *DescriptionRepository {
	return &DescriptionRepository{Conn: conn}
}

func (repo *DescriptionRepository) ListByAuthor(author uint) ([]domain.Description, error) {
	var descriptions []domain.Description
	err := repo.Conn.Where("author = ?", author).Find(&descriptions).Error
	return descriptions, err
}

func (repo *DescriptionRepository) ListByImplementation(impl string) ([]domain.Description, error) {
	var descriptions []domain.Description
	err := repo.Conn.Where("implementation = ?", impl).Find(&descriptions).Error
	return descriptions, err
}

func (repo *DescriptionRepository) Create(description *domain.Description) error {
	return repo.Conn.Create(&description).Error
}

func (repo *DescriptionRepository) Update(id uint, content string) error {
	return repo.Conn.Where("id = ?", id).Update("content", content).Error
}

func (repo *DescriptionRepository) Delete(id uint) error {
	return repo.Conn.Where("id = ?", id).Delete(&domain.Description{}).Error
}
