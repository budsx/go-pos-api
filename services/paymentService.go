package services

import (
	"go-pos-api/domain"
	"go-pos-api/dto"
	"go-pos-api/helpers"
	"go-pos-api/repositories"
)

type PaymentService interface {
	CreatePayment(dto.CreatePaymentInput) (domain.Payment, error)
}

type paymentService struct {
	paymentService  repositories.PaymentRepository
	midtransService MidtransService
}

func NewPaymentService(paymentRepository repositories.PaymentRepository, midtransService MidtransService) *paymentService {
	return &paymentService{paymentRepository, midtransService}
}

func (service *paymentService) CreatePayment(input dto.CreatePaymentInput) (domain.Payment, error) {
	payment := domain.Payment{}
	payment.OrderID = input.OrderID
	payment.Status = input.Status
	payment.Amount = input.Amount
	payment.PaymentCategory = input.PaymentCategory

	newPayment, err := service.paymentService.CreatePayment(payment)
	if err != nil {
		helpers.Error(err.Error())
		return newPayment, err
	}
	if input.PaymentCategory == "midtrans" {
		// get midtrans url and update transaction before return
		paymentURL, err := service.midtransService.GetPaymentURL(newPayment)
		if err != nil {
			helpers.Error("error" + err.Error())
			return newPayment, err
		}
		newPayment.SnapToken = paymentURL

		newPayment, err = service.paymentService.UpdatePayment(newPayment)
		if err != nil {
			helpers.Error("error" + err.Error())
			return newPayment, err
		}
		return newPayment, nil
	}
	return newPayment, nil
}
