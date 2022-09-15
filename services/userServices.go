package services

import (
	"errors"
	"go-pos-api/domain"
	"go-pos-api/dto"
	"go-pos-api/repositories"

	"golang.org/x/crypto/bcrypt"
)

type UserServices interface {
	RegisterUser(request dto.UserDTO) (domain.User, error)
	Login(request dto.LoginInput) (domain.User, error)
}

type userServices struct {
	userRepository repositories.UserRepositoryDB
}

func NewUserService(userRepository repositories.UserRepositoryDB) *userServices {
	return &userServices{userRepository}
}

func (services *userServices) Login(input dto.LoginInput) (domain.User, error) {
	email := input.Email
	password := input.Password

	user, err := services.userRepository.Login(email)
	if err != nil {
		return user, err
	}
	if user.Email == "" {
		return user, errors.New("no user found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return user, err
	}
	return user, nil
}

func (services *userServices) RegisterUser(input dto.UserDTO) (domain.User, error) {
	user := domain.User{}
	user.Name = input.Name
	user.Email = input.Email
	user.Role = input.Role
	user.Merchant = input.Merchant
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		return user, err
	}
	user.Password = string(passwordHash)
	newUser, err := services.userRepository.Register(user)
	if err != nil {
		return newUser, err
	}
	return newUser, nil
}
