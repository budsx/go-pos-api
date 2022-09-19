package repositories

import (
	"go-pos-api/domain"
	"go-pos-api/helpers"

	"gorm.io/gorm"
)

type PaymentRepository interface {
	CreatePayment(domain.Payment) (domain.Payment, error)
	UpdatePayment(domain.Payment) (domain.Payment, error)
}

type paymentRepository struct {
	db *gorm.DB
}

func NewPaymentRepository(db *gorm.DB) *paymentRepository {
	return &paymentRepository{db}
}

func (repository *paymentRepository) CreatePayment(payment domain.Payment) (domain.Payment, error) {
	err := repository.db.Create(&payment).Error
	if err != nil {
		helpers.Error("Error create payment" + err.Error())
		return payment, err
	}
	return payment, nil
}

func (repository *paymentRepository) UpdatePayment(payment domain.Payment) (domain.Payment, error) {
	err := repository.db.Save(&payment).Error
	if err != nil {
		helpers.Error(err.Error())
		return payment, err
	}
	return payment, nil
}
