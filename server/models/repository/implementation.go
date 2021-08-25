package repository

import (
	"fmt"

	"../domain"
	"github.com/lib/pq"
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
	if len(versions) == 0 {
		return domain.Version{}, gorm.ErrRecordNotFound
	}
	return versions[0], nil
}

func (repo *ImplementationRepository) Create(impl domain.Implementation) error {
	return repo.Conn.Create(&impl).Error
}

func (repo *ImplementationRepository) GetArguments(field, name string) ([]string, error) {
	var args []string
	row := repo.Conn.Model(&domain.Implementation{}).Where("name = ?", name).Select(field).Row()

	err := row.Scan(pq.Array(&args))
	return args, err
}

func (repo *ImplementationRepository) AddArgument(field, name, value string) error {
	var args []string

	row := repo.Conn.Model(&domain.Implementation{}).Where("name = ?", name).Select(field).Row()
	err := row.Scan(pq.Array(&args))
	if err != nil {
		return err
	}

	args = append(args, value)
	return repo.Conn.Model(&domain.Implementation{}).Where("name = ?", name).Update(field, pq.StringArray(args)).Error
}

func (repo *ImplementationRepository) UpdateArgument(field, name string, index uint, value string) error {
	var args []string

	row := repo.Conn.Model(&domain.Implementation{}).Where("name = ?", name).Select(field).Row()
	err := row.Scan(pq.Array(&args))
	if err != nil {
		return err
	}

	if index >= uint(len(args)) {
		return fmt.Errorf("index is out of bounds")
	}

	args[index] = value

	return repo.Conn.Model(&domain.Implementation{}).Where("name = ?", name).Update(field, pq.StringArray(args)).Error
}

func (repo *ImplementationRepository) RemoveArgument(field, name string, index uint) error {
	var args []string

	row := repo.Conn.Model(&domain.Implementation{}).Where("name = ?", name).Select(field).Row()
	err := row.Scan(pq.Array(&args))
	if err != nil {
		return err
	}

	if index >= uint(len(args)) {
		return fmt.Errorf("index is out of bounds")
	}

	copy(args[index:], args[index+1:])
	args[len(args)-1] = ""
	args = args[:len(args)-1]

	return repo.Conn.Model(&domain.Implementation{}).Where("name = ?", name).Update(field, pq.StringArray(args)).Error
}
