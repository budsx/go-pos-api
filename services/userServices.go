package services

import (
	"go-pos-api/domain"
	"go-pos-api/dto"
	"go-pos-api/helpers"
	"go-pos-api/repositories"

	"golang.org/x/crypto/bcrypt"
)

type UserServices interface {
	RegisterUser(request dto.RegisterRequest) (domain.User, *helpers.AppError)
	LoginUser(request dto.LoginRequest) (domain.User, *helpers.AppError)
}

type userServices struct {
	userRepository repositories.UserRepositoryDB
}

func NewUserService(userRepository repositories.UserRepositoryDB) *userServices {
	return &userServices{userRepository}
}

func (services *userServices) RegisterUser(request dto.RegisterRequest) (domain.User, *helpers.AppError) {
	user := domain.User{}
	user.Name = request.Name
	user.Email = request.Email
	user.Role = request.Role
	user.Merchant = request.Merchant
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.MinCost)
	if err != nil {
		return user, helpers.NewUnexpectedError("unexpected error")
	}
	user.Password = string(passwordHash)
	newUser, errRegis := services.userRepository.RegisterUser(user)
	if errRegis != nil {
		return newUser, helpers.NewUnexpectedError("unexpected error")
	}
	return newUser, nil
}

func (services *userServices) LoginUser(request dto.LoginRequest) (domain.User, *helpers.AppError) {
	email := request.Email
	password := request.Password
	newUser, errLogin := services.userRepository.LoginUser(email)
	if errLogin != nil {
		return newUser, helpers.NewUnexpectedError("Error Login")
	}
	if email == "" || password == "" {
		helpers.NewNotFoundError("Field email or password is empty")
	}
	err := helpers.VerifyPassword(password, newUser.Password)
	if err != nil {
		return newUser, helpers.NewBadRequestError("Wrong Password or Email")
	}
	return newUser, nil
}
