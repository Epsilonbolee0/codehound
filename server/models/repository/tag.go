package repository

import (
	"../domain"
	"gorm.io/gorm"
)

type TagRepository struct {
	Conn *gorm.DB
}

func NewTagRepository(conn *gorm.DB) *TagRepository {
	return &TagRepository{Conn: conn}
}

func (repo *TagRepository) List() ([]domain.Tag, error) {
	var tags []domain.Tag
	err := repo.Conn.Find(&tags).Error
	return tags, err
}

func (repo *TagRepository) ListByCategory(category string) ([]string, error) {
	var content []string
	err := repo.Conn.Where("category = ?", category).Pluck("content", &content).Error
	return content, err
}

func (repo *TagRepository) FindOrCreate(category, content string) (domain.Tag, error) {
	var tag domain.Tag
	err := repo.Conn.FirstOrCreate(&tag, domain.Tag{Category: category, Content: content}).Error
	return tag, err
}

func (repo *TagRepository) Find(category, content string) (domain.Tag, error) {
	var tag domain.Tag
	err := repo.Conn.Where("category = ? AND content = ?", category, content).First(&tag).Error
	return tag, err
}

func (repo *TagRepository) Update(id uint, content string) error {
	return repo.Conn.Where("id = ?", id).Update("content", content).Error
}

func (repo *TagRepository) Delete(id uint) error {
	return repo.Conn.Where("id = ?", id).Delete(&domain.Tag{}).Error
}
