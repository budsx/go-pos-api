package repositories

import (
	"go-pos-api/domain"
	"go-pos-api/helpers"

	"gorm.io/gorm"
)

type UserRepositoryDB interface {
	Register(domain.User) (domain.User, error)
	Login(string) (domain.User, error)
}

type userRepositoryDB struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepositoryDB {
	return &userRepositoryDB{db}
}

func (repo *userRepositoryDB) Register(user domain.User) (domain.User, error) {
	var err error
	if err = repo.db.Create(&user).Error; err != nil {
		helpers.Error("Unexpected Error: " + err.Error())
		return user, nil
	}
	return user, nil
}

func (repo *userRepositoryDB) Login(email string) (domain.User, error) {
	var user domain.User
	var err error
	if err = repo.db.Where("email = ?", email).Find(&user).Error; err != nil {
		helpers.Error("Unexpected Error: " + err.Error())
		return user, err
	}
	return user, nil
}
