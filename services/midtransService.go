package services

import (
	"go-pos-api/domain"
	"strconv"

	"github.com/veritrans/go-midtrans"
)

type midtransService struct {
}

type MidtransService interface {
	GetPaymentURL(domain.Payment) (string, error)
}

func NewMidTransService() *midtransService {
	return &midtransService{}
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
