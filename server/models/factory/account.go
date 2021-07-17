package factory

import (
	"../domain"
)

type accountModifier func(account *domain.Account)
type AccountBuilder struct {
	actions []accountModifier
}

func NewAccountBuilder() *AccountBuilder {
	return &AccountBuilder{}
}

func (b *AccountBuilder) Build() domain.Account {
	account := domain.Account{}
	for _, action := range b.actions {
		action(&account)
	}

	return account
}

func (b *AccountBuilder) Login(login string) *AccountBuilder {
	b.actions = append(b.actions, func(account *domain.Account) {
		account.Login = login
	})
	return b
}

func (b *AccountBuilder) Password(password string) *AccountBuilder {
	b.actions = append(b.actions, func(account *domain.Account) {
		account.Password = password
	})
	return b
}

func (b *AccountBuilder) Email(email string) *AccountBuilder {
	b.actions = append(b.actions, func(account *domain.Account) {
		account.Email = email
	})
	return b
}

func (b *AccountBuilder) Role(role string) *AccountBuilder {
	b.actions = append(b.actions, func(account *domain.Account) {
		account.Role = role
	})
	return b
}
