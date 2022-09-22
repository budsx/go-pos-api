package services

import (
	"go-pos-api/domain"
	"go-pos-api/repositories"
	"strconv"

	"github.com/veritrans/go-midtrans"
)

type midtransService struct {
	paymentRepository repositories.PaymentRepository
}

type MidtransService interface {
	GetPaymentURL(domain.Payment) (string, error)
	ProcessPayment(domain.TransactionNotificationFromMidtrans) error
}

func NewMidTransService(paymentRepository repositories.PaymentRepository) *midtransService {
	return &midtransService{paymentRepository: paymentRepository}
}

func (ms *midtransService) GetPaymentURL(payment domain.Payment) (string, error) {
	midclient := midtrans.NewClient()
	midclient.ServerKey = "SB-Mid-server-sliMiJsWEcB3xjZKzboB5BXO"
	midclient.ClientKey = "SB-Mid-client-Eev5kEZeGqUDcQXv"
	midclient.APIEnvType = midtrans.Sandbox
	snapGateway := midtrans.SnapGateway{
		Client: midclient,
	}
	snapReq := &midtrans.SnapReq{
		CustomerDetail: &midtrans.CustDetail{
			Email: "us.Email@email.com",
			FName: "us.name",
		},
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  strconv.Itoa(payment.ID),
			GrossAmt: int64(payment.Amount),
		},
	}
	snapToken, err := snapGateway.GetToken(snapReq)
	if err != nil {
		return "", err
	}

	return snapToken.RedirectURL, nil
}

func (ms *midtransService) ProcessPayment(input domain.TransactionNotificationFromMidtrans) error {
	// get result from midtrans
	intOrderId, _ := strconv.Atoi(input.OrderID)
	if (input.PaymentType == "qris" || input.PaymentType == "bank_transfer") && input.TransactionStatus == "settlement" && input.FraudStatus == "accept" {
		payment, err := ms.paymentRepository.GetPaymentByOrderIdAndAmount(intOrderId)
		if err != nil {
			return err
		}

		payment.Status = "success"
		_, errUpdatedPayment := ms.paymentRepository.UpdatePayment(payment)
		if errUpdatedPayment != nil {
			return errUpdatedPayment
		}
	}
	return nil
}
