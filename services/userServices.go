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
		return user, helpers.NewUnexpectedError("Bcrypt Error")
	}
	user.Password = string(passwordHash)

	avail, _ := services.userRepository.FindByEmail(user.Email)
	if avail.Email == user.Email {
		return avail, helpers.NewBadRequestError("User already exist")
	}
	newUser, errRegis := services.userRepository.RegisterUser(user)
	if errRegis != nil {
		return newUser, helpers.NewUnexpectedError("Failed register")
	}
	return newUser, nil
}

func (services *userServices) LoginUser(request dto.LoginRequest) (domain.User, *helpers.AppError) {
	email := request.Email
	password := request.Password
	user, errLogin := services.userRepository.FindByEmail(email)
	if errLogin != nil {
		return user, helpers.NewUnexpectedError("Error Login")
	}
	if email == "" || password == "" {
		helpers.NewNotFoundError("Field email or password is empty")
	}
	err := helpers.VerifyPassword(password, user.Password)
	if err != nil {
		return user, helpers.NewBadRequestError("Wrong Password or Email")
	}
	return user, nil
}
