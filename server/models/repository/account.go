package repository

import (
	"../domain"
	"gorm.io/gorm"
)

type AccountRepository struct {
	Conn *gorm.DB
}

func NewAccountRepository(conn *gorm.DB) *AccountRepository {
	return &AccountRepository{Conn: conn}
}

func (repo *AccountRepository) List() ([]domain.Account, error) {
	var accounts []domain.Account
	err := repo.Conn.Find(&accounts).Error
	return accounts, err
}

func (repo *AccountRepository) FindByLogin(login string) (domain.Account, error) {
	var account domain.Account
	err := repo.Conn.Where("login = ?", login).First(&account).Error
	return account, err
}

func (repo *AccountRepository) FindByEmail(email string) (domain.Account, error) {
	var account domain.Account
	err := repo.Conn.Where("email = ?", email).First(&account).Error
	return account, err
}

func (repo *AccountRepository) UpdateIcon(login, icon string) error {
	return repo.Conn.Model(&domain.Account{}).Where("login = ?", login).Update("icon", icon).Error
}

func (repo *AccountRepository) UpdateRole(login, role string) error {
	return repo.Conn.Model(&domain.Account{}).Where("login = ?", login).Update("icon", role).Error
}

func (repo *AccountRepository) Create(account domain.Account) error {
	return repo.Conn.Create(&account).Error
}
