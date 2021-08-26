package repository

import (
	"fmt"

	"../domain"
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type TestRepository struct {
	Conn *gorm.DB
}

func NewTestRepository(conn *gorm.DB) *TestRepository {
	return &TestRepository{Conn: conn}
}

func (repo *TestRepository) List(impl string) ([]domain.Test, error) {
	var tests []domain.Test
	err := repo.Conn.Where("implementation = ?", impl).Find(&tests).Error
	return tests, err
}

func (repo *TestRepository) ListByAuthor(author uint) ([]domain.Test, error) {
	var tests []domain.Test
	err := repo.Conn.Where("author = ?", author).Find(&tests).Error
	return tests, err
}

func (repo *TestRepository) Create(test domain.Test) error {
	return repo.Conn.Create(&test).Error
}

func (repo *TestRepository) UpdateDescription(id uint, description string) error {
	return repo.Conn.Model(&domain.Test{}).Where("id = ?", id).Update("description", description).Error
}

func (repo *TestRepository) UpdateArgument(field string, id, index uint, value string) error {
	var args []string
	row := repo.Conn.Model(&domain.Test{}).Where("id = ?", id).Select(field).Row()
	err := row.Scan(pq.Array(&args))
	if err != nil {
		return err
	}

	if index >= uint(len(args)) {
		return fmt.Errorf("index is out of bounds")
	}

	args[index] = value
	return repo.Conn.Model(&domain.Test{}).Where("id = ?", id).Update(field, pq.StringArray(args)).Error
}

func (repo *TestRepository) Delete(id uint) error {
	return repo.Conn.Where("id = ?", id).Delete(&domain.Test{}).Error
}
