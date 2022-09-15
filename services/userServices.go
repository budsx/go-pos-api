package services

import (
	"go-pos-api/domain"
	"go-pos-api/dto"
	"go-pos-api/helpers"
	"go-pos-api/repositories"

	"golang.org/x/crypto/bcrypt"
)

type UserServices interface {
	RegisterUser(request dto.RegisterRequest) (dto.RegisterResponse, *helpers.AppError)
}

type userServices struct {
	userServices repositories.UserRepositoryDB
}

func NewUserService(userRepository repositories.UserRepositoryDB) *userServices {
	return &userServices{userRepository}
}

func (services *userServices) RegisterUser(input dto.RegisterRequest) (dto.RegisterResponse, *helpers.AppError) {
	user := domain.User{}
	user.Name = input.Name
	user.Email = input.Email
	user.Role = input.Role
	user.Merchant = input.Merchant
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	registResponse := dto.RegisterResponse{}
	if err != nil {
		return registResponse, helpers.NewUnexpectedError("unexpected error")
	}
	user.Password = string(passwordHash)
	newUser, errRegis := services.userServices.RegisterUser(user)
	if errRegis != nil {
		return registResponse, helpers.NewUnexpectedError("unexpected error")
	}
	registResponse.ID = newUser.ID
	registResponse.Name = newUser.Name
	registResponse.Email = newUser.Email
	registResponse.Role = newUser.Role
	registResponse.Merchant = newUser.Merchant
	return registResponse, nil
}
