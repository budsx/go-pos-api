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
	GetAllUsers() ([]dto.LoginResponse, *helpers.AppError)
	GetUsersByID(int) (dto.LoginResponse, *helpers.AppError)
	UpdateUser(dto.RegisterRequest, int) (domain.User, *helpers.AppError)
	DeleteUser(int) (domain.User, *helpers.AppError)
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

func (services *userServices) GetAllUsers() ([]dto.LoginResponse, *helpers.AppError) {
	getUser, err := services.userRepository.GetAllUsers()
	userResponse := []dto.LoginResponse{}
	if err != nil {
		return userResponse, err
	}
	for _, user := range getUser {
		userResponse = append(userResponse, dto.LoginResponse{
			ID:       user.ID,
			Name:     user.Name,
			Email:    user.Email,
			Role:     user.Role,
			Merchant: user.Merchant,
		})
	}
	return userResponse, nil
}

func (services *userServices) GetUsersByID(user_id int) (dto.LoginResponse, *helpers.AppError) {
	getUser, err := services.userRepository.GetUsersByID(user_id)
	userResponse := dto.LoginResponse{}
	if err != nil {
		return userResponse, err
	} else {
		userResponse.ID = getUser.ID
		userResponse.Name = getUser.Name
		userResponse.Email = getUser.Email
		userResponse.Role = getUser.Role
		userResponse.Merchant = getUser.Merchant
		return userResponse, nil
	}
}

func (services *userServices) UpdateUser(request dto.RegisterRequest, user_id int) (domain.User, *helpers.AppError) {
	users := domain.User{}
	users.ID = user_id
	users.Name = request.Name
	users.Email = request.Email
	users.Role = request.Role
	users.Merchant = request.Merchant

	updateResponse := domain.User{}
	passwordHash, errHash := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.MinCost)
	if errHash != nil {
		return users, helpers.NewUnexpectedError("Bcrypt Error")
	}
	users.Password = string(passwordHash)

	updateUser, err := services.userRepository.UpdateUser(users, user_id)
	updateResponse.ID = user_id
	updateResponse.Name = updateUser.Name
	updateResponse.Email = updateUser.Email
	updateResponse.Role = updateUser.Role
	updateResponse.Merchant = updateUser.Merchant
	if err != nil {
		return updateResponse, err
	}
	return updateResponse, nil

}

func (services *userServices) DeleteUser(user_id int) (domain.User, *helpers.AppError) {
	delUser, err := services.userRepository.DeleteUser(user_id)
	if err != nil {
		return delUser, err
	}
	return delUser, nil
}
