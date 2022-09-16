package repositories

import (
	"go-pos-api/domain"
	"go-pos-api/helpers"

	"gorm.io/gorm"
)

type UserRepositoryDB interface {
	RegisterUser(domain.User) (domain.User, *helpers.AppError)
	LoginUser(string) (domain.User, *helpers.AppError)
}

type userRepositoryDB struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepositoryDB {
	return &userRepositoryDB{db}
}

func (repository *userRepositoryDB) RegisterUser(user domain.User) (domain.User, *helpers.AppError) {
	var err error
	if err = repository.db.Create(&user).Error; err != nil {
		helpers.Error("Unexpected Error: " + err.Error())
		return user, helpers.NewUnexpectedError("Failed Create User")
	}
	return user, nil
}

func (repository *userRepositoryDB) LoginUser(email string) (domain.User, *helpers.AppError) {
	var user domain.User
	var err error
	if err = repository.db.Where("email = ?", email).Find(&user).Error; err != nil {
		helpers.Error("Unexpected Error: " + err.Error())
		return user, helpers.NewNotFoundError("User Not Found")
	}
	return user, nil
}
