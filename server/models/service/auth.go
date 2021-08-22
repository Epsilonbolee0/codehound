package service

import (
	"regexp"
	"strings"

	"net/http"

	"../../utils"
	"../builder"
	"../domain"
	"../repository"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthService struct {
	accountRepo *repository.AccountRepository
}

func NewAuthService(accountRepo *repository.AccountRepository) *AuthService {
	return &AuthService{accountRepo}
}

func (auth *AuthService) Login(login, password string) map[string]interface{} {
	account, err := auth.accountRepo.FindByLogin(login)
	switch err {
	case nil:
		break
	case gorm.ErrRecordNotFound:
		return utils.Message(http.StatusNotFound, "Account not found!")
	default:
		return utils.Message(http.StatusInternalServerError, "Error occured while creating account!")
	}

	err = bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(password))
	if err == bcrypt.ErrMismatchedHashAndPassword {
		return utils.Message(http.StatusForbidden, "Invalid login credentials")
	}

	account.Password = ""

	resp := utils.Message(http.StatusOK, "Logged In")
	resp["account"] = account

	return resp
}

func (auth *AuthService) Register(login, email, password string) map[string]interface{} {
	accountBuilder := builder.NewAccountBuilder()
	account := accountBuilder.
		Login(login).
		Email(email).
		Password(password).
		Role("user").
		Build()

	if resp, ok := auth.validate(account); !ok {
		return resp
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(account.Password), bcrypt.DefaultCost)
	account.Password = string(hashedPassword)

	err := auth.accountRepo.Create(account)
	if err != nil {
		return utils.Message(http.StatusInternalServerError, "Failure occured while creating account")
	}
	account.Password = ""

	response := utils.Message(http.StatusOK, "Account has been created")
	response["account"] = account
	return response
}

func (auth *AuthService) validate(account domain.Account) (map[string]interface{}, bool) {
	if !auth.validateLogin(account.Login) {
		return utils.Message(http.StatusForbidden, "Login is invalid"), false
	}

	if !auth.validateEmail(account.Email) {
		return utils.Message(http.StatusForbidden, "Email address is invalid"), false
	}

	if !auth.validatePassword(account.Password) {
		return utils.Message(http.StatusForbidden, "Password is invalid"), false
	}

	if !auth.loginIsUnique(account.Login) {
		return utils.Message(http.StatusForbidden, "Login is already taken"), false
	}

	if !auth.emailIsUnique(account.Email) {
		return utils.Message(http.StatusForbidden, "Email is already used"), false
	}

	return utils.Message(http.StatusOK, "Validation passed"), true
}

func (auth *AuthService) validateLogin(login string) bool {
	matched, _ := regexp.MatchString(`^[a-zA-Z][a-zA-Z0-9-_\.]{1,20}$`, login)
	return matched
}

func (auth *AuthService) validateEmail(email string) bool {
	matched, _ := regexp.MatchString(`^[-\w.]+@([A-z0-9][-A-z0-9]+\.)+[A-z]{2,4}$`, email)
	return matched
}

func (auth *AuthService) validatePassword(password string) bool {
	capitalLetters := "ABCDEFGHIJKLMNOPQRSUVWXYZ"
	specialSymbols := "-_$.()#!&?/"
	digits := "1234567890"

	return len(password) >= 8 && strings.ContainsAny(password, capitalLetters) && strings.ContainsAny(password, specialSymbols) && strings.ContainsAny(password, digits)
}

func (auth *AuthService) loginIsUnique(login string) bool {
	_, err := auth.accountRepo.FindByLogin(login)
	return err == gorm.ErrRecordNotFound
}

func (auth *AuthService) emailIsUnique(login string) bool {
	_, err := auth.accountRepo.FindByEmail(login)
	return err == gorm.ErrRecordNotFound
}
